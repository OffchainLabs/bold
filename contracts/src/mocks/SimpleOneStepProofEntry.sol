// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "../challengeV2/EdgeChallengeManager.sol";
import "../state/Deserialize.sol";

contract SimpleOneStepProofEntry is IOneStepProofEntry {
    using GlobalStateLib for GlobalState;

    function proveOneStep(
        ExecutionContext calldata execCtx,
        uint256 step,
        bytes32 beforeHash,
        bytes calldata proof
    ) external view returns (bytes32 afterHash) {
        GlobalState memory globalState;
        uint256 offset;
        (globalState.u64Vals[0], offset) = Deserialize.u64(proof, offset);
        (globalState.u64Vals[1], offset) = Deserialize.u64(proof, offset);
        if (step > 0 && (beforeHash[0] == 0 || globalState.getPositionInMessage() == 0)) {
            // We end the block when the first byte of the hash hits 0 or we advance a batch
            return beforeHash;
        }
        if (globalState.getInboxPosition() >= execCtx.maxInboxMessagesRead) {
            // We can't continue further because we've hit the max inbox messages read
            return beforeHash;
        }
        require(globalState.hash() == beforeHash, "BAD_PROOF");
        globalState.u64Vals[0]++;
        if (globalState.u64Vals[0] % 200 == 0) {
            globalState.u64Vals[1]++;
        }
        return globalState.hash();
    }

    function getMachineHash(GlobalState calldata globalState, MachineStatus machineStatus) external pure returns (bytes32) {
        require(machineStatus == MachineStatus.FINISHED, "BAD_MACHINE_STATUS");
        return globalState.hash();
    }
}
