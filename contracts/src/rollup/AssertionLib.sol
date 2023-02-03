// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "./RollupLib.sol";

struct NewAssertionInputs {
    ExecutionState beforeState;
    ExecutionState afterState;
    uint64 numBlocks; // TODO: do we need this or can we just calc from height
    uint64 prevNum;
    bytes32 prevStateCommitment;
    uint256 prevNodeInboxMaxCount;
    bytes32 expectedAssertionHash;
}

enum AssertionStatus {
    Pending,
    Confirmed,
    Rejected
}

struct Assertion {
    bytes32 stateHash;
    bytes32 nodeHash;
    bytes32 challengeHash;
    bytes32 confirmHash;
    address staker;
    AssertionStatus status;
    bool notFirstChild;
    uint64 prevNum;
    uint64 firstChildCreationBlock;
    uint64 secondChildCreationBlock;
    uint64 createdAtBlock;
    bytes32 challengeId;
    uint64 challengeWinner;
}

/**
 * @notice Utility functions for Node
 */
library AssertionLib {
    using GlobalStateLib for GlobalState;

    /**
     * @notice Initialize a Node
     * @param _inputs Initial state
     * @param _inboxMaxCount Current inbox message count
     */
    function createAssertion(
        NewAssertionInputs memory _inputs,
        uint256 _inboxMaxCount // the current inbox size
    ) internal view returns (Assertion memory) {
        Assertion memory assertion;

        // assertion.nodeHash = _nodeHash;
        assertion.stateHash = RollupLib.stateHashMem(_inputs.afterState, _inboxMaxCount);
        // assertion.challengeHash = _challengeHash;
        assertion.confirmHash = RollupLib.confirmHash(
            _inputs.afterState.globalState.getBlockHash(),
            _inputs.afterState.globalState.getSendRoot()
        );

        // assertion.height = _inputs.height;
        // assertion.prevNum = _prev;
        assertion.createdAtBlock = uint64(block.number);

        return assertion;
    }

    function stateHashMem(NewAssertionInputs memory states, uint256 inboxMaxCount)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encode(states, inboxMaxCount));
    }

    // TODO: figure out what we need here
    function assertionHash(
        bool notFirstChild,
        bytes32 lastHash,
        bytes32 assertionExecHash,
        bytes32 inboxAcc,
        bytes32 wasmModuleRoot
    ) internal pure returns (bytes32) {
        uint8 notFirstChildInt = notFirstChild ? 1 : 0;
        return
            keccak256(
                abi.encode(notFirstChildInt, lastHash, assertionExecHash, inboxAcc, wasmModuleRoot)
            );
    }

    /**
     * @notice Update child properties
     * @param number The child number to set
     */
    function childCreated(Assertion storage self, uint64 number) internal {
        if (self.firstChildCreationBlock == 0) {
            self.firstChildCreationBlock = uint64(block.number);
        } else if (self.secondChildCreationBlock == 0) {
            self.secondChildCreationBlock = uint64(block.number);
        }
        // self.latestChildNumber = number; // TODO: do we need this?
    }
}
