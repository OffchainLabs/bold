// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.17;

import "forge-std/Test.sol";
import "../src/rollup/ValidatorWallet.sol";
import "../src/rollup/ValidatorWalletCreator.sol";
import "./RollupMock.sol";

contract ValidatorWalletTest is Test {
    // Contract instances
    ValidatorWalletCreator creator;
    ValidatorWallet wallet;
    RollupMock rollupMock1;
    RollupMock rollupMock2;

    // Set initial executor and owner
    address public executor = 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4;
    address public owner = 0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2;

    function setUp() public {
        // Create instances
        creator = new ValidatorWalletCreator();
        rollupMock1 = new RollupMock();
        rollupMock2 = new RollupMock();

        // Set up executor and owner
        address[] memory initialExecutors = new address[](1);
        initialExecutors[0] = executor;

        // Create wallet and explicitly initialize with executor and owner
        vm.prank(owner);
        address walletAddress = creator.createWallet(initialExecutors);
        wallet = ValidatorWallet(payable(walletAddress));
    }

    // Test wallet creation and executor setup
    function testWalletCreation() public {
        address[] memory newExecutors = new address[](1);
        newExecutors[0] = executor;
        bool[] memory isExecutor = new bool[](1);
        isExecutor[0] = true;

        // Set executor and check if it is marked as executor
        vm.expectEmit(true, false, false, false);
        emit ExecutorUpdated(executor, true);
        vm.prank(owner);
        wallet.setExecutor(newExecutors, isExecutor);
        require(wallet.executors(executor), "Executor should be marked as executor");
        require(wallet.allowedExecutorDestinations(executor), "Executor should be allowed");
    }

    // Test setting allowed executor destinations
    function testSetAllowedExecutorDestinations() public {
        address[] memory allowedAddrs = new address[](3);
        allowedAddrs[0] = address(0x1234567812345678123456781234567812345678);
        allowedAddrs[1] = address(0x0000000000000000000000000000000000000000);
        allowedAddrs[2] = address(0x0123000000000000000000000000000000000000);
        bool[] memory isSet = new bool[](3);
        isSet[0] = true;
        isSet[1] = true;
        isSet[2] = true;

        // Set allowed executor destinations
        vm.prank(owner);
        wallet.setAllowedExecutorDestinations(allowedAddrs, isSet);
        // Check if the set destinations are allowed
        require(wallet.allowedExecutorDestinations(allowedAddrs[0]), "Address 0 should be allowed");
        require(wallet.allowedExecutorDestinations(allowedAddrs[1]), "Address 1 should be allowed");
        require(wallet.allowedExecutorDestinations(allowedAddrs[2]), "Address 2 should be allowed");

        // Check if a non-set destination is not allowed
        require(
            !wallet.allowedExecutorDestinations(address(0x1114567812345678123456781234567812341111)),
            "Address 0x1114567812345678123456781234567812341111 should not be allowed"
        );
    }

    function testAllowExecutorToExecuteTxs() public {
        address[] memory newExecutors = new address[](1);
        newExecutors[0] = executor;
        bool[] memory isExecutor = new bool[](1);
        isExecutor[0] = true;
        address[] memory destinationAddrs = new address[](1);
        destinationAddrs[0] = address(rollupMock1);
        bool[] memory isSet = new bool[](1);
        isSet[0] = true;
        bytes memory data = abi.encodeWithSignature("withdrawStakerFunds()");

        // Set executor
        vm.prank(owner);
        wallet.setExecutor(newExecutors, isExecutor);
        // Expect the transaction to revert with a specific error message when called by the executor
        vm.expectRevert(abi.encodeWithSelector(OnlyOwnerDestination.selector, owner, executor, address(rollupMock1)));
        vm.prank(executor);
        wallet.executeTransaction(data, address(rollupMock1), 0);

        vm.expectEmit(false, false, false, false);
        // Expecet the transaction to go through when called by owenr
        emit WithdrawTriggered();
        vm.prank(owner);
        wallet.executeTransaction(data, address(rollupMock1), 0);

        vm.prank(owner);
        // Set allowed executor destinations using the provided arrays
        wallet.setAllowedExecutorDestinations(destinationAddrs, isSet);

        // Expect the transaction to go through when called by the executor
        emit WithdrawTriggered();
        vm.prank(executor);
        wallet.executeTransaction(data, address(rollupMock1), 0);
    }

    function testRejectBatchIfSingleTxNotAllowed() external {
        address[] memory addrs = new address[](1);
        addrs[0] = executor;
        bool[] memory isExecutor = new bool[](1);
        isExecutor[0] = true;

        // Encode function data for 'withdrawStakerFunds' in RollupMock
        bytes memory data1 = abi.encodeWithSignature("withdrawStakerFunds()");
        bytes memory data2 = abi.encodeWithSignature("withdrawStakerFunds()");

        bytes[] memory data = new bytes[](2);
        data[0] = data1;
        data[1] = data2;

        uint256[] memory amount = new uint256[](2);
        amount[0] = 0;
        amount[1] = 0;

        // Set allowed executor destinations in wallet
        address[] memory destination_addrs = new address[](2);
        destination_addrs[0] = address(rollupMock1);
        destination_addrs[1] = address(rollupMock2);
        bool[] memory isSet = new bool[](2);
        isSet[0] = true;
        isSet[1] = false;

        // Set executor
        vm.prank(owner);
        wallet.setExecutor(addrs, isExecutor);

        // Set allowed executor destinations using the provided arrays
        vm.prank(owner);
        wallet.setAllowedExecutorDestinations(destination_addrs, isSet);

        // Expect the transaction to revert with a specific error message when called by the executor
        vm.expectRevert(abi.encodeWithSelector(OnlyOwnerDestination.selector, owner, executor, address(rollupMock2)));
        vm.prank(executor);
        wallet.executeTransactions(data, destination_addrs, amount);
    }

    /**
     *
     * Event declarations
     *
     */

    event WithdrawTriggered();
    event ExecutorUpdated(address indexed executor, bool isExecutor);
}
