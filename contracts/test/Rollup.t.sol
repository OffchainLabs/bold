// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";

import "../src/rollup/RollupProxy.sol";

import "../src/rollup/RollupCore.sol";
import "../src/rollup/RollupUserLogic.sol";
import "../src/rollup/RollupAdminLogic.sol";
import "../src/rollup/RollupCreator.sol";

import "../src/osp/OneStepProver0.sol";
import "../src/osp/OneStepProverMemory.sol";
import "../src/osp/OneStepProverMath.sol";
import "../src/osp/OneStepProverHostIo.sol";
import "../src/osp/OneStepProofEntry.sol";
import "../src/challenge/ChallengeManager.sol";

contract RollupTest is Test {
    address constant owner = address(1337);
    bytes32 constant wasmModuleRoot = keccak256("wasmModuleRoot");
    uint256 constant BASE_STAKE = 10;

    RollupProxy rollup;
    RollupUserLogic userRollup;
    RollupAdminLogic adminRollup;

    event RollupCreated(
        address indexed rollupAddress,
        address inboxAddress,
        address adminProxy,
        address sequencerInbox,
        address bridge
    );

    function setUp() public {
        OneStepProver0 oneStepProver = new OneStepProver0();
        OneStepProverMemory oneStepProverMemory = new OneStepProverMemory();
        OneStepProverMath oneStepProverMath = new OneStepProverMath();
        OneStepProverHostIo oneStepProverHostIo = new OneStepProverHostIo();
        OneStepProofEntry oneStepProofEntry = new OneStepProofEntry(
            oneStepProver, oneStepProverMemory, oneStepProverMath, oneStepProverHostIo);
        ChallengeManager challengeManagerImpl = new ChallengeManager();
        BridgeCreator bridgeCreator = new BridgeCreator();
        RollupCreator rollupCreator = new RollupCreator();
        RollupAdminLogic rollupAdminLogicImpl = new RollupAdminLogic();
        RollupUserLogic rollupUserLogicImpl = new RollupUserLogic();

        rollupCreator.setTemplates(
            bridgeCreator,
            oneStepProofEntry,
            challengeManagerImpl,
            rollupAdminLogicImpl,
            rollupUserLogicImpl,
            address(0),
            address(0)
        );

        Config memory config = Config({
            baseStake: BASE_STAKE,
            chainId: 0,
            confirmPeriodBlocks: 100,
            extraChallengeTimeBlocks: 100,
            owner: owner,
            sequencerInboxMaxTimeVariation: ISequencerInbox.MaxTimeVariation({
                delayBlocks: (60 * 60 * 24) / 15,
                futureBlocks: 12,
                delaySeconds: 60 * 60 * 24,
                futureSeconds: 60 * 60
            }),
            stakeToken: address(0),
            wasmModuleRoot: wasmModuleRoot,
            loserStakeEscrow: address(0),
            genesisBlockNum: 0
        });

        address expectedRollupAddr = address(uint160(uint256(keccak256(abi.encodePacked(bytes1(0xd6), bytes1(0x94), address(rollupCreator), bytes1(0x03))))));

        vm.expectEmit(true, true, false, false);
        emit RollupCreated(expectedRollupAddr, address(0), address(0), address(0), address(0));
        rollupCreator.createRollup(config, expectedRollupAddr);

        userRollup = RollupUserLogic(address(expectedRollupAddr));
        adminRollup = RollupAdminLogic(address(expectedRollupAddr));

        vm.startPrank(owner);
        adminRollup.setValidatorWhitelistDisabled(true);
        vm.stopPrank();
    }

    function testCreateAssertion() public {
        ExecutionState memory beforeState;
        beforeState.machineStatus = MachineStatus.FINISHED;
        ExecutionState memory afterState;
        afterState.machineStatus = MachineStatus.FINISHED;
        afterState.globalState.u64Vals[0] = 1;
        NewAssertionInputs memory inputs = NewAssertionInputs({
            beforeState: beforeState,
            afterState: afterState,
            numBlocks: 1,
            prevNum: 0,
            prevStateCommitment: bytes32(0),
            prevNodeInboxMaxCount: 1,
            expectedAssertionHash: bytes32(0)
        });
        vm.roll(block.number + 75); // bypass TIME_DELTA error
        userRollup.newStakeOnNewAssertion{value: BASE_STAKE}(inputs);
    }
}
