// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

// TODO: assertionId is currently casted to a uint64 assertionNum
//       we might change it from index based to hash based in the future
interface IAssertionChain {
    function getPredecessorId(bytes32 assertionId) external view returns (bytes32);

    function getHeight(bytes32 assertionId) external view returns (uint256);

    function getInboxMsgCountSeen(bytes32 assertionId) external view returns (uint256);

    function getStateHash(bytes32 assertionId) external view returns (bytes32);

    function getSuccessionChallenge(bytes32 assertionId) external view returns (bytes32);

    function isFirstChild(bytes32 assertionId) external view returns (bool);
}
