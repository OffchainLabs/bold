// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "forge-std/StdJson.sol";

import "../src/rollup/BOLDUpgradeAction.sol";

/*
to use this test:

anvil --fork-url $ETH_URL > /dev/null &
PID=$!
sleep 1 && yarn script:bold-prepare && yarn script:bold-populate-lookup && forge test --match-contract BOLDUpgradeActionForkTest -vvv
kill $PID

*/

contract BOLDUpgradeActionForkTest is Test {
    using stdJson for string;

    string deployedContractsJson;
    string configJson;
    address boldAction;

    function setUp() external {
        uint256 forkId = vm.createFork("http://127.0.0.1:8545");
        vm.selectFork(forkId);

        deployedContractsJson = vm.readFile("./scripts/files/mainnetDeployedContracts.json");
        configJson = vm.readFile("./scripts/files/mainnetConfig.json");

        boldAction = deployedContractsJson.readAddress(".boldAction");
    }

    function testFoo() external {
        address[] memory validators = configJson.readAddressArray(".validators");

        bytes memory payload = abi.encodeWithSignature("execute(address,bytes)", boldAction, abi.encodeCall(
            BOLDUpgradeAction.perform,
            (validators)
        ));

        vm.prank(deployedContractsJson.readAddress(".l1Timelock"));
        (bool b,) = deployedContractsJson.readAddress(".l1Executor").call(payload);

        assertTrue(b, "failed to execute");
    }
}