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
    uint256 public chainId;

    // These 4 config should be stored into the prev and not used directly
    // An assertion can be confirmed after confirmPeriodBlocks when it is unchallenged
    uint64 public confirmPeriodBlocks;
    uint256 public baseStake;
    bytes32 public wasmModuleRoot;
    // When there is a challenge, we trust the challenge manager to determine the winner
    IEdgeChallengeManager public challengeManager;

    IInbox public inbox;
    IBridge public bridge;
    IOutbox public outbox;
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

    function sequencerInbox() public view virtual returns (ISequencerInbox) {
        return ISequencerInbox(bridge.sequencerInbox());
    }

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
        _assertions[assertionHash] = initialAssertion;
        _latestConfirmed = assertionHash;
    }

    /**
     * @dev This function will validate the parentAssertionHash, confirmState and inboxAcc against the assertionId
     *          and check if the assertionId is currently pending. If all checks pass, the assertion will be confirmed.
     */
    function confirmAssertionInternal(
        bytes32 assertionId,
        bytes32 parentAssertionHash,
        ExecutionState calldata confirmState,
        bytes32 inboxAcc
    ) internal {
        AssertionNode storage assertion = getAssertionStorage(assertionId);
        // Check that assertion is pending, this also checks that assertion exists
        require(assertion.status == AssertionStatus.Pending, "NOT_PENDING");

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
     * This should only be called when the staker is inactive
     * @param stakerAddress Address of the staker withdrawing their stake
     */
    function withdrawStaker(address stakerAddress) internal {
        Staker storage staker = _stakerMap[stakerAddress];
        uint256 initialStaked = staker.amountStaked;
        increaseWithdrawableFunds(stakerAddress, initialStaked);
        deleteStaker(stakerAddress);
        emit UserStakeUpdated(stakerAddress, initialStaked, 0);
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

    function createNewAssertion(
        AssertionInputs calldata assertion,
        bytes32 prevAssertionId,
        bytes32 expectedAssertionHash
    ) internal returns (bytes32) {
        // Validate the config hash
        RollupLib.validateConfigHash(
            assertion.beforeStateData.configData, getAssertionStorage(prevAssertionId).configHash
        );

        // reading inbox messages always terminates in either a finished or errored state
        // although the challenge protocol that any invalid terminal state will be proven incorrect
        // we can do a quick sanity check here
        require(
            assertion.afterState.machineStatus == MachineStatus.FINISHED
                || assertion.afterState.machineStatus == MachineStatus.ERRORED,
            "BAD_AFTER_STATUS"
        );

        // validate the provided before state is correct by checking that it's part of the prev assertion hash
        require(
            RollupLib.assertionHash(
                assertion.beforeStateData.prevPrevAssertionHash,
                assertion.beforeState,
                assertion.beforeStateData.sequencerBatchAcc
            ) == prevAssertionId,
            "INVALID_BEFORE_STATE"
        );

        // The rollup cannot advance from an errored state
        // If it reaches an errored state it must be corrected by an administrator
        // This will involve updating the wasm root and creating an alternative assertion
        // that consumes the correct number of inbox messages, and correctly transitions to the
        // FINISHED state so that normal progress can continue
        require(assertion.beforeState.machineStatus == MachineStatus.FINISHED, "BAD_PREV_STATUS");

        AssertionNode storage prevAssertion = getAssertionStorage(prevAssertionId);
        uint256 nextInboxPosition;
        bytes32 sequencerBatchAcc;
        {
            uint64 afterInboxPosition = assertion.afterState.globalState.getInboxPosition();
            uint64 prevInboxPosition = assertion.beforeState.globalState.getInboxPosition();
            require(afterInboxPosition >= prevInboxPosition, "INBOX_BACKWARDS");
            if (assertion.afterState.machineStatus == MachineStatus.ERRORED) {
                // the errored position must still be within the correct message bounds
                require(
                    afterInboxPosition <= assertion.beforeStateData.configData.nextInboxPosition,
                    "ERRORED_INBOX_TOO_FAR"
                );

                // and cannot go backwards
                require(afterInboxPosition >= prevInboxPosition, "ERRORED_INBOX_TOO_FEW");

                // See validator/assertion.go ExecutionState RequiredBatches() for
                // for why we move forward in the batch when the machine ends in an errored state
                // CHRIS: TODO: remove this
                afterInboxPosition++;
            } else if (assertion.afterState.machineStatus == MachineStatus.FINISHED) {
                // Assertions must consume exactly all inbox messages
                // that were in the inbox at the time the previous assertion was created
                require(
                    assertion.afterState.globalState.getInboxPosition()
                        == assertion.beforeStateData.configData.nextInboxPosition,
                    "INCORRECT_INBOX_POS"
                );
                // Assertions that finish correctly completely consume the message
                // Therefore their position in the message is 0
                require(assertion.afterState.globalState.getPositionInMessage() == 0, "FINISHED_NON_ZERO_POS");

                // We enforce that at least one inbox message is always consumed
                // so the after inbox position is always strictly greater than previous
                require(afterInboxPosition > prevInboxPosition, "INBOX_BACKWARDS");
            }

            uint256 currentInboxPosition = bridge.sequencerMessageCount();
            // Cannot read more messages than currently exist in the inbox
            require(afterInboxPosition <= currentInboxPosition, "INBOX_PAST_END");

            // The next assertion must consume all the messages that are currently found in the inbox
            if (assertion.afterState.globalState.getInboxPosition() == currentInboxPosition) {
                // No new messages have been added to the inbox since the last assertion
                // In this case if we set the next inbox position to the current one we would be insisting that
                // the next assertion process no messages. So instead we increment the next inbox position to current
                // plus one, so that the next assertion will process exactly one message
                nextInboxPosition = currentInboxPosition + 1;
            } else {
                nextInboxPosition = currentInboxPosition;
            }

            // only the genesis assertion processes no messages, and that assertion is created
            // when we initialize this contract. Therefore, all assertions created here should have a non
            // zero inbox position.
            require(afterInboxPosition != 0, "EMPTY_INBOX_COUNT");

            // Fetch the inbox accumulator for this message count. Fetching this and checking against it
            // allows the assertion creator to ensure they're creating an assertion against the expected
            // inbox messages
            sequencerBatchAcc = bridge.sequencerInboxAccs(afterInboxPosition - 1);
        }

        bytes32 newAssertionHash = RollupLib.assertionHash(prevAssertionId, assertion.afterState, sequencerBatchAcc);

        // allow an assertion creator to ensure that they're creating their assertion against the expected state
        require(
            newAssertionHash == expectedAssertionHash || expectedAssertionHash == bytes32(0),
            "UNEXPECTED_ASSERTION_HASH"
        );

        // the assertion hash is unique - it's only possible to have one correct assertion hash
        // per assertion. Therefore we can check if this assertion has already been made, and if so
        // we can revert
        require(getAssertionStorage(newAssertionHash).status == AssertionStatus.NoAssertion, "ASSERTION_SEEN");

        // state updates
        AssertionNode memory newAssertion = AssertionNodeLib.createAssertion(
            prevAssertion.firstChildBlock == 0, // assumes block 0 is impossible
            RollupLib.configHash({
                wasmModuleRoot: wasmModuleRoot,
                requiredStake: baseStake,
                challengeManager: address(challengeManager),
                confirmPeriodBlocks: confirmPeriodBlocks,
                nextInboxPosition: uint64(nextInboxPosition)
            })
        );

        // Fetch a storage reference to prevAssertion since we copied our other one into memory
        // and we don't have enough stack available to keep to keep the previous storage reference around
        prevAssertion.childCreated();
        _assertions[newAssertionHash] = newAssertion;

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

    function genesisAssertionId() external pure returns (bytes32) {
        GlobalState memory emptyGlobalState;
        ExecutionState memory emptyExecutionState = ExecutionState(emptyGlobalState, MachineStatus.FINISHED);
        bytes32 parentAssertionHash = bytes32(0);
        bytes32 inboxAcc = bytes32(0);
        return RollupLib.assertionHash({
            parentAssertionHash: parentAssertionHash,
            afterState: emptyExecutionState,
            inboxAcc: inboxAcc
        });
    }

    function getFirstChildCreationBlock(bytes32 assertionId) external view returns (uint256) {
        return getAssertionStorage(assertionId).firstChildBlock;
    }

    function getSecondChildCreationBlock(bytes32 assertionId) external view returns (uint256) {
        return getAssertionStorage(assertionId).secondChildBlock;
    }

    function validateAssertionId(
        bytes32 assertionId,
        ExecutionState calldata state,
        bytes32 prevAssertionId,
        bytes32 inboxAcc
    ) external pure {
        require(assertionId == RollupLib.assertionHash(prevAssertionId, state, inboxAcc), "INVALID_ASSERTION_HASH");
    }

    function validateConfig(bytes32 assertionId, ConfigData calldata configData) external view {
        RollupLib.validateConfigHash(configData, getAssertionStorage(assertionId).configHash);
    }

    function isFirstChild(bytes32 assertionId) external view returns (bool) {
        return getAssertionStorage(assertionId).isFirstChild;
    }

    function isPending(bytes32 assertionId) external view returns (bool) {
        return getAssertionStorage(assertionId).status == AssertionStatus.Pending;
    }

    /**
     * @notice Verify that the given staker is not active
     * @param stakerAddress Address to check
     */
    function requireInactiveStaker(address stakerAddress) internal view {
        require(isStaked(stakerAddress), "NOT_STAKED");
        // A staker is inactive if
        // a) their last staked assertion is the latest confirmed assertion
        // b) their last staked assertion have a child
        bytes32 lastestAssertion = latestStakedAssertion(stakerAddress);
        bool isLatestConfirmed = lastestAssertion == latestConfirmed();
        bool haveChild = getAssertionStorage(lastestAssertion).firstChildBlock > 0;
        require(isLatestConfirmed || haveChild, "STAKE_ACTIVE");
    }
}
