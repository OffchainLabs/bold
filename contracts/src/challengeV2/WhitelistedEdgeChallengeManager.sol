// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//
pragma solidity ^0.8.17;

import "./EdgeChallengeManager.sol";

interface IWhitelistedEdgeChallengeManager is IEdgeChallengeManager {
    function owner() external view returns (address);
    function isValidator(address account) external view returns (bool);
    function accountHasMadeRival(bytes32 mutualId, address account) external view returns (bool);

    function setValidator(address[] calldata accounts, bool[] calldata vals) external;
}

contract WhitelistedEdgeChallengeManager is EdgeChallengeManager, IWhitelistedEdgeChallengeManager {
    using ChallengeEdgeLib for ChallengeEdge;

    address public owner;
    mapping(address => bool) public isValidator;
    mapping(bytes32 => mapping(address => bool)) public accountHasMadeRival;

    error AccountHasMadeRival(address account);
    error NotValidator(address account);
    error NotOwner(address account);

    function initialize(
        IAssertionChain _assertionChain,
        uint64 _challengePeriodBlocks,
        IOneStepProofEntry _oneStepProofEntry,
        uint256 layerZeroBlockEdgeHeight,
        uint256 layerZeroBigStepEdgeHeight,
        uint256 layerZeroSmallStepEdgeHeight,
        IERC20 _stakeToken,
        address _excessStakeReceiver,
        uint8 _numBigStepLevel,
        uint256[] calldata _stakeAmounts
    ) public override(EdgeChallengeManager, IEdgeChallengeManager) initializer {
        super.initialize(
            _assertionChain,
            _challengePeriodBlocks,
            _oneStepProofEntry,
            layerZeroBlockEdgeHeight,
            layerZeroBigStepEdgeHeight,
            layerZeroSmallStepEdgeHeight,
            _stakeToken,
            _excessStakeReceiver,
            _numBigStepLevel,
            _stakeAmounts
        );

        owner = msg.sender;
    }

    function setValidator(address[] calldata accounts, bool[] calldata vals) public {
        if (msg.sender != owner) {
            revert NotOwner(msg.sender);
        }

        for (uint256 i = 0; i < accounts.length; i++) {
            isValidator[accounts[i]] = vals[i];
        }
    }

    function createLayerZeroEdge(CreateEdgeArgs calldata args) public override(EdgeChallengeManager, IEdgeChallengeManager) returns (bytes32) {
        if (!isValidator[msg.sender]) {
            revert NotValidator(msg.sender);
        }

        bytes32 edgeId = super.createLayerZeroEdge(args);
        bytes32 mutualId = getEdge(edgeId).mutualIdMem();

        if (accountHasMadeRival[mutualId][msg.sender]) {
            revert AccountHasMadeRival(msg.sender);
        }

        accountHasMadeRival[mutualId][msg.sender] = true;

        return edgeId;
    }

    function usesWhitelist() external pure override(EdgeChallengeManager, IEdgeChallengeManager) returns (bool) {
        return true;
    }
}