// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "../state/GlobalState.sol";
import "../state/Machine.sol";
import "../osp/IOneStepProofEntry.sol";

struct AssertionNode {
    // Hash of the state of the chain as of this assertion
    bytes32 stateHash;
    // Hash of the data that will be committed if this assertion is confirmed
    bytes32 confirmData;
    // Index of the assertion previous to this one
    uint64 prevNum;
    // Deadline at which this assertion can be confirmed
    uint64 deadlineBlock;
    // Deadline at which a child of this assertion can be confirmed
    uint64 noChildConfirmedBeforeBlock;
    // Number of stakers staked on this assertion. This includes real stakers and zombies
    uint64 stakerCount;
    // Number of stakers staked on a child assertion. This includes real stakers and zombies
    uint64 childStakerCount;
    // This value starts at zero and is set to a value when the first child is created. After that it is constant until the assertion is destroyed or the owner destroys pending assertions
    uint64 firstChildBlock;
    uint256 firstChildTime; // TODO: remove this after migrating to use block instead of timestamp
    // This value starts at zero and is set to a value when the second child is created. After that it is constant until the assertion is destroyed or the owner destroys pending assertions
    uint64 secondChildBlock;
    // The block number when this assertion was created
    uint64 createdAtBlock;
    // A hash of all the data needed to determine this assertion's validity, to protect against reorgs
    bytes32 assertionHash;

    bool isFirstChild; // no longer in assertionHash
}

struct AssertionInputs {
    ExecutionState beforeState;
    ExecutionState afterState;
}

/**
 * @notice Utility functions for Assertion
 */
library AssertionNodeLib {
    /**
     * @notice Initialize a Assertion
     * @param _stateHash Initial value of stateHash
     * @param _confirmData Initial value of confirmData
     * @param _prevNum Initial value of prevNum
     * @param _deadlineBlock Initial value of deadlineBlock
     * @param _assertionHash Initial value of assertionHash
     */
    function createAssertion(
        bytes32 _stateHash,
        bytes32 _confirmData,
        uint64 _prevNum,
        uint64 _deadlineBlock,
        bytes32 _assertionHash,
        bool _isFirstChild
    ) internal view returns (AssertionNode memory) {
        AssertionNode memory assertion;
        assertion.stateHash = _stateHash;
        assertion.confirmData = _confirmData;
        assertion.prevNum = _prevNum;
        assertion.deadlineBlock = _deadlineBlock;
        assertion.noChildConfirmedBeforeBlock = _deadlineBlock;
        assertion.createdAtBlock = uint64(block.number);
        assertion.assertionHash = _assertionHash;
        assertion.isFirstChild = _isFirstChild;
        return assertion;
    }

    /**
     * @notice Update child properties
     * @param number The child number to set
     */
    function childCreated(AssertionNode storage self, uint64 number, uint64 confirmPeriodBlocks) internal {
        if (self.firstChildBlock == 0) {
            self.firstChildBlock = uint64(block.number);
            self.firstChildTime = block.timestamp;
            self.noChildConfirmedBeforeBlock = uint64(block.number) + confirmPeriodBlocks;
        } else if (self.secondChildBlock == 0) {
            self.secondChildBlock = uint64(block.number);
        }
    }

    /**
     * @notice Update the child confirmed deadline
     * @param deadline The new deadline to set
     */
    function newChildConfirmDeadline(AssertionNode storage self, uint64 deadline) internal {
        self.noChildConfirmedBeforeBlock = deadline;
    }

    /**
     * @notice Check whether the current block number has met or passed the assertion's deadline
     */
    function requirePastDeadline(AssertionNode memory self) internal view {
        require(block.number >= self.deadlineBlock, "BEFORE_DEADLINE");
    }

    /**
     * @notice Check whether the current block number has met or passed deadline for children of this assertion to be confirmed
     */
    function requirePastChildConfirmDeadline(AssertionNode memory self) internal view {
        require(block.number >= self.noChildConfirmedBeforeBlock, "CHILD_TOO_RECENT");
    }
}
