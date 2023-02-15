// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "../bridge/ISequencerInbox.sol";

struct Config {
    uint64 confirmPeriodBlocks;
    uint64 extraChallengeTimeBlocks;
    address stakeToken;
    uint256 baseStake;
    bytes32 wasmModuleRoot;
    address owner;
    address loserStakeEscrow;
    uint256 chainId;
    uint64 genesisBlockNum;
    ISequencerInbox.MaxTimeVariation sequencerInboxMaxTimeVariation;
}

struct ContractDependencies {
    address bridge;
    address sequencerInbox;
    address inbox;
    address outbox;
    address rollupEventInbox;
    address challengeManager;
    address rollupAdminLogic;
    address rollupUserLogic;
    // misc contracts that are useful when interacting with the rollup
    address validatorUtils;
    address validatorWalletCreator;
}
