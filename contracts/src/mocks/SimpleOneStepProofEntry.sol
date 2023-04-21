// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "../../src/challengeV2/EdgeChallengeManager.sol";

contract SimpleOneStepProofEntry is IOneStepProofEntry {
    function proveOneStep(
        ExecutionContext calldata,
        uint256 step,
        bytes32 beforeHash,
        bytes calldata proof
    ) external view returns (bytes32 afterHash) {
        if (beforeHash[0] == 0 && step > 0) {
            // We end the block when the first byte of the hash hits 0
            return beforeHash;
        }
        uint256 state = uint256(bytes32(proof));
        require(keccak256(abi.encodePacked(state)) == beforeHash, "BAD_PROOF");
        state++;
        return keccak256(abi.encodePacked(state));
    }

    function getMachineHash(GlobalState calldata globalState, MachineStatus machineStatus) external pure returns (bytes32) {
        require(machineStatus == MachineStatus.FINISHED, "BAD_MACHINE_STATUS");
        return GlobalStateLib.hash(globalState);
    }
}
