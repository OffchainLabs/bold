// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "forge-std/Test.sol";
import {IAssertionChain} from "../src/challengeV2/IAssertionChain.sol";
import { IEdgeChallengeManager } from "../src/challengeV2/EdgeChallengeManager.sol";
import "../src/bridge/IBridge.sol";
import "../src/rollup/RollupLib.sol";
import "./challengeV2/StateTools.sol";

// CHRIS: TODO: we should update this to use the real assertion, not the mock
struct MockAssertion {
    bytes32 predecessorId;
    uint256 height;
    ExecutionState state;
    bytes32 successionChallenge;
    uint256 firstChildCreationBlock;
    uint256 secondChildCreationBlock;
    bool isFirstChild;
    bool isPending;
    bytes32 configHash;
}

contract MockAssertionChain is IAssertionChain {
    mapping(bytes32 => MockAssertion) assertions;
    IBridge public bridge; // TODO: set bridge in this mock
    bytes32 public wasmModuleRoot;
    uint256 public baseStake;
    address public challengeManager;
    uint64 public confirmPeriodBlocks;

    function assertionExists(bytes32 assertionId) public view returns (bool) {
        return assertions[assertionId].height != 0;
    }

    function stakeToken() public view returns(address) {
        return address(0);
    }

    function validateAssertionId(
        bytes32 assertionId,
        ExecutionState calldata state,
        bytes32 prevAssertionId,
        bytes32 inboxAcc
    ) external view {
        require(assertionExists(assertionId), "Assertion does not exist");
        // TODO: HN: This is not how the real assertion chain calculate assertion id
        require(assertionId == calculateAssertionId(prevAssertionId, state), "INVALID_ASSERTION_HASH");
    }

    function getFirstChildCreationBlock(bytes32 assertionId) external view returns (uint256) {
        require(assertionExists(assertionId), "Assertion does not exist");
        return assertions[assertionId].firstChildCreationBlock;
    }

    function getSecondChildCreationBlock(bytes32 assertionId) external view returns (uint256) {
        require(assertionExists(assertionId), "Assertion does not exist");
        return assertions[assertionId].secondChildCreationBlock;
    }

    function validateConfig(
        bytes32 assertionId,
        ConfigData calldata configData
    ) external view {
        require(
            RollupLib.configHash({
                wasmModuleRoot: configData.wasmModuleRoot,
                requiredStake: configData.requiredStake,
                challengeManager: configData.challengeManager,
                confirmPeriodBlocks: configData.confirmPeriodBlocks,
                nextInboxPosition: configData.nextInboxPosition
            }) == assertions[assertionId].configHash,
            "BAD_CONFIG"
        );
    }

    function isFirstChild(bytes32 assertionId) external view returns (bool) {
        require(assertionExists(assertionId), "Assertion does not exist");
        return assertions[assertionId].isFirstChild;
    }

    function isPending(bytes32 assertionId) external view returns (bool) {
        require(assertionExists(assertionId), "Assertion does not exist");
        return assertions[assertionId].isPending;
    }

    function calculateAssertionId(
        bytes32 predecessorId,
        ExecutionState memory afterState
    )
        public
        view
        returns (bytes32)
    {
        return RollupLib.assertionHash({
            parentAssertionHash: predecessorId,
            afterState: afterState,
            inboxAcc: keccak256(abi.encode(afterState.globalState.u64Vals[0])) // mock accumulator based on inbox count
        });
    }

    function childCreated(bytes32 assertionId) internal {
        if (assertions[assertionId].firstChildCreationBlock == 0) {
            assertions[assertionId].firstChildCreationBlock = block.number;
        } else if (assertions[assertionId].secondChildCreationBlock == 0) {
            assertions[assertionId].secondChildCreationBlock = block.number;
        }
    }

    function addAssertionUnsafe(
        bytes32 predecessorId,
        uint256 height,
        uint64 nextInboxPosition,
        ExecutionState memory afterState,
        bytes32 successionChallenge
    ) public returns (bytes32) {
        bytes32 assertionId = calculateAssertionId(predecessorId, afterState);
        assertions[assertionId] = MockAssertion({
            predecessorId: predecessorId,
            height: height,
            state: afterState,
            successionChallenge: successionChallenge,
            firstChildCreationBlock: 0,
            secondChildCreationBlock: 0,
            isFirstChild: assertions[predecessorId].firstChildCreationBlock == 0,
            isPending: true,
            configHash: RollupLib.configHash({
                wasmModuleRoot: wasmModuleRoot,
                requiredStake: baseStake,
                challengeManager: challengeManager,
                confirmPeriodBlocks: confirmPeriodBlocks,
                nextInboxPosition: nextInboxPosition
            })
        });
        childCreated(predecessorId);
        return assertionId;
    }

    function addAssertion(
        bytes32 predecessorId,
        uint256 height,
        uint64 nextInboxPosition,
        ExecutionState memory beforeState,
        ExecutionState memory afterState,
        bytes32 successionChallenge
    ) public returns (bytes32) {
        bytes32 beforeStateHash = StateToolsLib.hash(beforeState);
        bytes32 assertionId = calculateAssertionId(predecessorId, afterState);
        require(!assertionExists(assertionId), "Assertion already exists");
        require(assertionExists(predecessorId), "Predecessor does not exists");
        require(height > assertions[predecessorId].height, "Height too low");
        require(beforeStateHash == StateToolsLib.hash(assertions[predecessorId].state), "Before state hash does not match predecessor");

        return addAssertionUnsafe(predecessorId, height, nextInboxPosition, afterState, successionChallenge);
    }
}
