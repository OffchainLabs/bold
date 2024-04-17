// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";

import "../../src/assertionStakingPool/AbsBoldStakingPool.sol";
import "../../src/mocks/TestWETH9.sol";


contract FundsHolder {
    function withdraw(IERC20 stakeToken, uint256 amount) external {
        stakeToken.transfer(msg.sender, amount);
    }
}

contract FakeStakingPool is AbsBoldStakingPool {
    FundsHolder public immutable fundsHolder;
    uint256 immutable reqStake;

    constructor(IERC20 _stakeToken, FundsHolder _fundsHolder, uint256 _reqStake) AbsBoldStakingPool(_stakeToken) {
        fundsHolder = _fundsHolder;
        reqStake = _reqStake;
    }

    function getRequiredStake() public view override returns (uint256) {
        return reqStake;
    }

    function createMove() external override {
        stakeToken.transfer(address(fundsHolder), reqStake);
    }

    function withdrawStakeBackIntoPool() external override {
        fundsHolder.withdraw(stakeToken, reqStake);
    }
}

contract AbsBoldStakingPoolTest is Test {
    uint256 constant BASE_STAKE = 10 ether;

    address staker1 = address(4000001);
    address staker2 = address(4000002);
    address excessStaker = address(4000003);
    address fullStaker = address(4000004);

    uint256 staker1Bal = 6 ether;
    uint256 staker2Bal = 4 ether;
    uint256 fullStakerBal = 10 ether;
    uint256 excessStakerBal = 1 ether;

    IERC20 token;
    FakeStakingPool pool;

    event StakeDeposited(address indexed sender, uint256 amount);

    function setUp() public {
        token = new TestWETH9("Test", "TEST");
        pool = new FakeStakingPool(token, new FundsHolder(), BASE_STAKE);

        
        IWETH9(address(token)).deposit{value: 21 ether}();

        token.transfer(staker1, staker1Bal);
        token.transfer(staker2, staker2Bal);
        token.transfer(fullStaker, fullStakerBal);

        token.transfer(excessStaker, excessStakerBal);

        vm.prank(staker1);
        token.approve(address(pool), type(uint256).max);

        vm.prank(staker2);
        token.approve(address(pool), type(uint256).max);

        vm.prank(fullStaker);
        token.approve(address(pool), type(uint256).max);

        vm.prank(excessStaker);
        token.approve(address(pool), type(uint256).max);
    }

    function testDepositCap() external {
        vm.prank(fullStaker);
        vm.expectEmit(true, false, false, true);
        emit StakeDeposited(fullStaker, 9 ether);
        pool.depositIntoPool(9 ether);
        vm.prank(staker2);
        vm.expectEmit(true, false, false, true);
        emit StakeDeposited(staker2, 1 ether);
        pool.depositIntoPool(2 ether);

        assertEq(token.balanceOf(address(pool)), 10 ether, "tokens depositted into pool");
        assertEq(token.balanceOf(address(fullStaker)), 1 ether, "full staker balance");
        assertEq(token.balanceOf(address(staker2)), staker2Bal - 1 ether, "staker2 balance");

        assertEq(pool.depositedTokenBalances(fullStaker), 9 ether, "full staker balance");
        assertEq(pool.depositedTokenBalances(staker2), 1 ether, "staker2 balance");

        // since we are at cap, we can't deposit more
        vm.expectRevert(AbsBoldStakingPool.RequiredStakeAmountMet.selector);
        vm.prank(staker1);
        pool.depositIntoPool(1 ether);
    }

    function testCanDepositAndWithdrawWhilePending() external {
        vm.prank(staker1);
        pool.depositIntoPool(staker1Bal);
        vm.prank(staker2);
        pool.depositIntoPool(staker2Bal);

        assertEq(token.balanceOf(address(pool)), staker1Bal + staker2Bal, "tokens depositted into pool");
        assertEq(token.balanceOf(address(staker1)), uint256(0), "tokens depositted into pool");
        assertEq(token.balanceOf(address(staker2)), uint256(0), "tokens depositted into pool");

        vm.prank(staker1);
        pool.withdrawFromPool();

        vm.prank(staker2);
        pool.withdrawFromPool();

        assertEq(token.balanceOf(address(pool)), uint256(0), "tokens withdrawn from pool");
        assertEq(token.balanceOf(address(staker1)), staker1Bal, "tokens withdrawn from pool");
        assertEq(token.balanceOf(address(staker2)), staker2Bal, "tokens withdrawn from pool");
    }

    function testCantAssertWithInsufficientStake() external {
        vm.prank(staker1);
        pool.depositIntoPool(staker1Bal);

        vm.expectRevert("ERC20: transfer amount exceeds balance");
        pool.createMove();
    }

    function testPartialWithdraw() external {
        vm.startPrank(fullStaker);
        pool.depositIntoPool(fullStakerBal);

        pool.withdrawFromPool(1000);
        assertEq(token.balanceOf(fullStaker), 1000, "partial stake returned");

        vm.stopPrank();
    }

    function testReturnStake() external {
        vm.prank(staker1);
        pool.depositIntoPool(staker1Bal);
        vm.prank(staker2);
        pool.depositIntoPool(staker2Bal);

        pool.createMove();
        assertEq(token.balanceOf(address(pool)), 0, "tokens sent from pool");

        vm.expectRevert("ERC20: transfer amount exceeds balance");
        vm.prank(staker1);
        pool.withdrawFromPool();

        pool.withdrawStakeBackIntoPool();
        assertEq(token.balanceOf(address(pool)), BASE_STAKE, "tokens returned to pool");

        vm.prank(staker1);
        pool.withdrawFromPool();

        vm.prank(staker2);
        pool.withdrawFromPool();

        assertEq(token.balanceOf(address(pool)), 0, "tokens returned to users");
        assertEq(token.balanceOf(staker1), staker1Bal, "tokens returned to users");
        assertEq(token.balanceOf(staker2), staker2Bal, "tokens returned to users");
    }

    function testCantWithdrawTwice() external {
        vm.prank(staker1);
        pool.depositIntoPool(staker1Bal);
        vm.prank(staker2);
        pool.depositIntoPool(staker2Bal);

        pool.createMove();
        assertEq(token.balanceOf(address(pool)), 0, "tokens sent from pool");

        vm.expectRevert("ERC20: transfer amount exceeds balance");
        vm.prank(staker1);
        pool.withdrawFromPool();

        pool.withdrawStakeBackIntoPool();
        assertEq(token.balanceOf(address(pool)), BASE_STAKE, "tokens returned to pool");

        vm.startPrank(staker1);
        pool.withdrawFromPool();
        vm.expectRevert(abi.encodeWithSelector(AbsBoldStakingPool.ZeroAmount.selector));
        pool.withdrawFromPool();
        vm.stopPrank();
    }
}
