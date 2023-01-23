// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

library RollupLib {
    struct StateCommitment {
        uint64 height;
        bytes32 stateRoot;
    }
    enum Status {
        Pending,
        Confirmed,
        Rejected
    }
    struct Assertion {
        uint256 seqNum;
        StateCommitment stateCommitment;
        address staker;
        bytes32 prev;
        Status status;
        bool isFirstChild;
        uint256 firstChildCreationBlock;
        uint256 secondChildCreationBlock;
        bytes32 challenge;
        uint256 createdAtBlock;
    }

    function assertionIsNone(Assertion memory assertion) internal pure returns (bool) {
        return assertion.createdAtBlock == 0;
    }
    function stateCommitmentHash(
        StateCommitment memory stateCommitment
    ) internal pure returns (bytes32) {
        return keccak256(abi.encode(stateCommitment));
    }
}