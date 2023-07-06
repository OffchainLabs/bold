// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/utils/Create2Upgradeable.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "./RollupProxy.sol";
import "./RollupLib.sol";
import "./RollupAdminLogic.sol";

struct Node {
    // Hash of the state of the chain as of this node
    bytes32 stateHash;
    // Hash of the data that can be challenged
    bytes32 challengeHash;
    // Hash of the data that will be committed if this node is confirmed
    bytes32 confirmData;
    // Index of the node previous to this one
    uint64 prevNum;
    // Deadline at which this node can be confirmed
    uint64 deadlineBlock;
    // Deadline at which a child of this node can be confirmed
    uint64 noChildConfirmedBeforeBlock;
    // Number of stakers staked on this node. This includes real stakers and zombies
    uint64 stakerCount;
    // Number of stakers staked on a child node. This includes real stakers and zombies
    uint64 childStakerCount;
    // This value starts at zero and is set to a value when the first child is created. After that it is constant until the node is destroyed or the owner destroys pending nodes
    uint64 firstChildBlock;
    // The number of the latest child of this node to be created
    uint64 latestChildNumber;
    // The block number when this node was created
    uint64 createdAtBlock;
    // A hash of all the data needed to determine this node's validity, to protect against reorgs
    bytes32 nodeHash;
}

struct OldStaker {
    uint256 amountStaked;
    uint64 index;
    uint64 latestStakedNode;
    // currentChallenge is 0 if staker is not in a challenge
    uint64 currentChallenge; // 1. cannot have current challenge
    bool isStaked; // 2. must be staked
}

interface IOldRollup {
    function wasmModuleRoot() external view returns (bytes32);
    function latestConfirmed() external view returns (uint64);
    function getNode(uint64 nodeNum) external view returns (Node memory);
    function getStakerAddress(uint64 stakerNum) external view returns (address);
    function stakerCount() external view returns (uint64);
    function getStaker(address staker) external view returns (OldStaker memory);
}

interface IOldRollupAdmin {
    function forceRefundStaker(address[] memory stacker) external;
    function pause() external;
}

/// @title  Provides pre-images to a state hash
/// @notice We need to use the execution state of the latest confirmed node as the genesis
///         in the new rollup. However the this full state is not available on chain, only
///         the state hash is, which commits to this. This lookup contract should be deployed
///         before the upgrade, and just before the upgrade is executed the pre-image of the
///         latest confirmed state hash should be populated here. The upgrade contact can then
///         fetch this information and verify it before using it.
contract StateHashPreImageLookup {
    using GlobalStateLib for GlobalState;

    event HashSet(bytes32 h, ExecutionState execState, uint inboxMaxCount);

    mapping(bytes32 => bytes) internal preImages;

    function stateHash(ExecutionState calldata execState, uint256 inboxMaxCount) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(execState.globalState.hash(), inboxMaxCount, execState.machineStatus));
    }

    function set(bytes32 h, ExecutionState calldata execState, uint256 inboxMaxCount) public {
        require(h == stateHash(execState, inboxMaxCount), "Invalid hash");
        preImages[h] = abi.encode(execState, inboxMaxCount);
        emit HashSet(h, execState, inboxMaxCount);
    }

    function get(bytes32 h) public view returns (ExecutionState memory execState, uint256 inboxMaxCount) {
        return abi.decode(preImages[h], (ExecutionState, uint256));
    }
}

/// @title  Forwards calls to the rollup so that they can be interpreted as a user
/// @notice In the upgrade executor we need to access functions on the rollup
///         but since the upgrade executor is the admin it will always be forwarded to the
///         rollup admin logic. We create a separate forwarder contract here that just relays
///         information, since it's not the admin it can access rollup user logic.
contract RollupReader is IOldRollup {
    IOldRollup public immutable rollup;

    constructor(IOldRollup _rollup) {
        rollup = _rollup;
    }

    function wasmModuleRoot() external view returns (bytes32) {
        return rollup.wasmModuleRoot();
    }

    function latestConfirmed() external view returns (uint64) {
        return rollup.latestConfirmed();
    }

    function getNode(uint64 nodeNum) external view returns (Node memory) {
        return rollup.getNode(nodeNum);
    }

    function getStakerAddress(uint64 stakerNum) external view returns (address) {
        return rollup.getStakerAddress(stakerNum);
    }

    function stakerCount() external view returns (uint64) {
        return rollup.stakerCount();
    }

    function getStaker(address staker) external view returns (OldStaker memory) {
        return rollup.getStaker(staker);
    }
}

