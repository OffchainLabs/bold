// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//
pragma solidity ^0.8.17;

import "./EdgeChallengeManager.sol";

contract EdgeChallengeManagerFactory {
    function createChallengeManager(
        IAssertionChain _assertionChain,
        uint64 _challengePeriodBlocks,
        IOneStepProofEntry _oneStepProofEntry,
        uint256 layerZeroBlockEdgeHeight,
        uint256 layerZeroBigStepEdgeHeight,
        uint256 layerZeroSmallStepEdgeHeight,
        IERC20 _stakeToken,
        uint256 _stakeAmount,
        address _excessStakeReceiver,
        uint8 _numBigStepLevel
    ) external returns (EdgeChallengeManager) {
        return new EdgeChallengeManager(
            _assertionChain,
            _challengePeriodBlocks,
            _oneStepProofEntry,
            layerZeroBlockEdgeHeight,
            layerZeroBigStepEdgeHeight,
            layerZeroSmallStepEdgeHeight,
            _stakeToken,
            _stakeAmount,
            _excessStakeReceiver,
            _numBigStepLevel
        );
    }
}