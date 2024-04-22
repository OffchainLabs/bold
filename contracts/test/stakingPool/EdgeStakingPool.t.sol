// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../ERC20Mock.sol";
import "../../src/assertionStakingPool/EdgeStakingPoolCreator.sol";
import "../../src/challengeV2/EdgeChallengeManager.sol";

contract MockChallengeManager {
    uint256 i;
    IERC20 public immutable stakeToken;
    bytes32 public edgeIdToReturn = keccak256("real");

    event EdgeCreated(CreateEdgeArgs args);

    constructor(IERC20 _token) {
        i = 0;
        stakeToken = _token;
    }

    function createLayerZeroEdge(CreateEdgeArgs calldata args) external returns (bytes32) {
        stakeToken.transferFrom(msg.sender, address(this), stakeAmounts(args.level));

        emit EdgeCreated(args);

        return edgeIdToReturn;
    }

    function setEdgeIdToReturn(bytes32 edgeId) public {
        edgeIdToReturn = edgeId;
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

    function testProperInitialization(bytes32 edgeId, uint8 edgeLevel) public {
        EdgeStakingPool stakingPool = stakingPoolCreator.createPool(address(challengeManager), edgeId, edgeLevel);

        assertEq(address(stakingPoolCreator.getPool(address(challengeManager), edgeId, edgeLevel)), address(stakingPool));

        assertEq(address(stakingPool.challengeManager()), address(challengeManager));
        assertEq(stakingPool.edgeId(), edgeId);
        assertEq(address(stakingPool.stakeToken()), address(token));
        assertEq(stakingPool.edgeLevel(), edgeLevel);
        assertEq(stakingPool.requiredStake(), challengeManager.stakeAmounts(edgeLevel));
    }

    function testCreateEdge(CreateEdgeArgs memory args) public {
        EdgeStakingPool stakingPool = stakingPoolCreator.createPool(address(challengeManager), keccak256("real"), args.level);

        // simulate deposits
        // we don't need to deposit using the staking pool's deposit function because we're not testing that here
        token.transfer(address(stakingPool), stakingPool.requiredStake() - 1);
        vm.expectRevert("ERC20: transfer amount exceeds balance");
        stakingPool.createEdge(args);
        token.transfer(address(stakingPool), 1);

        // simulate an incorrect edge id
        challengeManager.setEdgeIdToReturn(keccak256("fake"));
        vm.expectRevert(abi.encodeWithSelector(EdgeStakingPool.IncorrectEdgeId.selector, keccak256("fake"), keccak256("real")));
        stakingPool.createEdge(args);
        challengeManager.setEdgeIdToReturn(keccak256("real"));

        vm.expectEmit(false, false, false, true);
        emit EdgeCreated(args);
        stakingPool.createEdge(args);

        assertEq(token.balanceOf(address(stakingPool)), 0);
        assertEq(token.balanceOf(address(challengeManager)), stakingPool.requiredStake());
    }
}