interface IAddressRegistry {
    // CHRIS: TODO: consider creating a generic cross chain receiver contract for the escrow funds
    function l1Timelock() external returns (address);
    function rollup() external returns (IOldRollup);
    function bridge() external returns (address);
    function sequencerInbox() external returns (address);
    function rollupEventInbox() external returns (address);
    function outbox() external returns (address);
    function inbox() external returns (address);
}

/// @title  Upgrades an Arbitrum rollup to the new challenge protocol
/// @notice Requires implementation contracts to be pre-deployed and provided in the constructor
///         Also requires a lookup contract to be provided that contains the pre-image of the state hash
///         that is in the latest confirmed assertion in the current rollup.
contract BOLDUpgradeAction {
    bytes32 public constant _ADMIN_SLOT = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
    uint256 public constant BLOCK_LEAF_SIZE = 2 ^ 26; // CHRIS: TODO: get final number for this
    uint256 public constant BIGSTEP_LEAF_SIZE = 2 ^ 23;
    uint256 public constant SMALLSTEP_LEAF_SIZE = 2 ^ 20;

    IAddressRegistry public immutable ADDRESS_REG;
    
    uint64 public immutable CONFIRM_PERIOD_BLOCKS;
    address public immutable STAKE_TOKEN;
    uint256 public immutable STAKE_AMOUNT;
    uint256 public immutable MINI_STAKE_AMOUNT;
    uint256 public immutable CHAIN_ID;

    IOneStepProofEntry public immutable OSP;
    // proxy admins of the contracts to be upgraded
    ProxyAdmin public immutable PROXY_ADMIN_OUTBOX;
    ProxyAdmin public immutable PROXY_ADMIN_BRIDGE;
    ProxyAdmin public immutable PROXY_ADMIN_REI;
    ProxyAdmin public immutable PROXY_ADMIN_SEQUENCER_INBOX;
    StateHashPreImageLookup public immutable PREIMAGE_LOOKUP;
    RollupReader public immutable ROLLUP_READER;

    // new contract implementations
    address public immutable IMPL_BRIDGE;
    address public immutable IMPL_SEQUENCER_INBOX;
    address public immutable IMPL_REI;
    address public immutable IMPL_OUTBOX;
    address public immutable IMPL_OLD_ROLLUP_USER;
    address public immutable IMPL_NEW_ROLLUP_USER;
    address public immutable IMPL_NEW_ROLLUP_ADMIN;
    address public immutable IMPL_CHALLENGE_MANAGER;

    event RollupMigrated(address rollup, address challengeManager);

    struct Settings {
        uint64 confirmPeriodBlocks;
        address stakeToken;
        uint256 stakeAmt;
        uint256 miniStakeAmt;
        uint256 chainId;
    }

    // Unfortunately these are not discoverable on-chain, so we need to supply them
    struct ProxyAdmins {
        address outbox;
        address bridge;
        address rei;
        address seqInbox;
    }

    struct Implementations {
        address bridge;
        address seqInbox;
        address rei;
        address outbox;
        address oldRollupUser;
        address newRollupUser;
        address newRollupAdmin;
        address challengeManager;
    }

    constructor(
        IAddressRegistry addressReg,
        IOneStepProofEntry osp,
        ProxyAdmins memory proxyAdmins,
        StateHashPreImageLookup lookup,
        RollupReader rollupReader,
        Implementations memory implementations,
        Settings memory settings
    ) {
        ADDRESS_REG = addressReg;
        OSP = osp;
        PROXY_ADMIN_OUTBOX = ProxyAdmin(proxyAdmins.outbox);
        PROXY_ADMIN_BRIDGE = ProxyAdmin(proxyAdmins.bridge);
        PROXY_ADMIN_REI = ProxyAdmin(proxyAdmins.rei);
        PROXY_ADMIN_SEQUENCER_INBOX = ProxyAdmin(proxyAdmins.seqInbox);
        PREIMAGE_LOOKUP = lookup;
        ROLLUP_READER = rollupReader;

        IMPL_BRIDGE = implementations.bridge;
        IMPL_SEQUENCER_INBOX = implementations.seqInbox;
        IMPL_REI = implementations.rei;
        IMPL_OUTBOX = implementations.outbox;
        IMPL_OLD_ROLLUP_USER = implementations.oldRollupUser;
        IMPL_NEW_ROLLUP_USER = implementations.newRollupUser;
        IMPL_NEW_ROLLUP_ADMIN = implementations.newRollupAdmin;
        IMPL_CHALLENGE_MANAGER = implementations.challengeManager;

        CHAIN_ID = settings.chainId;
        CONFIRM_PERIOD_BLOCKS = settings.confirmPeriodBlocks;
        STAKE_TOKEN = settings.stakeToken;
        STAKE_AMOUNT = settings.stakeAmt;
        MINI_STAKE_AMOUNT = settings.miniStakeAmt;
    }

    /// @dev    Refund the existing stakers, pause and upgrade the current rollup to
    ///         allow them to withdraw after pausing
    function cleanupOldRollup() private {
        IOldRollup oldRollup = ADDRESS_REG.rollup();
        IOldRollupAdmin(address(oldRollup)).pause();

        uint64 stakerCount = ROLLUP_READER.stakerCount();
        // since we for-loop these stakers we set an arbitrary limit - we dont
        // expect any instances to have close to this number of stakers
        if (stakerCount > 50) {
            stakerCount = 50;
        }
        for (uint64 i = 0; i < stakerCount; i++) {
            address stakerAddr = ROLLUP_READER.getStakerAddress(i);
            OldStaker memory staker = ROLLUP_READER.getStaker(stakerAddr);
            if (staker.isStaked && staker.currentChallenge == 0) {
                address[] memory stakersToRefund = new address[](1);
                stakersToRefund[0] = stakerAddr;

                IOldRollupAdmin(address(oldRollup)).forceRefundStaker(stakersToRefund);
            }
        }

        // upgrade the rollup to one that allows validators to withdraw even whilst paused
        UUPSUpgradeable(address(oldRollup)).upgradeTo(IMPL_OLD_ROLLUP_USER);
    }

    /// @dev    Create a config for the new rollup - fetches the latest confirmed
    ///         assertion from the old rollup and uses it as genesis
    function createConfig() private returns (Config memory) {
        // fetch the assertion associated with the latest confirmed state
        bytes32 latestConfirmedStateHash = ROLLUP_READER.getNode(ROLLUP_READER.latestConfirmed()).stateHash;
        (ExecutionState memory genesisExecState, uint256 inboxMaxCount) = PREIMAGE_LOOKUP.get(latestConfirmedStateHash);
        // double check the hash
        require(
            RollupLib.stateHashMem(genesisExecState, inboxMaxCount) == latestConfirmedStateHash,
            "Invalid latest execution hash"
        );

        ISequencerInbox.MaxTimeVariation memory maxTimeVariation; // can be empty as it's not used in rollup creation
        return Config({
            confirmPeriodBlocks: CONFIRM_PERIOD_BLOCKS,
            stakeToken: STAKE_TOKEN,
            baseStake: STAKE_AMOUNT,
            wasmModuleRoot: ROLLUP_READER.wasmModuleRoot(),
            owner: address(this), // upgrade executor is the owner
            loserStakeEscrow: ADDRESS_REG.l1Timelock(), // additional funds get sent to the l1 timelock
            chainId: CHAIN_ID,
            miniStakeValue: MINI_STAKE_AMOUNT,
            sequencerInboxMaxTimeVariation: maxTimeVariation,
            layerZeroBlockEdgeHeight: BLOCK_LEAF_SIZE,
            layerZeroBigStepEdgeHeight: BIGSTEP_LEAF_SIZE,
            layerZeroSmallStepEdgeHeight: SMALLSTEP_LEAF_SIZE,
            genesisExecutionState: genesisExecState,
            genesisInboxCount: inboxMaxCount
        });
    }

    function upgradeSurroundingContracts(address newRollupAddress) private {
        // now we upgrade each of the contracts that a reference to the rollup address
        // first we upgrade to an implementation which allows setting, then set the rollup address
        // then we revert to the previous implementation since we dont require this functionality going forward
        TransparentUpgradeableProxy bridge = TransparentUpgradeableProxy(payable(ADDRESS_REG.bridge()));
        address currentBridgeImpl = bridge.implementation();
        PROXY_ADMIN_BRIDGE.upgradeAndCall(
            bridge, IMPL_BRIDGE, abi.encodeWithSelector(IBridge.updateRollupAddress.selector, newRollupAddress)
        );
        PROXY_ADMIN_BRIDGE.upgrade(bridge, currentBridgeImpl);

        TransparentUpgradeableProxy sequencerInbox = TransparentUpgradeableProxy(payable(ADDRESS_REG.sequencerInbox()));
        address currentSequencerInboxImpl = sequencerInbox.implementation();
        PROXY_ADMIN_SEQUENCER_INBOX.upgradeAndCall(
            sequencerInbox, IMPL_SEQUENCER_INBOX, abi.encodeWithSelector(IOutbox.updateRollupAddress.selector)
        );
        PROXY_ADMIN_SEQUENCER_INBOX.upgrade(sequencerInbox, currentSequencerInboxImpl);

        TransparentUpgradeableProxy rollupEventInbox =
            TransparentUpgradeableProxy(payable(ADDRESS_REG.rollupEventInbox()));
        address currentRollupEventInboxImpl = rollupEventInbox.implementation();
        PROXY_ADMIN_REI.upgradeAndCall(
            rollupEventInbox, IMPL_REI, abi.encodeWithSelector(IOutbox.updateRollupAddress.selector)
        );
        PROXY_ADMIN_REI.upgrade(rollupEventInbox, currentRollupEventInboxImpl);

        TransparentUpgradeableProxy outbox = TransparentUpgradeableProxy(payable(ADDRESS_REG.outbox()));
        address currentOutboxImpl = outbox.implementation();
        PROXY_ADMIN_OUTBOX.upgradeAndCall(
            outbox, IMPL_OUTBOX, abi.encodeWithSelector(IOutbox.updateRollupAddress.selector)
        );
        PROXY_ADMIN_OUTBOX.upgrade(outbox, currentOutboxImpl);
    }

    function perform() external {
        // tidy up the old rollup - pause it and refund stakes
        cleanupOldRollup();

        // create the config, we do this now so that we compute the expected rollup address
        Config memory config = createConfig();

        // upgrade the surrounding contracts eg bridge, outbox, seq inbox, rollup event inbox
        // to set of the new rollup address
        bytes32 rollupSalt = keccak256(abi.encode(config));
        // CHRIS: TODO: as it stands we have the address wrong here since we dont append params to the creation code
        //              however in nitro we've moved away from this and have an initializer
        //              So this line and the new RollupProxy below need to be updated after updating from nitro
        address expectedRollupAddress =
            Create2Upgradeable.computeAddress(rollupSalt, keccak256(type(RollupProxy).creationCode));
        upgradeSurroundingContracts(expectedRollupAddress);

        // deploy the new challenge manager
        IEdgeChallengeManager challengeManager = IEdgeChallengeManager(
            address(
                new TransparentUpgradeableProxy(
                    address(IMPL_CHALLENGE_MANAGER),
                    address(PROXY_ADMIN_BRIDGE), // use the same proxy admin as the bridge
                    ""
                )
            )
        );
        challengeManager.initialize({
            _assertionChain: IAssertionChain(expectedRollupAddress),
            // confirm period and challenge period are the same atm
            _challengePeriodBlocks: config.confirmPeriodBlocks,
            _oneStepProofEntry: OSP,
            layerZeroBlockEdgeHeight: config.layerZeroBlockEdgeHeight,
            layerZeroBigStepEdgeHeight: config.layerZeroBigStepEdgeHeight,
            layerZeroSmallStepEdgeHeight: config.layerZeroSmallStepEdgeHeight,
            _stakeToken: IERC20(config.stakeToken),
            _stakeAmount: config.miniStakeValue,
            _excessStakeReceiver: ADDRESS_REG.l1Timelock()
        });

        // now that all the dependent contracts are pointed at the new address we can
        // deploy and init the new rollup
        ContractDependencies memory connectedContracts = ContractDependencies({
            bridge: IBridge(ADDRESS_REG.bridge()),
            sequencerInbox: ISequencerInbox(ADDRESS_REG.sequencerInbox()),
            inbox: IInbox(ADDRESS_REG.inbox()),
            outbox: IOutbox(ADDRESS_REG.outbox()),
            rollupEventInbox: IRollupEventInbox(ADDRESS_REG.rollupEventInbox()),
            challengeManager: challengeManager,
            rollupAdminLogic: IMPL_NEW_ROLLUP_ADMIN,
            rollupUserLogic: IRollupUser(IMPL_NEW_ROLLUP_USER),
            validatorUtils: address(0), // CHRIS: TODO: remove this from the admin contract
            validatorWalletCreator: address(0) // CHRIS: TODO: remove this from the admin contract
        });

        RollupProxy rollup = new RollupProxy{ salt: rollupSalt}(
            config, connectedContracts
        );
        require(address(rollup) == expectedRollupAddress, "UNEXPCTED_ROLLUP_ADDR");

        emit RollupMigrated(expectedRollupAddress, address(challengeManager));
    }
}
