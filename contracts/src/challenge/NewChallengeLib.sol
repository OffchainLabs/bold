// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "../state/Machine.sol";
import "../state/GlobalState.sol";

library NewChallengeLib {
    using MachineLib for Machine;
    using NewChallengeLib for Challenge;

    enum VertexStatus {
        Pending,
        Confirmed,
        Rejected
    }

    struct HistoryCommitment {
        uint256 height;
        bytes32 merkleRoot;
    }

    struct Challenge {
        bytes32 wasmModuleRoot;
        bytes32 challengeStateHash;
        uint64 maxInboxMessages;
        uint256 confirmPeriodBlocks;
    }

    struct Vertex {
        address validator;
        bool isLeaf;
        VertexStatus status;
        bytes32 history;
        uint64 prev;
        uint64 presumptivSuccessor;
        uint64 psTimer;
        uint64 lastPsUpdate;
        uint64 subChallenge;
        uint64 winnerIfConfirmed;
        uint256 height;
    }

    function historyHash(HistoryCommitment memory commitment) internal pure returns (bytes32) {
        return keccak256(abi.encode(commitment.height, commitment.merkleRoot));
    }

    function bisectHeight(uint256 height) internal pure returns (uint256) {
        // TODO: implement
        return height-2;
    }

    function verifyPrefixProof(bytes32 prefix, bytes32 root, bytes32[] memory proof) internal pure returns (bool) {
        // TODO: implement
        return true;
    }
}
