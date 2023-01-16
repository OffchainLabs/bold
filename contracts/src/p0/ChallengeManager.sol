// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "../libraries/DelegateCallAware.sol";
import "../osp/IOneStepProofEntry.sol";
// import "../state/GlobalState.sol";
import "../challenge/IChallengeResultReceiver.sol";
// import "./ChallengeLib.sol";
import "../challenge/IChallengeManager.sol";

import {NO_CHAL_INDEX} from "../libraries/Constants.sol";

contract ChallengeManager is DelegateCallAware, IChallengeManager {
    enum ChallengeModeRequirement {
        ANY,
        BLOCK,
        EXECUTION
    }

    string private constant NO_CHAL = "NO_CHAL";
    uint256 private constant MAX_CHALLENGE_DEGREE = 40;

    uint64 public totalChallengesCreated;
    mapping(uint256 => ChallengeLib.Challenge) public challenges;

    IChallengeResultReceiver public resultReceiver;

    ISequencerInbox public sequencerInbox;
    IBridge public bridge;
    IOneStepProofEntry public osp;

    function challengeInfo(uint64 challengeIndex)
        external
        view
        override
        returns (ChallengeLib.Challenge memory)
    {
        revert("UNIMPLEMENTED");
    }

    modifier takeTurn(
        uint64 challengeIndex,
        ChallengeLib.SegmentSelection calldata selection,
        ChallengeModeRequirement expectedMode
    ) {
        revert("UNIMPLEMENTED");
        _;
    }

    function initialize(
        IChallengeResultReceiver resultReceiver_,
        ISequencerInbox sequencerInbox_,
        IBridge bridge_,
        IOneStepProofEntry osp_
    ) external override onlyDelegated {
        revert("UNIMPLEMENTED");
    }

    function createChallenge(
        bytes32 wasmModuleRoot_,
        MachineStatus[2] calldata startAndEndMachineStatuses_,
        GlobalState[2] calldata startAndEndGlobalStates_,
        uint64 numBlocks,
        address asserter_,
        address challenger_,
        uint256 asserterTimeLeft_,
        uint256 challengerTimeLeft_
    ) external override returns (uint64) {
        revert("UNIMPLEMENTED");
    }

    /**
     * @notice Initiate the next round in the bisection by objecting to execution correctness with a bisection
     * of an execution segment with the same length but a different endpoint. This is either the initial move
     * or follows another execution objection
     */
    function bisectExecution(
        uint64 challengeIndex,
        ChallengeLib.SegmentSelection calldata selection,
        bytes32[] calldata newSegments
    ) external takeTurn(challengeIndex, selection, ChallengeModeRequirement.ANY) {
        revert("UNIMPLEMENTED");
    }

    function challengeExecution(
        uint64 challengeIndex,
        ChallengeLib.SegmentSelection calldata selection,
        MachineStatus[2] calldata machineStatuses,
        bytes32[2] calldata globalStateHashes,
        uint256 numSteps
    ) external takeTurn(challengeIndex, selection, ChallengeModeRequirement.BLOCK) {
        revert("UNIMPLEMENTED");
    }

    function oneStepProveExecution(
        uint64 challengeIndex,
        ChallengeLib.SegmentSelection calldata selection,
        bytes calldata proof
    ) external takeTurn(challengeIndex, selection, ChallengeModeRequirement.EXECUTION) {
        revert("UNIMPLEMENTED");
    }

    function timeout(uint64 challengeIndex) external override {
        revert("UNIMPLEMENTED");
    }

    function clearChallenge(uint64 challengeIndex) external override {
        revert("UNIMPLEMENTED");
    }

    function currentResponder(uint64 challengeIndex) public view override returns (address) {
        revert("UNIMPLEMENTED");
    }

    function isTimedOut(uint64 challengeIndex) public view override returns (bool) {
        revert("UNIMPLEMENTED");
    }

    function requireValidBisection(
        ChallengeLib.SegmentSelection calldata selection,
        bytes32 startHash,
        bytes32 endHash
    ) private pure {
        revert("UNIMPLEMENTED");
    }

    function completeBisection(
        uint64 challengeIndex,
        uint256 challengeStart,
        uint256 challengeLength,
        bytes32[] memory newSegments
    ) private {
        revert("UNIMPLEMENTED");
    }

    /// @dev This function causes the mode of the challenge to be set to NONE by deleting the challenge
    function _nextWin(uint64 challengeIndex, ChallengeTerminationType reason) private {
        revert("UNIMPLEMENTED");
    }

    /**
     * @dev this currently sets a challenge hash of 0 - no move is possible for the next participant to progress the
     * state. It is assumed that wherever this function is consumed, the turn is then adjusted for the opposite party
     * to timeout. This is done as a safety measure so challenges can only be resolved by timeouts during mainnet beta.
     */
    function _currentWin(
        uint64 challengeIndex,
        ChallengeTerminationType /* reason */
    ) private {
        revert("UNIMPLEMENTED");
    }
}
