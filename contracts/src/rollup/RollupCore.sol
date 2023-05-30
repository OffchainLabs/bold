// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";

import "./Assertion.sol";
import "./RollupLib.sol";
import "./IRollupEventInbox.sol";
import "./IRollupCore.sol";

import "../state/Machine.sol";

import "../bridge/ISequencerInbox.sol";
import "../bridge/IBridge.sol";
import "../bridge/IOutbox.sol";
import "../challengeV2/EdgeChallengeManager.sol";
import {NO_CHAL_INDEX} from "../libraries/Constants.sol";

abstract contract RollupCore is IRollupCore, PausableUpgradeable {
    using AssertionNodeLib for AssertionNode;
    using GlobalStateLib for GlobalState;

    // Rollup Config
    uint64 public confirmPeriodBlocks;
    uint64 public extraChallengeTimeBlocks; // TODO: unused
    uint256 public chainId;
    uint256 public baseStake;
    bytes32 public wasmModuleRoot;

    IInbox public inbox;
    IBridge public bridge;
    IOutbox public outbox;
    ISequencerInbox public sequencerInbox;
    IRollupEventInbox public rollupEventInbox;

    // misc useful contracts when interacting with the rollup
    address public validatorUtils;
    address public validatorWalletCreator;

    // only 1 child can be confirmed, the excess/loser stake will be sent to this address
    address public loserStakeEscrow;
    address public stakeToken;
    uint256 public minimumAssertionPeriod;

    mapping(address => bool) public isValidator;

    bytes32 private _latestConfirmed;
    mapping(bytes32 => AssertionNode) private _assertions;

    address[] private _stakerList;
    mapping(address => Staker) public _stakerMap;

    mapping(address => uint256) private _withdrawableFunds;
    uint256 public totalWithdrawableFunds;
    uint256 public rollupDeploymentBlock;

    // The assertion number of the initial assertion
    uint64 internal constant GENESIS_NODE = 1;

    bool public validatorWhitelistDisabled;

    IEdgeChallengeManager public challengeManager;

    /**
     * @notice Get a storage reference to the Assertion for the given assertion id
     * @dev The assertion may not exists
     * @param assertionId Id of the assertion
     * @return Assertion struct
     */
    function getAssertionStorage(bytes32 assertionId) internal view returns (AssertionNode storage) {
        require(assertionId != bytes32(0), "ASSERTION_ID_CANNOT_BE_ZERO");
        return _assertions[assertionId];
    }

    /**
     * @notice Get the Assertion for the given index.
     */
    function getAssertion(bytes32 assertionId) public view override returns (AssertionNode memory) {
        return getAssertionStorage(assertionId);
    }

    /**
     * @notice Get the address of the staker at the given index
     * @param stakerNum Index of the staker
     * @return Address of the staker
     */
    function getStakerAddress(uint64 stakerNum) external view override returns (address) {
        return _stakerList[stakerNum];
    }

    /**
     * @notice Check whether the given staker is staked
     * @param staker Staker address to check
     * @return True or False for whether the staker was staked
     */
    function isStaked(address staker) public view override returns (bool) {
        return _stakerMap[staker].isStaked;
    }

    /**
     * @notice Get the latest staked assertion of the given staker
     * @param staker Staker address to lookup
     * @return Latest assertion staked of the staker
     */
    function latestStakedAssertion(address staker) public view override returns (bytes32) {
        return _stakerMap[staker].latestStakedAssertion;
    }

    /**
     * @notice Get the amount staked of the given staker
     * @param staker Staker address to lookup
     * @return Amount staked of the staker
     */
    function amountStaked(address staker) public view override returns (uint256) {
        return _stakerMap[staker].amountStaked;
    }

    /**
     * @notice Retrieves stored information about a requested staker
     * @param staker Staker address to retrieve
     * @return A structure with information about the requested staker
     */
    function getStaker(address staker) external view override returns (Staker memory) {
        return _stakerMap[staker];
    }

    /**
     * @notice Get the amount of funds withdrawable by the given address
     * @param user Address to check the funds of
     * @return Amount of funds withdrawable by user
     */
    function withdrawableFunds(address user) external view override returns (uint256) {
        return _withdrawableFunds[user];
    }

    /// @return Index of the latest confirmed assertion
    function latestConfirmed() public view override returns (bytes32) {
        return _latestConfirmed;
    }

    /// @return Number of active stakers currently staked
    function stakerCount() public view override returns (uint64) {
        return uint64(_stakerList.length);
    }

    /**
     * @notice Initialize the core with an initial assertion
     * @param initialAssertion Initial assertion to start the chain with
     */
    function initializeCore(AssertionNode memory initialAssertion, bytes32 assertionHash) internal {
        __Pausable_init();
        // TODO: HN: prolly should use the internal function to create genesis
        _assertions[assertionHash] = initialAssertion;
        _latestConfirmed = assertionHash;
    }

    /**
     * @notice React to a new assertion being created by storing it an incrementing the latest assertion counter
     * @param assertion Assertion that was newly created
     */
    function assertionCreated(AssertionNode memory assertion, bytes32 assertionHash) internal {
        _assertions[assertionHash] = assertion;
    }

    function confirmAssertion(
        bytes32 assertionId,
        bytes32 parentAssertionHash,
        ExecutionState calldata confirmState,
        bytes32 inboxAcc
    ) internal {
        AssertionNode storage assertion = getAssertionStorage(assertionId);

        // Authenticate data against assertionHash pre-image
        require(
            assertionId
                == RollupLib.assertionHash({
                    parentAssertionHash: parentAssertionHash,
                    afterState: confirmState,
                    inboxAcc: inboxAcc
                }),
            "CONFIRM_DATA"
        );

        bytes32 blockHash = confirmState.globalState.bytes32Vals[0];
        bytes32 sendRoot = confirmState.globalState.bytes32Vals[1];

        // trusted external call to outbox
        outbox.updateSendRoot(sendRoot, blockHash);

        _latestConfirmed = assertionId;
        assertion.status = AssertionStatus.Confirmed;

        emit AssertionConfirmed(assertionId, blockHash, sendRoot);
    }

    /**
     * @notice Create a new stake at latest confirmed assertion
     * @param stakerAddress Address of the new staker
     * @param depositAmount Stake amount of the new staker
     */
    function createNewStake(address stakerAddress, uint256 depositAmount) internal {
        uint64 stakerIndex = uint64(_stakerList.length);
        _stakerList.push(stakerAddress);
        _stakerMap[stakerAddress] = Staker(depositAmount, _latestConfirmed, stakerIndex, true);
        emit UserStakeUpdated(stakerAddress, 0, depositAmount);
    }

    /**
     * @notice Add to the stake of the given staker by the given amount
     * @param stakerAddress Address of the staker to increase the stake of
     * @param amountAdded Amount of stake to add to the staker
     */
    function increaseStakeBy(address stakerAddress, uint256 amountAdded) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        uint256 initialStaked = staker.amountStaked;
        uint256 finalStaked = initialStaked + amountAdded;
        staker.amountStaked = finalStaked;
        emit UserStakeUpdated(stakerAddress, initialStaked, finalStaked);
    }

    /**
     * @notice Reduce the stake of the given staker to the given target
     * @param stakerAddress Address of the staker to reduce the stake of
     * @param target Amount of stake to leave with the staker
     * @return Amount of value released from the stake
     */
    function reduceStakeTo(address stakerAddress, uint256 target) internal returns (uint256) {
        Staker storage staker = _stakerMap[stakerAddress];
        uint256 current = staker.amountStaked;
        require(target <= current, "TOO_LITTLE_STAKE");
        uint256 amountWithdrawn = current - target;
        staker.amountStaked = target;
        increaseWithdrawableFunds(stakerAddress, amountWithdrawn);
        emit UserStakeUpdated(stakerAddress, current, target);
        return amountWithdrawn;
    }

    /**
     * @notice Remove the given staker and return their stake
     * This should not be called if the staker is staked on a descendent of the latest confirmed assertion
     * @param stakerAddress Address of the staker withdrawing their stake
     */
    function withdrawStaker(address stakerAddress) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        // TODO: HN: review if we need additional checks here
        //           the user logic already checked if the staker is inactive
        uint256 initialStaked = staker.amountStaked;
        increaseWithdrawableFunds(stakerAddress, initialStaked);
        deleteStaker(stakerAddress);
        emit UserStakeUpdated(stakerAddress, initialStaked, 0);
    }

    /**
     * @notice Advance the given staker to the given assertion
     * @param stakerAddress Address of the staker adding their stake
     * @param assertionId Id of the assertion to stake on
     */
    function stakeOnAssertion(address stakerAddress, bytes32 assertionId) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        staker.latestStakedAssertion = assertionId;
    }

    /**
     * @notice Clear the withdrawable funds for the given address
     * @param account Address of the account to remove funds from
     * @return Amount of funds removed from account
     */
    function withdrawFunds(address account) internal returns (uint256) {
        uint256 amount = _withdrawableFunds[account];
        _withdrawableFunds[account] = 0;
        totalWithdrawableFunds -= amount;
        emit UserWithdrawableFundsUpdated(account, amount, 0);
        return amount;
    }

    /**
     * @notice Increase the withdrawable funds for the given address
     * @param account Address of the account to add withdrawable funds to
     */
    function increaseWithdrawableFunds(address account, uint256 amount) internal {
        uint256 initialWithdrawable = _withdrawableFunds[account];
        uint256 finalWithdrawable = initialWithdrawable + amount;
        _withdrawableFunds[account] = finalWithdrawable;
        totalWithdrawableFunds += amount;
        emit UserWithdrawableFundsUpdated(account, initialWithdrawable, finalWithdrawable);
    }

    /**
     * @notice Remove the given staker
     * @param stakerAddress Address of the staker to remove
     */
    function deleteStaker(address stakerAddress) private {
        Staker storage staker = _stakerMap[stakerAddress];
        require(staker.isStaked, "NOT_STAKED");
        uint64 stakerIndex = staker.index;
        _stakerList[stakerIndex] = _stakerList[_stakerList.length - 1];
        _stakerMap[_stakerList[stakerIndex]].index = stakerIndex;
        _stakerList.pop();
        delete _stakerMap[stakerAddress];
    }

    struct StakeOnNewAssertionFrame {
        uint256 currentInboxSize;
        AssertionNode assertion;
        bytes32 stateHash;
        AssertionNode prevAssertion;
        bytes32 lastHash;
        bool hasSibling;
        uint64 deadlineBlock;
        bytes32 sequencerBatchAcc;
    }

    function createNewAssertion(
        AssertionInputs calldata assertion,
        bytes32 prevAssertionId,
        bytes32 expectedAssertionHash
    ) internal returns (bytes32) {
        require(
            assertion.afterState.machineStatus == MachineStatus.FINISHED
                || assertion.afterState.machineStatus == MachineStatus.ERRORED,
            "BAD_AFTER_STATUS"
        );

        AssertionNode storage prevAssertion = getAssertionStorage(prevAssertionId);
        // validate the before state
        require(
            RollupLib.assertionHash(
                assertion.beforeStateData.prevAssertionHash,
                assertion.beforeState,
                assertion.beforeStateData.sequencerBatchAcc
            ) == prevAssertionId,
            "INVALID_BEFORE_STATE"
        );

        uint256 nextInboxPosition;
        bytes32 sequencerBatchAcc;
        {
            // Validate the inbox positions
            uint64 afterInboxCount = assertion.afterState.globalState.getInboxPosition();
            uint64 prevInboxPosition = assertion.beforeState.globalState.getInboxPosition();
            require(afterInboxCount >= prevInboxPosition, "INBOX_BACKWARDS");
            if (afterInboxCount == prevInboxPosition) {
                require(
                    assertion.afterState.globalState.getPositionInMessage()
                        >= assertion.beforeState.globalState.getPositionInMessage(),
                    "INBOX_POS_IN_MSG_BACKWARDS"
                );
            }

            // See validator/assertion.go ExecutionState RequiredBatches() for reasoning
            if (
                assertion.afterState.machineStatus == MachineStatus.ERRORED
                    || assertion.afterState.globalState.getPositionInMessage() > 0
            ) {
                // The current inbox message was read
                afterInboxCount++;
            }
            // Cannot read more messages than currently exist
            uint256 currentInboxPosition = bridge.sequencerMessageCount();
            require(afterInboxCount <= currentInboxPosition, "INBOX_PAST_END");

            if (assertion.afterState.globalState.getInboxPosition() == currentInboxPosition) {
                // assertions must consume exactly up to the message count that was in the inbox
                // when the prev assertion was made. However if no new messages are sent, the next assertion
                // would need to consume the same number of messages as the prev, meaning the chain
                // would be unable to make progress. To avoid this we say that if no new messages have been
                // made between the prev and now, then the next assertion should consume one message
                nextInboxPosition = currentInboxPosition + 1;
            } else {
                nextInboxPosition = currentInboxPosition;
            }

            // we don't create an assertion until messages are added to the inbox
            require(afterInboxCount != 0, "EMPTY_INBOX_COUNT");

            // This gives replay protection against the state of the inbox
            sequencerBatchAcc = bridge.sequencerInboxAccs(afterInboxCount - 1);
        }

        bytes32 newAssertionHash = RollupLib.assertionHash(prevAssertionId, assertion.afterState, sequencerBatchAcc);
        require(
            newAssertionHash == expectedAssertionHash || expectedAssertionHash == bytes32(0), "UNEXPECTED_NODE_HASH"
        );

        require(!isAssertionExists(newAssertionHash), "ASSERTION_SEEN");

        AssertionNode memory newAssertion = AssertionNodeLib.createAssertion(
            uint64(nextInboxPosition),
            prevAssertionId,
            uint64(block.number) + confirmPeriodBlocks,
            prevAssertion.firstChildBlock == 0, // assume block 0 is impossible
            RollupLib.configHash({
                wasmModuleRoot: wasmModuleRoot,
                requiredStake: baseStake,
                challengeManager: address(challengeManager),
                confirmPeriodBlocks: confirmPeriodBlocks
            })
        );

        {
            // Fetch a storage reference to prevAssertion since we copied our other one into memory
            // and we don't have enough stack available to keep to keep the previous storage reference around
            prevAssertion.childCreated(confirmPeriodBlocks); // TODO: HN: this should use the prev's confirmPeriodBlocks
            assertionCreated(newAssertion, newAssertionHash);
        }

        emit AssertionCreated(
            newAssertionHash,
            prevAssertionId,
            assertion,
            sequencerBatchAcc,
            nextInboxPosition,
            wasmModuleRoot,
            baseStake,
            address(challengeManager),
            confirmPeriodBlocks
        );

        return newAssertionHash;
    }

    function getPredecessorId(bytes32 assertionId) external view returns (bytes32) {
        bytes32 prevId = getAssertionStorage(assertionId).prevId;
        return prevId;
    }

    function proveExecutionState(bytes32 assertionId, ExecutionState memory state, bytes memory proof)
        external
        pure
        returns (ExecutionState memory)
    {
        (bytes32 parentAssertionHash, bytes32 inboxAcc) = abi.decode(proof, (bytes32, bytes32));

        require(assertionId == RollupLib.assertionHash(parentAssertionHash, state, inboxAcc), "Invalid assertion hash");

        return state;
    }

    function getNextInboxPosition(bytes32 assertionId) external view returns (uint64) {
        return getAssertionStorage(assertionId).nextInboxPosition;
    }

    function hasSibling(bytes32 assertionId) external view returns (bool) {
        return getAssertionStorage(getAssertionStorage(assertionId).prevId).secondChildBlock != 0;
    }

    // HN: TODO: use block or timestamp?
    function getFirstChildCreationBlock(bytes32 assertionId) external view returns (uint256) {
        return getAssertionStorage(assertionId).firstChildBlock;
    }

    function getSecondChildCreationBlock(bytes32 assertionId) external view returns (uint256) {
        return getAssertionStorage(assertionId).secondChildBlock;
    }

    function proveWasmModuleRoot(bytes32 assertionId, bytes32 root, bytes memory proof)
        external
        view
        returns (bytes32)
    {
        (uint256 requiredStake, address _challengeManager, uint256 _confirmPeriodBlocks) =
            abi.decode(proof, (uint256, address, uint256));
        require(
            RollupLib.configHash({
                wasmModuleRoot: root,
                requiredStake: requiredStake,
                challengeManager: _challengeManager,
                confirmPeriodBlocks: _confirmPeriodBlocks
            }) == getAssertionStorage(assertionId).configHash,
            "BAD_WASM_MODULE_ROOT_PROOF"
        );
        return root;
    }

    function isFirstChild(bytes32 assertionId) external view returns (bool) {
        return getAssertionStorage(assertionId).isFirstChild;
    }

    function isPending(bytes32 assertionId) external view returns (bool) {
        return getAssertionStorage(assertionId).status == AssertionStatus.Pending;
    }

    function isAssertionExists(bytes32 id) public view returns (bool) {
        return _assertions[id].createdAtBlock > 0;
    }
}
