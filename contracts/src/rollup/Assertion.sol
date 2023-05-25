// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "../state/GlobalState.sol";
import "../state/Machine.sol";
import "../osp/IOneStepProofEntry.sol";

struct AssertionNode {
    // The inbox position that the assertion that succeeds should process up to and including
    uint64 nextInboxPosition;
    // Index of the assertion previous to this one
    uint64 prevNum;
    // Deadline at which this assertion can be confirmed
    uint64 deadlineBlock;
    // Deadline at which a child of this assertion can be confirmed
    uint64 noChildConfirmedBeforeBlock;
    // This value starts at zero and is set to a value when the first child is created. After that it is constant until the assertion is destroyed or the owner destroys pending assertions
    uint64 firstChildBlock;
    // This value starts at zero and is set to a value when the second child is created. After that it is constant until the assertion is destroyed or the owner destroys pending assertions
    uint64 secondChildBlock;
    // The block number when this assertion was created
    uint64 createdAtBlock;
    // True if this assertion is the first child of its prev
    bool isFirstChild;
    // A hash of all the data needed to determine this assertion's validity, to protect against reorgs
    bytes32 assertionHash;
    // A hash of all configuration data when the assertion is created
    bytes32 configHash;
}

struct BeforeStateData {
    bytes32 wasmRoot;
    bytes32 prevAssertionHash;
    bytes32 sequencerBatchAcc;
    uint256 requiredStake;
    address challengeManager;
    uint256 confirmPeriodBlocks;
}

struct AssertionInputs {
    // Additional data used to validate the before state
    BeforeStateData beforeStateData;
    ExecutionState beforeState;
    ExecutionState afterState;
}

/**
 * @notice Utility functions for Assertion
 */
library AssertionNodeLib {
    /**
     * @notice Initialize a Assertion
     * @param _nextInboxPosition The inbox position that the assertion that succeeds should process up to and including
     * @param _prevNum Initial value of prevNum
     * @param _deadlineBlock Initial value of deadlineBlock
     * @param _assertionHash Initial value of assertionHash
     */
    function createAssertion(
        uint64 _nextInboxPosition,
        uint64 _prevNum,
        uint64 _deadlineBlock,
        bytes32 _assertionHash,
        bool _isFirstChild,
        bytes32 _configHash
    ) internal view returns (AssertionNode memory) {
        AssertionNode memory assertion;
        assertion.nextInboxPosition = _nextInboxPosition;
        assertion.prevNum = _prevNum;
        assertion.deadlineBlock = _deadlineBlock;
        assertion.noChildConfirmedBeforeBlock = _deadlineBlock;
        assertion.createdAtBlock = uint64(block.number);
        assertion.assertionHash = _assertionHash;
        assertion.isFirstChild = _isFirstChild;
        assertion.configHash = _configHash;
        return assertion;
    }

    /**
     * @notice Update child properties
     * @param number The child number to set
     */
    function childCreated(AssertionNode storage self, uint64 number, uint64 confirmPeriodBlocks) internal {
        if (self.firstChildBlock == 0) {
            self.firstChildBlock = uint64(block.number);
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

    function requireMoreThanOneChild(AssertionNode memory self) internal pure {
        require(self.secondChildBlock > 0, "TOO_FEW_CHILD");
    }

    function requireExists(AssertionNode memory self) internal pure {
        require(self.createdAtBlock > 0, "ASSERTION_NOT_EXIST");
    }
}
