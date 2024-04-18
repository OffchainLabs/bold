// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../ERC20Mock.sol";
import "../../src/assertionStakingPool/EdgeStakingPoolCreator.sol";
import "../../src/challengeV2/EdgeChallengeManager.sol";

contract MockChallengeManager {
    uint256 i;
    IERC20 public immutable stakeToken;
    struct MockEdge {
        address staker;
        uint256 amount;
    }

    mapping(bytes32 => MockEdge) public edges;

    constructor(IERC20 _token) {
        i = 0;
        stakeToken = _token;
    }

    function createLayerZeroEdge(CreateEdgeArgs calldata args) external returns (bytes32) {
        bytes32 edgeId = keccak256(abi.encode(i));
        edges[edgeId] = MockEdge(msg.sender, stakeAmounts(args.level));
        i++;

        stakeToken.transferFrom(msg.sender, address(this), stakeAmounts(args.level));

        return edgeId;
    }

    function refundStake(bytes32 edgeId) external {
        stakeToken.transfer(edges[edgeId].staker, edges[edgeId].amount);
        edges[edgeId].amount = 0;
    }

    function stakeAmounts(uint256 lvl) public pure returns (uint256) {
        return 100 * (lvl + 1);
    }
}

contract EdgeStakingPoolTest is Test {
    IERC20 token;
    MockChallengeManager challengeManager;
    EdgeStakingPoolCreator stakingPoolCreator;

    function setUp() public {
        token = new ERC20Mock("TEST", "TST", address(this), 100 ether);
        challengeManager = new MockChallengeManager(token);
        stakingPoolCreator = new EdgeStakingPoolCreator();
    }

    function testProperInitialization(CreateEdgeArgs memory args) public {
        args.level &= 0xFF; // Ensure level is within 0-255
        EdgeStakingPool stakingPool = stakingPoolCreator.createPool(address(challengeManager), args);

        assertEq(address(stakingPool.challengeManager()), address(challengeManager));
        assertEq(stakingPool.createEdgeArgsHash(), keccak256(abi.encode(args)));
        assertEq(address(stakingPool.stakeToken()), address(token));
        assertEq(stakingPool.requiredStake(), challengeManager.stakeAmounts(args.level));
    }
}
