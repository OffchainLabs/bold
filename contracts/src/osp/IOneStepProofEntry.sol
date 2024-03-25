// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "./IOneStepProver.sol";
import "../state/Machine.sol";

library OneStepProofEntryLib {
    uint256 internal constant MAX_STEPS = 1 << 43;
}

struct ExecutionState {
    GlobalState globalState;
    MachineStatus machineStatus;
    bytes32 endHistoryRoot;
}

interface IOneStepProofEntry {
    function proveOneStep(
        ExecutionContext calldata execCtx,
        uint256 machineStep,
        bytes32 beforeHash,
        bytes calldata proof
    ) external view returns (bytes32 afterHash);

    function getMachineHash(ExecutionState calldata execState) external pure returns (bytes32);
}
