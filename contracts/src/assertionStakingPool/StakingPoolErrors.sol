// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//
pragma solidity ^0.8.0;
import "../rollup/IRollupLogic.sol";

enum PoolState {
    PENDING,
    ASSERTED,
    INACTIVE
}

error PoolNotInPendingState(PoolState poolState);

error PoolNotInAssertedState(PoolState poolState);

error PoolNotInPendingOrInactiveState(PoolState poolState);

error PoolNotInInactiveState(PoolState poolState);

error PoolStakeAlreadyReached(uint256 baseStake);

error NotEnoughStake(uint256 balance, uint256 baseStake);

error NoBalanceToWithdraw(address sender);

error AssertionNotInactive(bytes32 assertionHash);

error PoolDoesntExist(address rollup, AssertionInputs assertionInputs, bytes32 assertionHash);
