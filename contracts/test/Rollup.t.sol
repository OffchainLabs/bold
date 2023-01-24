// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";

import "../src/rollup/RollupProxy.sol";

import "../src/p0/RollupCore.sol";
import "../src/p0/RollupUserLogic.sol";
import "../src/p0/RollupAdminLogic.sol";

contract RollupTest is Test {
    
    address owner = address(1);
    RollupProxy rollup;
    RollupUserLogic userRollup;
    RollupAdminLogic adminRollup;
    function setUp() public {
        Config memory config = Config({
            stakeToken: address(0),
            owner: owner
        });
        RollupUserLogic userLogic = new RollupUserLogic();
        RollupAdminLogic adminLogic = new RollupAdminLogic();
        ContractDependencies memory connectedContracts = ContractDependencies({
            rollupAdminLogic: address(adminLogic),
            rollupUserLogic: address(userLogic)
        });
        rollup = new RollupProxy(config, connectedContracts);
        userRollup = RollupUserLogic(address(rollup));
        adminRollup = RollupAdminLogic(address(rollup));

        vm.startPrank(owner);
        adminRollup.setValidatorWhitelistDisabled(true);
        vm.stopPrank();
    }

    function testFoo() public {

        RollupLib.StateCommitment memory sc1 = RollupLib.StateCommitment({
            height: 1,
            stateRoot: bytes32(0)
        });
        userRollup.stakeOnNewAssertion(bytes32(0), sc1);
        RollupLib.StateCommitment memory sc2 = RollupLib.StateCommitment({
            height: 3,
            stateRoot: bytes32(0)
        });
        userRollup.stakeOnNewAssertion(RollupLib.stateCommitmentHash(sc1), sc2);
        RollupLib.StateCommitment memory sc3 = RollupLib.StateCommitment({
            height: 2,
            stateRoot: bytes32(0)
        });
        userRollup.stakeOnNewAssertion(RollupLib.stateCommitmentHash(sc1), sc3);
    }
}
