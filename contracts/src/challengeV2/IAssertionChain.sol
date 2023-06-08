// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "../bridge/IBridge.sol";
import "../osp/IOneStepProofEntry.sol";

/// @title  Assertion chain interface
/// @notice The interface required by the EdgeChallengeManager for requesting assertion data from the AssertionChain
interface IAssertionChain {
    function bridge() external view returns (IBridge);
    function getPredecessorId(bytes32 assertionId) external view returns (bytes32);
    function proveExecutionState(bytes32 assertionId, ExecutionState memory state, bytes memory proof)
        external
        view
        returns (ExecutionState memory);
    function validateConfig(
        bytes32 assertionId,
        bytes32 _wasmModuleRoot,
        uint256 _requiredStake,
        address _challengeManager,
        uint64 _confirmPeriodBlocks,
        uint64 _nextInboxPosition
    ) external view;
    function hasSibling(bytes32 assertionId) external view returns (bool);
    function getFirstChildCreationBlock(bytes32 assertionId) external view returns (uint256);
    function getSecondChildCreationBlock(bytes32 assertionId) external view returns (uint256);
    function isFirstChild(bytes32 assertionId) external view returns (bool);
    function isPending(bytes32 assertionId) external view returns (bool);
}
