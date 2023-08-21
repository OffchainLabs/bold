// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//
pragma solidity ^0.8.17;

/// @notice The status of the edge
/// - Pending: Yet to be confirmed. Not all edges can be confirmed.
/// - Confirmed: Once confirmed it cannot transition back to pending
enum EdgeStatus {
    Pending,
    Confirmed
}
