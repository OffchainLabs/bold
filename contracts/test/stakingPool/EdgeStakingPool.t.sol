// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../ERC20Mock.sol";
import "../../src/assertionStakingPool/EdgeStakingPoolCreator.sol";
import "../../src/challengeV2/EdgeChallengeManager.sol";

contract MockChallengeManager {
    uint256 i;
    IERC20 public immutable stakeToken;

    event EdgeCreated(CreateEdgeArgs args);

    constructor(IERC20 _token) {
        i = 0;
        stakeToken = _token;
    }

    function createLayerZeroEdge(CreateEdgeArgs calldata args) external returns (bytes32) {
        bytes32 edgeId = keccak256(abi.encode(i++));

        stakeToken.transferFrom(msg.sender, address(this), stakeAmounts(args.level));

        emit EdgeCreated(args);

        return edgeId;
    }

    function stakeAmounts(uint256 lvl) public pure returns (uint256) {
        return 100 * (lvl + 1);
    }
}

contract EdgeStakingPoolTest is Test {
    IERC20 token;
    MockChallengeManager challengeManager;
    EdgeStakingPoolCreator stakingPoolCreator;

    event EdgeCreated(CreateEdgeArgs args);

    function setUp() public {
        token = new ERC20Mock("TEST", "TST", address(this), 100 ether);
        challengeManager = new MockChallengeManager(token);
        stakingPoolCreator = new EdgeStakingPoolCreator();
    }

    function testProperInitialization(CreateEdgeArgs memory args) public {
        EdgeStakingPool stakingPool = stakingPoolCreator.createPool(address(challengeManager), args);

        assertEq(address(stakingPool.challengeManager()), address(challengeManager));
        assertEq(stakingPool.createEdgeArgsHash(), keccak256(abi.encode(args)));
        assertEq(address(stakingPool.stakeToken()), address(token));
        assertEq(stakingPool.requiredStake(), challengeManager.stakeAmounts(args.level));
    }

    function testCheckCreateEdgeArgs(CreateEdgeArgs memory args) public {
        EdgeStakingPool stakingPool = stakingPoolCreator.createPool(address(challengeManager), args);

        assertTrue(stakingPool.isCorrectCreateEdgeArgs(args));
        args.level = ~args.level;
        assertFalse(stakingPool.isCorrectCreateEdgeArgs(args));
    }

    function testCreateEdge(CreateEdgeArgs memory args) public {
        EdgeStakingPool stakingPool = stakingPoolCreator.createPool(address(challengeManager), args);

        // simulate deposits
        // we don't need to deposit using the staking pool's deposit function because we're not testing that here
        token.transfer(address(stakingPool), stakingPool.requiredStake() - 1);
        vm.expectRevert("ERC20: transfer amount exceeds balance");
        stakingPool.createEdge(args);
        token.transfer(address(stakingPool), 1);

        args.level = ~args.level;
        vm.expectRevert(EdgeStakingPool.IncorrectCreateEdgeArgs.selector);
        stakingPool.createEdge(args);
        args.level = ~args.level;

        vm.expectEmit(false, false, false, true);
        emit EdgeCreated(args);
        stakingPool.createEdge(args);

        assertEq(token.balanceOf(address(stakingPool)), 0);
        assertEq(token.balanceOf(address(challengeManager)), stakingPool.requiredStake());
    }
}
