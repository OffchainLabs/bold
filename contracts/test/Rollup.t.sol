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
    }

    function testFoo() public {
        assertTrue(true);
    }
}
