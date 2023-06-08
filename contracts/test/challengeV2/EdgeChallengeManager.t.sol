// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "forge-std/Test.sol";
import "./Utils.sol";
import "../MockAssertionChain.sol";
import "../../src/challengeV2/EdgeChallengeManager.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "../ERC20Mock.sol";
import "./StateTools.sol";

contract MockOneStepProofEntry is IOneStepProofEntry {
    function proveOneStep(ExecutionContext calldata, uint256, bytes32, bytes calldata proof)
        external
        pure
        returns (bytes32 afterHash)
    {
        return bytes32(proof);
    }

    function getMachineHash(ExecutionState calldata execState) external pure override returns (bytes32) {
        require(execState.machineStatus == MachineStatus.FINISHED, "BAD_MACHINE_STATUS");
        return GlobalStateLib.hash(execState.globalState);
    }
}

contract EdgeChallengeManagerTest is Test {
    Random rand = new Random();
    bytes32 genesisBlockHash = rand.hash();
    ExecutionState genesisState = StateToolsLib.randomState(rand, 4, genesisBlockHash, MachineStatus.FINISHED);
    bytes32 genesisStateHash = StateToolsLib.mockMachineHash(genesisState);
    bytes32 genesisAfterStateHash = RollupLib.executionStateHash(genesisState);
    ExecutionStateData genesisStateData = ExecutionStateData(genesisState, "");

    function genesisStates() internal view returns (bytes32[] memory) {
        bytes32[] memory genStates = new bytes32[](1);
        genStates[0] = genesisStateHash;
        return genStates;
    }

    bytes32 genesisRoot = MerkleTreeLib.root(ProofUtils.expansionFromLeaves(genesisStates(), 0, 1));

    uint256 genesisHeight = 2;
    uint64 inboxMsgCountGenesis = 7;
    uint64 inboxMsgCountAssertion = 12;

    bytes32 h1 = rand.hash();
    bytes32 h2 = rand.hash();
    uint256 height1 = 32;

    uint256 miniStakeVal = 1 ether;
    address excessStakeReceiver = address(77);
    address nobody = address(78);

    uint256 challengePeriodBlock = 1000;
    ExecutionStateData empty;

    function appendRandomStates(bytes32[] memory currentStates, uint256 numStates)
        internal
        returns (bytes32[] memory, bytes32[] memory)
    {
        bytes32[] memory newStates = rand.hashes(numStates);
        bytes32[] memory full = ArrayUtilsLib.concat(currentStates, newStates);
        bytes32[] memory exp = ProofUtils.expansionFromLeaves(full, 0, full.length);

        return (full, exp);
    }

    function deploy() internal returns (MockAssertionChain, EdgeChallengeManager, bytes32) {
        MockAssertionChain assertionChain = new MockAssertionChain();
        EdgeChallengeManager challengeManagerTemplate = new EdgeChallengeManager();
        EdgeChallengeManager challengeManager = EdgeChallengeManager(
            address(
                new TransparentUpgradeableProxy(
                    address(challengeManagerTemplate),
                    address(new ProxyAdmin()),
                    ""
                )
            )
        );
        challengeManager.initialize(
            assertionChain,
            challengePeriodBlock,
            new MockOneStepProofEntry(),
            2 ** 5,
            2 ** 5,
            2 ** 5,
            new ERC20Mock(
                "StakeToken",
                "ST",
                address(this),
                1000000 ether
            ),
            miniStakeVal,
            excessStakeReceiver
        );

        challengeManager.stakeToken().approve(address(challengeManager), type(uint256).max);

        bytes32 genesis = assertionChain.addAssertionUnsafe(0, genesisHeight, inboxMsgCountGenesis, genesisState, 0);
        return (assertionChain, challengeManager, genesis);
    }

    struct EdgeInitData {
        MockAssertionChain assertionChain;
        EdgeChallengeManager challengeManager;
        bytes32 genesis;
        bytes32 a1;
        bytes32 a2;
        ExecutionState a1State;
        ExecutionState a2State;
    }

    function deployAndInit() internal returns (EdgeInitData memory) {
        (MockAssertionChain assertionChain, EdgeChallengeManager challengeManager, bytes32 genesis) = deploy();

        ExecutionState memory a1State = StateToolsLib.randomState(
            rand, GlobalStateLib.getInboxPosition(genesisState.globalState), h1, MachineStatus.FINISHED
        );
        ExecutionState memory a2State = StateToolsLib.randomState(
            rand, GlobalStateLib.getInboxPosition(genesisState.globalState), h2, MachineStatus.FINISHED
        );

        // add one since heights are zero indexed in the history states
        bytes32 a1 = assertionChain.addAssertion(
            genesis, genesisHeight + height1, inboxMsgCountAssertion, genesisState, a1State, 0
        );
        bytes32 a2 = assertionChain.addAssertion(
            genesis, genesisHeight + height1, inboxMsgCountAssertion, genesisState, a2State, 0
        );

        return EdgeInitData({
            assertionChain: assertionChain,
            challengeManager: challengeManager,
            genesis: genesis,
            a1: a1,
            a2: a2,
            a1State: a1State,
            a2State: a2State
        });
    }

    function testRevertBlockNoFork() public {
        (MockAssertionChain assertionChain, EdgeChallengeManager challengeManager, bytes32 genesis) = deploy();

        ExecutionState memory a1State = StateToolsLib.randomState(
            rand, GlobalStateLib.getInboxPosition(genesisState.globalState), h1, MachineStatus.FINISHED
        );

        bytes32 a1 = assertionChain.addAssertion(
            genesis, genesisHeight + height1, inboxMsgCountAssertion, genesisState, a1State, 0
        );

        (bytes32[] memory states, bytes32[] memory exp) =
            appendRandomStatesBetween(genesisStates(), StateToolsLib.mockMachineHash(a1State), height1);

        vm.expectRevert(abi.encodeWithSelector(AssertionNoSibling.selector));
        challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.Block,
                endHistoryRoot: MerkleTreeLib.root(exp),
                endHeight: height1,
                claimId: a1,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(states, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(states, 1, states.length))
                    ),
                proof: abi.encode(
                    ProofUtils.generateInclusionProof(ProofUtils.rehashed(states), states.length - 1),
                    ExecutionStateData(genesisState, ""),
                    ExecutionStateData(a1State, "")
                    )
            })
        );
    }

    function testRevertBlockInvalidHeight() public {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states, bytes32[] memory exp) =
            appendRandomStatesBetween(genesisStates(), StateToolsLib.mockMachineHash(ei.a1State), height1);

        vm.expectRevert(abi.encodeWithSelector(InvalidEndHeight.selector, 1, 32));
        ei.challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.Block,
                endHistoryRoot: MerkleTreeLib.root(exp),
                endHeight: 1,
                claimId: ei.a1,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(states, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(states, 1, states.length))
                    ),
                proof: abi.encode(
                    ProofUtils.generateInclusionProof(ProofUtils.rehashed(states), states.length - 1),
                    ExecutionStateData(genesisState, ""),
                    ExecutionStateData(ei.a1State, "")
                    )
            })
        );
    }

    function testRevertBlockNoProof() public {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states, bytes32[] memory exp) =
            appendRandomStatesBetween(genesisStates(), StateToolsLib.mockMachineHash(ei.a1State), height1);

        vm.expectRevert(abi.encodeWithSelector(EmptyEdgeSpecificProof.selector));
        ei.challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.Block,
                endHistoryRoot: MerkleTreeLib.root(exp),
                endHeight: height1,
                claimId: ei.a1,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(states, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(states, 1, states.length))
                    ),
                proof: ""
            })
        );
    }

    function testRevertBlockInvalidProof() public {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states, bytes32[] memory exp) =
            appendRandomStatesBetween(genesisStates(), StateToolsLib.mockMachineHash(ei.a1State), height1);

        vm.expectRevert("Invalid inclusion proof");
        ei.challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.Block,
                endHistoryRoot: MerkleTreeLib.root(exp),
                endHeight: height1,
                claimId: ei.a1,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(states, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(states, 1, states.length))
                    ),
                proof: abi.encode(
                    ProofUtils.generateInclusionProof(ProofUtils.rehashed(states), 0),
                    ExecutionStateData(genesisState, ""),
                    ExecutionStateData(ei.a1State, "")
                    )
            })
        );
    }

    function testCanCreateEdgeWithStake()
        public
        returns (EdgeInitData memory, bytes32[] memory, bytes32[] memory, bytes32)
    {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states, bytes32[] memory exp) =
            appendRandomStatesBetween(genesisStates(), StateToolsLib.mockMachineHash(ei.a1State), height1);

        IERC20 stakeToken = ei.challengeManager.stakeToken();
        uint256 beforeBalance = stakeToken.balanceOf(address(this));
        bytes32 edgeId = ei.challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.Block,
                endHistoryRoot: MerkleTreeLib.root(exp),
                endHeight: height1,
                claimId: ei.a1,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(states, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(states, 1, states.length))
                    ),
                proof: abi.encode(
                    ProofUtils.generateInclusionProof(ProofUtils.rehashed(states), states.length - 1),
                    ExecutionStateData(genesisState, ""),
                    ExecutionStateData(ei.a1State, "")
                    )
            })
        );
        uint256 afterBalance = stakeToken.balanceOf(address(this));
        assertEq(beforeBalance - afterBalance, ei.challengeManager.stakeAmount(), "Staked");

        return (ei, states, exp, edgeId);
    }

    function testCanConfirmPs() public {
        (EdgeInitData memory ei,,, bytes32 edgeId) = testCanCreateEdgeWithStake();

        vm.roll(challengePeriodBlock + 2);

        bytes32[] memory ancestorEdges = new bytes32[](0);
        ei.challengeManager.confirmEdgeByTime(edgeId, ancestorEdges);

        assertTrue(ei.challengeManager.getEdge(edgeId).status == EdgeStatus.Confirmed, "Edge confirmed");
    }

    function testCanConfirmByChildren() public {
        (EdgeInitData memory ei, bytes32[] memory states1,, bytes32 edge1Id) = testCanCreateEdgeWithStake();

        vm.roll(block.number + 1);

        assertEq(ei.challengeManager.timeUnrivaled(edge1Id), 1, "Edge1 timer");
        {
            (bytes32[] memory states2, bytes32[] memory exp2) =
                appendRandomStatesBetween(genesisStates(), StateToolsLib.mockMachineHash(ei.a2State), height1);
            bytes32 edge2Id = ei.challengeManager.createLayerZeroEdge(
                CreateEdgeArgs({
                    edgeType: EdgeType.Block,
                    endHistoryRoot: MerkleTreeLib.root(exp2),
                    endHeight: height1,
                    claimId: ei.a2,
                    prefixProof: abi.encode(
                        ProofUtils.expansionFromLeaves(states2, 0, 1),
                        ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(states2, 1, states2.length))
                        ),
                    proof: abi.encode(
                        ProofUtils.generateInclusionProof(ProofUtils.rehashed(states2), states2.length - 1),
                        ExecutionStateData(genesisState, ""),
                        ExecutionStateData(ei.a2State, "")
                        )
                })
            );

            vm.roll(block.number + 2);
            assertEq(ei.challengeManager.timeUnrivaled(edge1Id), 1, "Edge1 timer");
            assertEq(ei.challengeManager.timeUnrivaled(edge2Id), 0, "Edge2 timer");
        }

        BisectionChildren memory children = bisect(ei.challengeManager, edge1Id, states1, 16, states1.length - 1);

        vm.roll(challengePeriodBlock + 5);

        bytes32[] memory ancestors = new bytes32[](1);
        ancestors[0] = edge1Id;
        ei.challengeManager.confirmEdgeByTime(children.lowerChildId, ancestors);
        ei.challengeManager.confirmEdgeByTime(children.upperChildId, ancestors);
        ei.challengeManager.confirmEdgeByChildren(edge1Id);

        assertTrue(ei.challengeManager.getEdge(edge1Id).status == EdgeStatus.Confirmed, "Edge confirmed");
    }

    function bisect(
        EdgeChallengeManager challengeManager,
        bytes32 edgeId,
        bytes32[] memory states,
        uint256 bisectionSize,
        uint256 endSize
    ) internal returns (BisectionChildren memory) {
        bytes32[] memory middleExp = ProofUtils.expansionFromLeaves(states, 0, bisectionSize + 1);
        bytes32[] memory upperStates = ArrayUtilsLib.slice(states, bisectionSize + 1, endSize + 1);

        (bytes32 lowerChildId, bytes32 upperChildId) = challengeManager.bisectEdge(
            edgeId,
            MerkleTreeLib.root(middleExp),
            abi.encode(middleExp, ProofUtils.generatePrefixProof(bisectionSize + 1, upperStates))
        );

        return BisectionChildren(lowerChildId, upperChildId);
    }

    struct BisectionChildren {
        bytes32 lowerChildId;
        bytes32 upperChildId;
    }

    struct BisectToForkOnlyArgs {
        EdgeChallengeManager challengeManager;
        bytes32 winningId;
        bytes32 losingId;
        bytes32[] winningLeaves;
        bytes32[] losingLeaves;
        bool skipLast;
    }

    function bisectToForkOnly(BisectToForkOnlyArgs memory args)
        internal
        returns (BisectionChildren[6] memory, BisectionChildren[6] memory)
    {
        BisectionChildren[6] memory winningEdges;
        BisectionChildren[6] memory losingEdges;

        winningEdges[5] = BisectionChildren(args.winningId, 0);
        losingEdges[5] = BisectionChildren(args.losingId, 0);

        // height 16
        winningEdges[4] = bisect(
            args.challengeManager, winningEdges[5].lowerChildId, args.winningLeaves, 16, args.winningLeaves.length - 1
        );
        losingEdges[4] = bisect(
            args.challengeManager, losingEdges[5].lowerChildId, args.losingLeaves, 16, args.losingLeaves.length - 1
        );

        // height 8
        winningEdges[3] = bisect(args.challengeManager, winningEdges[4].lowerChildId, args.winningLeaves, 8, 16);
        losingEdges[3] = bisect(args.challengeManager, losingEdges[4].lowerChildId, args.losingLeaves, 8, 16);

        // height 4
        winningEdges[2] = bisect(args.challengeManager, winningEdges[3].lowerChildId, args.winningLeaves, 4, 8);
        losingEdges[2] = bisect(args.challengeManager, losingEdges[3].lowerChildId, args.losingLeaves, 4, 8);

        winningEdges[1] = bisect(args.challengeManager, winningEdges[2].lowerChildId, args.winningLeaves, 2, 4);
        losingEdges[1] = bisect(args.challengeManager, losingEdges[2].lowerChildId, args.losingLeaves, 2, 4);

        // height 2
        winningEdges[0] = bisect(args.challengeManager, winningEdges[1].lowerChildId, args.winningLeaves, 1, 2);
        if (!args.skipLast) {
            losingEdges[0] = bisect(args.challengeManager, losingEdges[1].lowerChildId, args.losingLeaves, 1, 2);
        }

        return (winningEdges, losingEdges);
    }

    function appendRandomStatesBetween(bytes32[] memory currentStates, bytes32 endState, uint256 numStates)
        internal
        returns (bytes32[] memory, bytes32[] memory)
    {
        (bytes32[] memory states, bytes32[] memory exp) = appendRandomStates(currentStates, numStates - 1);
        bytes32[] memory fullStates = ArrayUtilsLib.append(states, endState);
        bytes32[] memory fullExp = MerkleTreeLib.appendLeaf(exp, endState);
        return (fullStates, fullExp);
    }

    function toDynamic(BisectionChildren[6] memory l) internal pure returns (BisectionChildren[] memory) {
        BisectionChildren[] memory d = new BisectionChildren[](6);
        for (uint256 i = 0; i < d.length; i++) {
            d[i] = l[i];
        }
        return d;
    }

    function concat(BisectionChildren[] memory arr1, BisectionChildren[] memory arr2)
        internal
        pure
        returns (BisectionChildren[] memory)
    {
        BisectionChildren[] memory full = new BisectionChildren[](arr1.length + arr2.length);
        for (uint256 i = 0; i < arr1.length; i++) {
            full[i] = arr1[i];
        }
        for (uint256 i = 0; i < arr2.length; i++) {
            full[arr1.length + i] = arr2[i];
        }
        return full;
    }

    function getAncestorsAbove(BisectionChildren[] memory layers, uint256 layer)
        internal
        pure
        returns (bytes32[] memory)
    {
        bytes32[] memory ancestors = new bytes32[](layers.length - 1 - layer);
        for (uint256 i = 0; i < layers.length - layer - 1; i++) {
            ancestors[i] = layers[i + layer + 1].lowerChildId;
        }
        return ancestors;
    }

    function generateEdgeProof(bytes32[] memory states1, bytes32[] memory bigStepStates)
        internal
        pure
        returns (bytes memory)
    {
        bytes32[] memory claimStartInclusionProof =
            ProofUtils.generateInclusionProof(ProofUtils.rehashed(ArrayUtilsLib.slice(states1, 0, 1)), 0);
        bytes32[] memory claimEndInclusionProof =
            ProofUtils.generateInclusionProof(ProofUtils.rehashed(ArrayUtilsLib.slice(states1, 0, 2)), 1);
        bytes32[] memory edgeInclusionProof =
            ProofUtils.generateInclusionProof(ProofUtils.rehashed(bigStepStates), bigStepStates.length - 1);
        return abi.encode(states1[0], states1[1], claimStartInclusionProof, claimEndInclusionProof, edgeInclusionProof);
    }

    function testRevertEmptyPrefixProof() public {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states1,, BisectionChildren[6] memory edges1,) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        (bytes32[] memory bigStepStates, bytes32[] memory bigStepExp) =
            appendRandomStatesBetween(genesisStates(), states1[1], height1);

        vm.expectRevert(abi.encodeWithSelector(EmptyPrefixProof.selector));
        ei.challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.BigStep,
                endHistoryRoot: MerkleTreeLib.root(bigStepExp),
                endHeight: height1,
                claimId: edges1[0].lowerChildId,
                prefixProof: "",
                proof: generateEdgeProof(states1, bigStepStates)
            })
        );
    }

    function testRevertInvalidPrefixProof() public {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states1,, BisectionChildren[6] memory edges1,) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        (bytes32[] memory bigStepStates, bytes32[] memory bigStepExp) =
            appendRandomStatesBetween(genesisStates(), states1[1], height1);

        vm.expectRevert("Post expansion root not equal post");
        ei.challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.BigStep,
                endHistoryRoot: MerkleTreeLib.root(bigStepExp),
                endHeight: height1,
                claimId: edges1[0].lowerChildId,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(states1, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(states1, 1, states1.length))
                    ),
                proof: generateEdgeProof(states1, bigStepStates)
            })
        );
    }

    function testRevertSubChallengeNotOneStepFork() public {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states1,, BisectionChildren[6] memory edges1,) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(
                ei.challengeManager,
                ei.a1,
                ei.a2,
                ei.a1State,
                ei.a2State,
                true // skipLast
            )
        );

        (bytes32[] memory bigStepStates, bytes32[] memory bigStepExp) =
            appendRandomStatesBetween(genesisStates(), states1[1], height1);

        vm.expectRevert(abi.encodeWithSelector(ClaimEdgeNotLengthOneRival.selector, edges1[0].lowerChildId));
        ei.challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.BigStep,
                endHistoryRoot: MerkleTreeLib.root(bigStepExp),
                endHeight: height1,
                claimId: edges1[0].lowerChildId,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(bigStepStates, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(bigStepStates, 1, bigStepStates.length))
                    ),
                proof: generateEdgeProof(states1, bigStepStates)
            })
        );
    }

    function testRevertSubChallengeNoProof() public {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states1,, BisectionChildren[6] memory edges1,) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        (bytes32[] memory bigStepStates, bytes32[] memory bigStepExp) =
            appendRandomStatesBetween(genesisStates(), states1[1], height1);

        vm.expectRevert(abi.encodeWithSelector(EmptyEdgeSpecificProof.selector));
        ei.challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.BigStep,
                endHistoryRoot: MerkleTreeLib.root(bigStepExp),
                endHeight: height1,
                claimId: edges1[0].lowerChildId,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(bigStepStates, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(bigStepStates, 1, bigStepStates.length))
                    ),
                proof: ""
            })
        );
    }

    function testRevertSubChallengeInvalidStartClaimProof() public {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states1,, BisectionChildren[6] memory edges1,) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        (bytes32[] memory bigStepStates, bytes32[] memory bigStepExp) =
            appendRandomStatesBetween(genesisStates(), states1[1], height1);

        bytes32[] memory claimEndInclusionProof =
            ProofUtils.generateInclusionProof(ProofUtils.rehashed(ArrayUtilsLib.slice(states1, 0, 2)), 1);
        bytes32[] memory edgeInclusionProof =
            ProofUtils.generateInclusionProof(ProofUtils.rehashed(bigStepStates), bigStepStates.length - 1);

        vm.expectRevert("Invalid inclusion proof");
        ei.challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.BigStep,
                endHistoryRoot: MerkleTreeLib.root(bigStepExp),
                endHeight: height1,
                claimId: edges1[0].lowerChildId,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(bigStepStates, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(bigStepStates, 1, bigStepStates.length))
                    ),
                proof: abi.encode(
                    states1[0], states1[1], claimEndInclusionProof, claimEndInclusionProof, edgeInclusionProof
                    )
            })
        );
    }

    function testRevertSubChallengeInvalidEndClaimProof() public {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states1,, BisectionChildren[6] memory edges1,) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        (bytes32[] memory bigStepStates, bytes32[] memory bigStepExp) =
            appendRandomStatesBetween(genesisStates(), states1[1], height1);

        bytes32[] memory claimStartInclusionProof =
            ProofUtils.generateInclusionProof(ProofUtils.rehashed(ArrayUtilsLib.slice(states1, 0, 1)), 0);
        bytes32[] memory edgeInclusionProof =
            ProofUtils.generateInclusionProof(ProofUtils.rehashed(bigStepStates), bigStepStates.length - 1);

        vm.expectRevert("Invalid inclusion proof");
        ei.challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.BigStep,
                endHistoryRoot: MerkleTreeLib.root(bigStepExp),
                endHeight: height1,
                claimId: edges1[0].lowerChildId,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(bigStepStates, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(bigStepStates, 1, bigStepStates.length))
                    ),
                proof: abi.encode(
                    states1[0], states1[1], claimStartInclusionProof, claimStartInclusionProof, edgeInclusionProof
                    )
            })
        );
    }

    function testRevertSubChallengeInvalidEdgeProof() public {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states1,, BisectionChildren[6] memory edges1,) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        (bytes32[] memory bigStepStates, bytes32[] memory bigStepExp) =
            appendRandomStatesBetween(genesisStates(), states1[1], height1);

        bytes32[] memory claimStartInclusionProof =
            ProofUtils.generateInclusionProof(ProofUtils.rehashed(ArrayUtilsLib.slice(states1, 0, 1)), 0);
        bytes32[] memory claimEndInclusionProof =
            ProofUtils.generateInclusionProof(ProofUtils.rehashed(ArrayUtilsLib.slice(states1, 0, 2)), 1);

        vm.expectRevert("Invalid inclusion proof");
        ei.challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.BigStep,
                endHistoryRoot: MerkleTreeLib.root(bigStepExp),
                endHeight: height1,
                claimId: edges1[0].lowerChildId,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(bigStepStates, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(bigStepStates, 1, bigStepStates.length))
                    ),
                proof: abi.encode(
                    states1[0], states1[1], claimStartInclusionProof, claimEndInclusionProof, claimStartInclusionProof
                    )
            })
        );
    }

    function testRevertBigStepInvalidHeight() public {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states1,, BisectionChildren[6] memory edges1,) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        (bytes32[] memory bigStepStates, bytes32[] memory bigStepExp) =
            appendRandomStatesBetween(genesisStates(), states1[1], height1);

        vm.expectRevert(abi.encodeWithSelector(InvalidEndHeight.selector, 1, 32));
        ei.challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.BigStep,
                endHistoryRoot: MerkleTreeLib.root(bigStepExp),
                endHeight: 1,
                claimId: edges1[0].lowerChildId,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(bigStepStates, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(bigStepStates, 1, bigStepStates.length))
                    ),
                proof: generateEdgeProof(states1, bigStepStates)
            })
        );
    }

    function testRevertBigStepInvalidClaimType() public {
        EdgeInitData memory ei = deployAndInit();

        (
            bytes32[] memory states1,
            bytes32[] memory states2,
            BisectionChildren[6] memory edges1,
            BisectionChildren[6] memory edges2
        ) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        bytes32[] memory bigStepStates1;
        bytes32 edge1BigStepId;
        {
            bytes32[] memory bigStepExp1;
            (bigStepStates1, bigStepExp1) = appendRandomStatesBetween(genesisStates(), states1[1], height1);

            edge1BigStepId = ei.challengeManager.createLayerZeroEdge(
                CreateEdgeArgs({
                    edgeType: EdgeType.BigStep,
                    endHistoryRoot: MerkleTreeLib.root(bigStepExp1),
                    endHeight: height1,
                    claimId: edges1[0].lowerChildId,
                    prefixProof: abi.encode(
                        ProofUtils.expansionFromLeaves(bigStepStates1, 0, 1),
                        ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(bigStepStates1, 1, bigStepStates1.length))
                        ),
                    proof: generateEdgeProof(states1, bigStepStates1)
                })
            );
        }

        bytes32[] memory bigStepStates2;
        bytes32 edge2BigStepId;
        {
            bytes32[] memory bigStepExp2;
            (bigStepStates2, bigStepExp2) = appendRandomStatesBetween(genesisStates(), states2[1], height1);

            edge2BigStepId = ei.challengeManager.createLayerZeroEdge(
                CreateEdgeArgs({
                    edgeType: EdgeType.BigStep,
                    endHistoryRoot: MerkleTreeLib.root(bigStepExp2),
                    endHeight: height1,
                    claimId: edges2[0].lowerChildId,
                    prefixProof: abi.encode(
                        ProofUtils.expansionFromLeaves(bigStepStates2, 0, 1),
                        ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(bigStepStates2, 1, bigStepStates2.length))
                        ),
                    proof: generateEdgeProof(states2, bigStepStates2)
                })
            );
        }

        (BisectionChildren[6] memory bigstepedges1,) = bisectToForkOnly(
            BisectToForkOnlyArgs(
                ei.challengeManager, edge1BigStepId, edge2BigStepId, bigStepStates1, bigStepStates2, false
            )
        );

        bytes32[] memory smallStepStates1;
        bytes32 edge1SmallStepId;
        {
            bytes32[] memory smallStepExp1;
            (smallStepStates1, smallStepExp1) = appendRandomStatesBetween(genesisStates(), bigStepStates1[1], height1);

            vm.expectRevert(abi.encodeWithSelector(ClaimEdgeInvalidType.selector, EdgeType.BigStep, EdgeType.BigStep));
            edge1SmallStepId = ei.challengeManager.createLayerZeroEdge(
                CreateEdgeArgs({
                    edgeType: EdgeType.BigStep,
                    endHistoryRoot: MerkleTreeLib.root(smallStepExp1),
                    endHeight: 1,
                    claimId: bigstepedges1[0].lowerChildId,
                    prefixProof: abi.encode(
                        ProofUtils.expansionFromLeaves(smallStepStates1, 0, 1),
                        ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(smallStepStates1, 1, smallStepStates1.length))
                        ),
                    proof: generateEdgeProof(bigStepStates1, smallStepStates1)
                })
            );
        }
    }

    function testRevertSmallStepInvalidClaimType() public {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states1,, BisectionChildren[6] memory edges1,) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        bytes32[] memory bigStepStates1;
        bytes32 edge1BigStepId;
        {
            bytes32[] memory bigStepExp1;
            (bigStepStates1, bigStepExp1) = appendRandomStatesBetween(genesisStates(), states1[1], height1);

            vm.expectRevert(abi.encodeWithSelector(ClaimEdgeInvalidType.selector, EdgeType.SmallStep, EdgeType.Block));
            edge1BigStepId = ei.challengeManager.createLayerZeroEdge(
                CreateEdgeArgs({
                    edgeType: EdgeType.SmallStep,
                    endHistoryRoot: MerkleTreeLib.root(bigStepExp1),
                    endHeight: height1,
                    claimId: edges1[0].lowerChildId,
                    prefixProof: abi.encode(
                        ProofUtils.expansionFromLeaves(bigStepStates1, 0, 1),
                        ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(bigStepStates1, 1, bigStepStates1.length))
                        ),
                    proof: generateEdgeProof(states1, bigStepStates1)
                })
            );
        }
    }

    function testRevertSmallStepInvalidHeight() public {
        EdgeInitData memory ei = deployAndInit();

        (
            bytes32[] memory states1,
            bytes32[] memory states2,
            BisectionChildren[6] memory edges1,
            BisectionChildren[6] memory edges2
        ) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        bytes32[] memory bigStepStates1;
        bytes32 edge1BigStepId;
        {
            bytes32[] memory bigStepExp1;
            (bigStepStates1, bigStepExp1) = appendRandomStatesBetween(genesisStates(), states1[1], height1);

            edge1BigStepId = ei.challengeManager.createLayerZeroEdge(
                CreateEdgeArgs({
                    edgeType: EdgeType.BigStep,
                    endHistoryRoot: MerkleTreeLib.root(bigStepExp1),
                    endHeight: height1,
                    claimId: edges1[0].lowerChildId,
                    prefixProof: abi.encode(
                        ProofUtils.expansionFromLeaves(bigStepStates1, 0, 1),
                        ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(bigStepStates1, 1, bigStepStates1.length))
                        ),
                    proof: generateEdgeProof(states1, bigStepStates1)
                })
            );
        }

        bytes32[] memory bigStepStates2;
        bytes32 edge2BigStepId;
        {
            bytes32[] memory bigStepExp2;
            (bigStepStates2, bigStepExp2) = appendRandomStatesBetween(genesisStates(), states2[1], height1);

            edge2BigStepId = ei.challengeManager.createLayerZeroEdge(
                CreateEdgeArgs({
                    edgeType: EdgeType.BigStep,
                    endHistoryRoot: MerkleTreeLib.root(bigStepExp2),
                    endHeight: height1,
                    claimId: edges2[0].lowerChildId,
                    prefixProof: abi.encode(
                        ProofUtils.expansionFromLeaves(bigStepStates2, 0, 1),
                        ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(bigStepStates2, 1, bigStepStates2.length))
                        ),
                    proof: generateEdgeProof(states2, bigStepStates2)
                })
            );
        }

        (BisectionChildren[6] memory bigstepedges1,) = bisectToForkOnly(
            BisectToForkOnlyArgs(
                ei.challengeManager, edge1BigStepId, edge2BigStepId, bigStepStates1, bigStepStates2, false
            )
        );

        bytes32[] memory smallStepStates1;
        bytes32 edge1SmallStepId;
        {
            bytes32[] memory smallStepExp1;
            (smallStepStates1, smallStepExp1) = appendRandomStatesBetween(genesisStates(), bigStepStates1[1], height1);

            vm.expectRevert(abi.encodeWithSelector(InvalidEndHeight.selector, 1, 32));
            edge1SmallStepId = ei.challengeManager.createLayerZeroEdge(
                CreateEdgeArgs({
                    edgeType: EdgeType.SmallStep,
                    endHistoryRoot: MerkleTreeLib.root(smallStepExp1),
                    endHeight: 1,
                    claimId: bigstepedges1[0].lowerChildId,
                    prefixProof: abi.encode(
                        ProofUtils.expansionFromLeaves(smallStepStates1, 0, 1),
                        ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(smallStepStates1, 1, smallStepStates1.length))
                        ),
                    proof: generateEdgeProof(bigStepStates1, smallStepStates1)
                })
            );
        }
    }

    function testCanConfirmByClaim() public {
        EdgeInitData memory ei = deployAndInit();

        (bytes32[] memory states1,, BisectionChildren[6] memory edges1,) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        (bytes32[] memory bigStepStates, bytes32[] memory bigStepExp) =
            appendRandomStatesBetween(genesisStates(), states1[1], height1);

        bytes32 edge1BigStepId = ei.challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.BigStep,
                endHistoryRoot: MerkleTreeLib.root(bigStepExp),
                endHeight: height1,
                claimId: edges1[0].lowerChildId,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(bigStepStates, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(bigStepStates, 1, bigStepStates.length))
                    ),
                proof: generateEdgeProof(states1, bigStepStates)
            })
        );

        vm.roll(challengePeriodBlock + 5);
        bytes32[] memory ancestors = new bytes32[](edges1.length);
        for (uint256 i = 0; i < edges1.length; i++) {
            ancestors[i] = edges1[i].lowerChildId;
        }

        ei.challengeManager.confirmEdgeByTime(edge1BigStepId, ancestors);

        ei.challengeManager.confirmEdgeByClaim(edges1[0].lowerChildId, edge1BigStepId);
        ei.challengeManager.confirmEdgeByTime(edges1[0].upperChildId, getAncestorsAbove(toDynamic(edges1), 0));

        ei.challengeManager.confirmEdgeByChildren(edges1[1].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(edges1[1].upperChildId, getAncestorsAbove(toDynamic(edges1), 1));

        ei.challengeManager.confirmEdgeByChildren(edges1[2].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(edges1[2].upperChildId, getAncestorsAbove(toDynamic(edges1), 2));

        ei.challengeManager.confirmEdgeByChildren(edges1[3].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(edges1[3].upperChildId, getAncestorsAbove(toDynamic(edges1), 3));

        ei.challengeManager.confirmEdgeByChildren(edges1[4].lowerChildId);

        assertTrue(ei.challengeManager.getEdge(edges1[4].lowerChildId).status == EdgeStatus.Confirmed, "Edge confirmed");
    }

    struct CreateBlockEdgesBisectArgs {
        EdgeChallengeManager challengeManager;
        bytes32 claim1Id;
        bytes32 claim2Id;
        ExecutionState endState1;
        ExecutionState endState2;
        bool skipLast;
    }

    struct CreateMachineEdgesBisectArgs {
        EdgeChallengeManager challengeManager;
        EdgeType eType;
        bytes32 claim1Id;
        bytes32 claim2Id;
        bytes32 endState1;
        bytes32 endState2;
        bool skipLast;
        bytes32[] forkStates1;
        bytes32[] forkStates2;
    }

    function createLayerZeroEdge(
        EdgeChallengeManager challengeManager,
        bytes32 claimId,
        ExecutionState memory endState,
        bytes32[] memory states,
        bytes32[] memory exp
    ) internal returns (bytes32) {
        bytes memory typeSpecificProof1 = abi.encode(
            ProofUtils.generateInclusionProof(ProofUtils.rehashed(states), states.length - 1),
            genesisStateData,
            ExecutionStateData(endState, "")
        );

        return challengeManager.createLayerZeroEdge(
            CreateEdgeArgs({
                edgeType: EdgeType.Block,
                endHistoryRoot: MerkleTreeLib.root(exp),
                endHeight: height1,
                claimId: claimId,
                prefixProof: abi.encode(
                    ProofUtils.expansionFromLeaves(states, 0, 1),
                    ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(states, 1, states.length))
                    ),
                proof: typeSpecificProof1
            })
        );
    }

    function createBlockEdgesAndBisectToFork(CreateBlockEdgesBisectArgs memory args)
        internal
        returns (bytes32[] memory, bytes32[] memory, BisectionChildren[6] memory, BisectionChildren[6] memory)
    {
        (bytes32[] memory states1, bytes32[] memory exp1) =
            appendRandomStatesBetween(genesisStates(), StateToolsLib.mockMachineHash(args.endState1), height1);
        bytes32 edge1Id = createLayerZeroEdge(args.challengeManager, args.claim1Id, args.endState1, states1, exp1);

        vm.roll(block.number + 1);

        assertEq(args.challengeManager.timeUnrivaled(edge1Id), 1, "Edge1 timer");

        (bytes32[] memory states2, bytes32[] memory exp2) =
            appendRandomStatesBetween(genesisStates(), StateToolsLib.mockMachineHash(args.endState2), height1);
        bytes32 edge2Id = createLayerZeroEdge(args.challengeManager, args.claim2Id, args.endState2, states2, exp2);

        vm.roll(block.number + 2);

        (BisectionChildren[6] memory edges1, BisectionChildren[6] memory edges2) = bisectToForkOnly(
            BisectToForkOnlyArgs(args.challengeManager, edge1Id, edge2Id, states1, states2, args.skipLast)
        );

        return (states1, states2, edges1, edges2);
    }

    function createMachineEdgesAndBisectToFork(CreateMachineEdgesBisectArgs memory args)
        internal
        returns (bytes32[] memory, bytes32[] memory, BisectionChildren[6] memory, BisectionChildren[6] memory)
    {
        (bytes32[] memory states1, bytes32[] memory exp1) =
            appendRandomStatesBetween(genesisStates(), args.endState1, height1);
        bytes32 edge1Id;
        {
            bytes memory typeSpecificProof1;
            {
                bytes32[] memory claimStartInclusionProof = ProofUtils.generateInclusionProof(
                    ProofUtils.rehashed(ArrayUtilsLib.slice(args.forkStates1, 0, 1)), 0
                );
                bytes32[] memory claimEndInclusionProof =
                    ProofUtils.generateInclusionProof(ProofUtils.rehashed(args.forkStates1), 1);
                bytes32[] memory edgeInclusionProof =
                    ProofUtils.generateInclusionProof(ProofUtils.rehashed(states1), states1.length - 1);
                typeSpecificProof1 = abi.encode(
                    genesisStateHash,
                    args.endState1,
                    claimStartInclusionProof,
                    claimEndInclusionProof,
                    edgeInclusionProof
                );
            }
            edge1Id = args.challengeManager.createLayerZeroEdge(
                CreateEdgeArgs({
                    edgeType: args.eType,
                    endHistoryRoot: MerkleTreeLib.root(exp1),
                    endHeight: height1,
                    claimId: args.claim1Id,
                    prefixProof: abi.encode(
                        ProofUtils.expansionFromLeaves(states1, 0, 1),
                        ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(states1, 1, states1.length))
                        ),
                    proof: typeSpecificProof1
                })
            );
        }

        vm.roll(block.number + 1);

        assertEq(args.challengeManager.timeUnrivaled(edge1Id), 1, "Edge1 timer");

        (bytes32[] memory states2, bytes32[] memory exp2) =
            appendRandomStatesBetween(genesisStates(), args.endState2, height1);
        bytes32 edge2Id;
        {
            bytes memory typeSpecificProof2;
            {
                bytes32[] memory claimStartInclusionProof = ProofUtils.generateInclusionProof(
                    ProofUtils.rehashed(ArrayUtilsLib.slice(args.forkStates2, 0, 1)), 0
                );
                bytes32[] memory claimEndInclusionProof =
                    ProofUtils.generateInclusionProof(ProofUtils.rehashed(args.forkStates2), 1);
                bytes32[] memory edgeInclusionProof =
                    ProofUtils.generateInclusionProof(ProofUtils.rehashed(states2), states2.length - 1);
                typeSpecificProof2 = abi.encode(
                    genesisStateHash,
                    args.endState2,
                    claimStartInclusionProof,
                    claimEndInclusionProof,
                    edgeInclusionProof
                );
            }
            edge2Id = args.challengeManager.createLayerZeroEdge(
                CreateEdgeArgs({
                    edgeType: args.eType,
                    endHistoryRoot: MerkleTreeLib.root(exp2),
                    endHeight: height1,
                    claimId: args.claim2Id,
                    prefixProof: abi.encode(
                        ProofUtils.expansionFromLeaves(states2, 0, 1),
                        ProofUtils.generatePrefixProof(1, ArrayUtilsLib.slice(states2, 1, states2.length))
                        ),
                    proof: typeSpecificProof2
                })
            );
        }

        vm.roll(block.number + 2);

        (BisectionChildren[6] memory edges1, BisectionChildren[6] memory edges2) = bisectToForkOnly(
            BisectToForkOnlyArgs(args.challengeManager, edge1Id, edge2Id, states1, states2, args.skipLast)
        );

        return (states1, states2, edges1, edges2);
    }

    function testCanConfirmByClaimSubChallenge() public {
        EdgeInitData memory ei = deployAndInit();

        (
            bytes32[] memory blockStates1,
            bytes32[] memory blockStates2,
            BisectionChildren[6] memory blockEdges1,
            BisectionChildren[6] memory blockEdges2
        ) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        (
            bytes32[] memory bigStepStates1,
            bytes32[] memory bigStepStates2,
            BisectionChildren[6] memory bigStepEdges1,
            BisectionChildren[6] memory bigStepEdges2
        ) = createMachineEdgesAndBisectToFork(
            CreateMachineEdgesBisectArgs(
                ei.challengeManager,
                EdgeType.BigStep,
                blockEdges1[0].lowerChildId,
                blockEdges2[0].lowerChildId,
                blockStates1[1],
                blockStates2[1],
                false,
                ArrayUtilsLib.slice(blockStates1, 0, 2),
                ArrayUtilsLib.slice(blockStates2, 0, 2)
            )
        );

        (,, BisectionChildren[6] memory smallStepEdges1,) = createMachineEdgesAndBisectToFork(
            CreateMachineEdgesBisectArgs(
                ei.challengeManager,
                EdgeType.SmallStep,
                bigStepEdges1[0].lowerChildId,
                bigStepEdges2[0].lowerChildId,
                bigStepStates1[1],
                bigStepStates2[1],
                true,
                ArrayUtilsLib.slice(bigStepStates1, 0, 2),
                ArrayUtilsLib.slice(bigStepStates2, 0, 2)
            )
        );

        vm.roll(challengePeriodBlock + 11);

        BisectionChildren[] memory allWinners =
            concat(concat(toDynamic(smallStepEdges1), toDynamic(bigStepEdges1)), toDynamic(blockEdges1));

        ei.challengeManager.confirmEdgeByTime(allWinners[0].lowerChildId, getAncestorsAbove(allWinners, 0));
        ei.challengeManager.confirmEdgeByTime(allWinners[0].upperChildId, getAncestorsAbove(allWinners, 0));

        ei.challengeManager.confirmEdgeByChildren(allWinners[1].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[1].upperChildId, getAncestorsAbove(allWinners, 1));

        ei.challengeManager.confirmEdgeByChildren(allWinners[2].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[2].upperChildId, getAncestorsAbove(allWinners, 2));

        ei.challengeManager.confirmEdgeByChildren(allWinners[3].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[3].upperChildId, getAncestorsAbove(allWinners, 3));

        ei.challengeManager.confirmEdgeByChildren(allWinners[4].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[4].upperChildId, getAncestorsAbove(allWinners, 4));

        ei.challengeManager.confirmEdgeByChildren(allWinners[5].lowerChildId);

        ei.challengeManager.confirmEdgeByClaim(allWinners[6].lowerChildId, allWinners[5].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[6].upperChildId, getAncestorsAbove(allWinners, 6));

        ei.challengeManager.confirmEdgeByChildren(allWinners[7].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[7].upperChildId, getAncestorsAbove(allWinners, 7));

        ei.challengeManager.confirmEdgeByChildren(allWinners[8].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[8].upperChildId, getAncestorsAbove(allWinners, 8));

        ei.challengeManager.confirmEdgeByChildren(allWinners[9].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[9].upperChildId, getAncestorsAbove(allWinners, 9));

        ei.challengeManager.confirmEdgeByChildren(allWinners[10].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[10].upperChildId, getAncestorsAbove(allWinners, 10));

        ei.challengeManager.confirmEdgeByChildren(allWinners[11].lowerChildId);

        ei.challengeManager.confirmEdgeByClaim(allWinners[12].lowerChildId, allWinners[11].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[12].upperChildId, getAncestorsAbove(allWinners, 12));

        ei.challengeManager.confirmEdgeByChildren(allWinners[13].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[13].upperChildId, getAncestorsAbove(allWinners, 13));

        ei.challengeManager.confirmEdgeByChildren(allWinners[14].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[14].upperChildId, getAncestorsAbove(allWinners, 14));

        ei.challengeManager.confirmEdgeByChildren(allWinners[15].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[15].upperChildId, getAncestorsAbove(allWinners, 15));

        ei.challengeManager.confirmEdgeByChildren(allWinners[16].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[16].upperChildId, getAncestorsAbove(allWinners, 16));

        ei.challengeManager.confirmEdgeByChildren(allWinners[17].lowerChildId);

        assertTrue(
            ei.challengeManager.getEdge(allWinners[14].lowerChildId).status == EdgeStatus.Confirmed, "Edge confirmed"
        );
    }

    function testCanConfirmByOneStep() public returns (EdgeInitData memory, BisectionChildren[] memory) {
        EdgeInitData memory ei = deployAndInit();

        (
            bytes32[] memory blockStates1,
            bytes32[] memory blockStates2,
            BisectionChildren[6] memory blockEdges1,
            BisectionChildren[6] memory blockEdges2
        ) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        (
            bytes32[] memory bigStepStates1,
            bytes32[] memory bigStepStates2,
            BisectionChildren[6] memory bigStepEdges1,
            BisectionChildren[6] memory bigStepEdges2
        ) = createMachineEdgesAndBisectToFork(
            CreateMachineEdgesBisectArgs(
                ei.challengeManager,
                EdgeType.BigStep,
                blockEdges1[0].lowerChildId,
                blockEdges2[0].lowerChildId,
                blockStates1[1],
                blockStates2[1],
                false,
                ArrayUtilsLib.slice(blockStates1, 0, 2),
                ArrayUtilsLib.slice(blockStates2, 0, 2)
            )
        );

        (bytes32[] memory smallStepStates1,, BisectionChildren[6] memory smallStepEdges1,) =
        createMachineEdgesAndBisectToFork(
            CreateMachineEdgesBisectArgs(
                ei.challengeManager,
                EdgeType.SmallStep,
                bigStepEdges1[0].lowerChildId,
                bigStepEdges2[0].lowerChildId,
                bigStepStates1[1],
                bigStepStates2[1],
                false,
                ArrayUtilsLib.slice(bigStepStates1, 0, 2),
                ArrayUtilsLib.slice(bigStepStates2, 0, 2)
            )
        );

        vm.roll(challengePeriodBlock + 11);

        BisectionChildren[] memory allWinners =
            concat(concat(toDynamic(smallStepEdges1), toDynamic(bigStepEdges1)), toDynamic(blockEdges1));

        bytes32[] memory firstStates = new bytes32[](2);
        firstStates[0] = smallStepStates1[0];
        firstStates[1] = smallStepStates1[1];

        ei.challengeManager.confirmEdgeByOneStepProof(
            allWinners[0].lowerChildId,
            OneStepData({beforeHash: firstStates[0], proof: abi.encodePacked(firstStates[1])}),
            ConfigData({
                wasmModuleRoot: bytes32(0),
                requiredStake: 0,
                challengeManager: address(0),
                confirmPeriodBlocks: 0,
                nextInboxPosition: inboxMsgCountGenesis
            }),
            ProofUtils.generateInclusionProof(ProofUtils.rehashed(genesisStates()), 0),
            ProofUtils.generateInclusionProof(ProofUtils.rehashed(firstStates), 1)
        );
        bytes32[] memory above = getAncestorsAbove(allWinners, 0);
        ei.challengeManager.confirmEdgeByTime(allWinners[0].upperChildId, above);

        ei.challengeManager.confirmEdgeByChildren(allWinners[1].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[1].upperChildId, getAncestorsAbove(allWinners, 1));

        ei.challengeManager.confirmEdgeByChildren(allWinners[2].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[2].upperChildId, getAncestorsAbove(allWinners, 2));

        ei.challengeManager.confirmEdgeByChildren(allWinners[3].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[3].upperChildId, getAncestorsAbove(allWinners, 3));

        ei.challengeManager.confirmEdgeByChildren(allWinners[4].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[4].upperChildId, getAncestorsAbove(allWinners, 4));

        ei.challengeManager.confirmEdgeByChildren(allWinners[5].lowerChildId);

        ei.challengeManager.confirmEdgeByClaim(allWinners[6].lowerChildId, allWinners[5].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[6].upperChildId, getAncestorsAbove(allWinners, 6));

        ei.challengeManager.confirmEdgeByChildren(allWinners[7].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[7].upperChildId, getAncestorsAbove(allWinners, 7));

        ei.challengeManager.confirmEdgeByChildren(allWinners[8].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[8].upperChildId, getAncestorsAbove(allWinners, 8));

        ei.challengeManager.confirmEdgeByChildren(allWinners[9].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[9].upperChildId, getAncestorsAbove(allWinners, 9));

        ei.challengeManager.confirmEdgeByChildren(allWinners[10].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[10].upperChildId, getAncestorsAbove(allWinners, 10));

        ei.challengeManager.confirmEdgeByChildren(allWinners[11].lowerChildId);

        ei.challengeManager.confirmEdgeByClaim(allWinners[12].lowerChildId, allWinners[11].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[12].upperChildId, getAncestorsAbove(allWinners, 12));

        ei.challengeManager.confirmEdgeByChildren(allWinners[13].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[13].upperChildId, getAncestorsAbove(allWinners, 13));

        ei.challengeManager.confirmEdgeByChildren(allWinners[14].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[14].upperChildId, getAncestorsAbove(allWinners, 14));

        ei.challengeManager.confirmEdgeByChildren(allWinners[15].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[15].upperChildId, getAncestorsAbove(allWinners, 15));

        ei.challengeManager.confirmEdgeByChildren(allWinners[16].lowerChildId);
        ei.challengeManager.confirmEdgeByTime(allWinners[16].upperChildId, getAncestorsAbove(allWinners, 16));

        ei.challengeManager.confirmEdgeByChildren(allWinners[17].lowerChildId);

        assertTrue(
            ei.challengeManager.getEdge(allWinners[17].lowerChildId).status == EdgeStatus.Confirmed, "Edge confirmed"
        );

        return (ei, allWinners);
    }

    function testExcessStakeReceived() external {
        (EdgeInitData memory ei,) = testCanConfirmByOneStep();
        IERC20 stakeToken = ei.challengeManager.stakeToken();
        assertEq(stakeToken.balanceOf(excessStakeReceiver), ei.challengeManager.stakeAmount(), "Excess stake received");
    }

    function testCanRefundStake() external {
        (EdgeInitData memory ei, BisectionChildren[] memory allWinners) = testCanConfirmByOneStep();

        IERC20 stakeToken = ei.challengeManager.stakeToken();
        uint256 beforeBalance = stakeToken.balanceOf(address(this));
        vm.prank(nobody); // call refund as nobody
        ei.challengeManager.refundStake(allWinners[17].lowerChildId);
        uint256 afterBalance = stakeToken.balanceOf(address(this));
        assertEq(afterBalance - beforeBalance, ei.challengeManager.stakeAmount(), "Stake refunded");
    }

    function testRevertRefundStakeTwice() external {
        (EdgeInitData memory ei, BisectionChildren[] memory allWinners) = testCanConfirmByOneStep();
        ei.challengeManager.refundStake(allWinners[17].lowerChildId);
        vm.expectRevert(abi.encodeWithSelector(EdgeAlreadyRefunded.selector, allWinners[17].lowerChildId));
        ei.challengeManager.refundStake(allWinners[17].lowerChildId);
    }

    function testRevertRefundStakeNotLayerZero() external {
        (EdgeInitData memory ei, BisectionChildren[] memory allWinners) = testCanConfirmByOneStep();
        vm.expectRevert(abi.encodeWithSelector(EdgeNotLayerZero.selector, allWinners[16].lowerChildId, 0, 0));
        ei.challengeManager.refundStake(allWinners[16].lowerChildId);
    }

    function testRevertRefundStakeBigStep() external {
        (EdgeInitData memory ei, BisectionChildren[] memory allWinners) = testCanConfirmByOneStep();
        vm.expectRevert(abi.encodeWithSelector(EdgeTypeNotBlock.selector, EdgeType.BigStep));
        ei.challengeManager.refundStake(allWinners[11].lowerChildId);
    }

    function testRevertRefundStakeSmallStep() external {
        (EdgeInitData memory ei, BisectionChildren[] memory allWinners) = testCanConfirmByOneStep();
        vm.expectRevert(abi.encodeWithSelector(EdgeTypeNotBlock.selector, EdgeType.SmallStep));
        ei.challengeManager.refundStake(allWinners[5].lowerChildId);
    }

    function testRevertRefundStakeNotConfirmed() external {
        (EdgeInitData memory ei,,, bytes32 edgeId) = testCanCreateEdgeWithStake();

        vm.expectRevert(abi.encodeWithSelector(EdgeNotConfirmed.selector, edgeId, EdgeStatus.Pending));
        ei.challengeManager.refundStake(edgeId);
    }

    function testGetPrevAssertionId() public {
        EdgeInitData memory ei = deployAndInit();

        (
            bytes32[] memory blockStates1,
            bytes32[] memory blockStates2,
            BisectionChildren[6] memory blockEdges1,
            BisectionChildren[6] memory blockEdges2
        ) = createBlockEdgesAndBisectToFork(
            CreateBlockEdgesBisectArgs(ei.challengeManager, ei.a1, ei.a2, ei.a1State, ei.a2State, false)
        );

        (
            bytes32[] memory bigStepStates1,
            bytes32[] memory bigStepStates2,
            BisectionChildren[6] memory bigStepEdges1,
            BisectionChildren[6] memory bigStepEdges2
        ) = createMachineEdgesAndBisectToFork(
            CreateMachineEdgesBisectArgs(
                ei.challengeManager,
                EdgeType.BigStep,
                blockEdges1[0].lowerChildId,
                blockEdges2[0].lowerChildId,
                blockStates1[1],
                blockStates2[1],
                false,
                ArrayUtilsLib.slice(blockStates1, 0, 2),
                ArrayUtilsLib.slice(blockStates2, 0, 2)
            )
        );

        (,, BisectionChildren[6] memory smallStepEdges1, BisectionChildren[6] memory smallStepEdges2) =
        createMachineEdgesAndBisectToFork(
            CreateMachineEdgesBisectArgs(
                ei.challengeManager,
                EdgeType.SmallStep,
                bigStepEdges1[0].lowerChildId,
                bigStepEdges2[0].lowerChildId,
                bigStepStates1[1],
                bigStepStates2[1],
                false,
                ArrayUtilsLib.slice(bigStepStates1, 0, 2),
                ArrayUtilsLib.slice(bigStepStates2, 0, 2)
            )
        );

        for (uint256 i = 0; i < smallStepEdges1.length; i++) {
            bytes32 childId = smallStepEdges1[i].lowerChildId;
            assertEq(ei.challengeManager.getPrevAssertionId(childId), ei.genesis);
        }

        for (uint256 i = 0; i < smallStepEdges2.length; i++) {
            bytes32 childId = smallStepEdges2[i].lowerChildId;
            assertEq(ei.challengeManager.getPrevAssertionId(childId), ei.genesis);
        }

        for (uint256 i = 0; i < bigStepEdges1.length; i++) {
            bytes32 childId = bigStepEdges1[i].lowerChildId;
            assertEq(ei.challengeManager.getPrevAssertionId(childId), ei.genesis);
        }

        for (uint256 i = 0; i < bigStepEdges2.length; i++) {
            bytes32 childId = bigStepEdges2[i].lowerChildId;
            assertEq(ei.challengeManager.getPrevAssertionId(childId), ei.genesis);
        }

        for (uint256 i = 0; i < blockEdges1.length; i++) {
            bytes32 childId = blockEdges1[i].lowerChildId;
            assertEq(ei.challengeManager.getPrevAssertionId(childId), ei.genesis);
        }

        for (uint256 i = 0; i < blockEdges2.length; i++) {
            bytes32 childId = blockEdges2[i].lowerChildId;
            assertEq(ei.challengeManager.getPrevAssertionId(childId), ei.genesis);
        }
    }
}
