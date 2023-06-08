// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "../state/GlobalState.sol";
import "../state/Machine.sol";
import "../osp/IOneStepProofEntry.sol";

enum AssertionStatus {
    // No assertion at this index
    NoAssertion,
    // Assertion is being computed
    Pending,
    // Assertion is confirmed
    Confirmed
}

struct AssertionNode {
    // Deadline at which this assertion can be confirmed
    // TODO: HN: remove this and derive from createdAtBlock?
    uint64 deadlineBlock;
    // This value starts at zero and is set to a value when the first child is created. After that it is constant until the assertion is destroyed or the owner destroys pending assertions
    uint64 firstChildBlock;
    // This value starts at zero and is set to a value when the second child is created. After that it is constant until the assertion is destroyed or the owner destroys pending assertions
    uint64 secondChildBlock;
    // The block number when this assertion was created
    uint64 createdAtBlock;
    // True if this assertion is the first child of its prev
    bool isFirstChild;
    // Status of the Assertion
    AssertionStatus status;
    // Id of the assertion previous to this one
    bytes32 prevId;
    // A hash of all configuration data when the assertion is created, used for the creation and resolution of its successor
    bytes32 configHash;
}

struct BeforeStateData {
    // The assertion hash of the prev of the beforeState(prev)
    bytes32 prevPrevAssertionHash;
    // The sequencer inbox accumulator asserted by the beforeState(prev)
    bytes32 sequencerBatchAcc;
    // below are the components of config hash
    bytes32 wasmRoot;
    uint256 requiredStake;
    address challengeManager;
    uint64 confirmPeriodBlocks;
    uint64 nextInboxPosition;
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
     * @param _prevId Initial value of prevId
     * @param _deadlineBlock Initial value of deadlineBlock
     */
    function createAssertion(
        bytes32 _prevId,
        uint64 _deadlineBlock,
        bool _isFirstChild,
        bytes32 _configHash
    ) internal view returns (AssertionNode memory) {
        AssertionNode memory assertion;
        assertion.prevId = _prevId;
        assertion.deadlineBlock = _deadlineBlock;
        assertion.createdAtBlock = uint64(block.number);
        assertion.isFirstChild = _isFirstChild;
        assertion.configHash = _configHash;
        assertion.status = AssertionStatus.Pending;
        return assertion;
    }

    /**
     * @notice Update child properties
     */
    function childCreated(AssertionNode storage self) internal {
        if (self.firstChildBlock == 0) {
            self.firstChildBlock = uint64(block.number);
        } else if (self.secondChildBlock == 0) {
            self.secondChildBlock = uint64(block.number);
        }
    }

    /**
     * @notice Check whether the current block number has met or passed the assertion's deadline
     */
    function requirePastDeadline(AssertionNode memory self) internal view {
        require(block.number >= self.deadlineBlock, "BEFORE_DEADLINE");
    }

    function requireMoreThanOneChild(AssertionNode memory self) internal pure {
        require(self.secondChildBlock > 0, "TOO_FEW_CHILD");
    }

    function requireExists(AssertionNode memory self) internal pure {
        require(self.status != AssertionStatus.NoAssertion, "ASSERTION_NOT_EXIST");
    }
}
