// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "./BridgeCreator.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

import "./RollupProxy.sol";

contract RollupCreator is Ownable {
    event RollupCreated(
        address indexed rollupAddress, address inboxAddress, address adminProxy, address sequencerInbox, address bridge
    );
    event TemplatesUpdated();

    BridgeCreator public bridgeCreator;
    IOneStepProofEntry public osp;
    IEdgeChallengeManager public challengeManagerTemplate;
    IRollupAdmin public rollupAdminLogic;
    IRollupUser public rollupUserLogic;

    address public validatorUtils;
    address public validatorWalletCreator;

    constructor() Ownable() {}

    function setTemplates(
        BridgeCreator _bridgeCreator,
        IOneStepProofEntry _osp,
        IEdgeChallengeManager _challengeManagerLogic,
        IRollupAdmin _rollupAdminLogic,
        IRollupUser _rollupUserLogic,
        address _validatorUtils,
        address _validatorWalletCreator
    ) external onlyOwner {
        bridgeCreator = _bridgeCreator;
        osp = _osp;
        challengeManagerTemplate = _challengeManagerLogic;
        rollupAdminLogic = _rollupAdminLogic;
        rollupUserLogic = _rollupUserLogic;
        validatorUtils = _validatorUtils;
        validatorWalletCreator = _validatorWalletCreator;
        emit TemplatesUpdated();
    }

    // internal function to workaround stack limit
    function createChallengeManager(address rollupAddr, address proxyAdminAddr, Config memory config)
        internal
        returns (IEdgeChallengeManager)
    {
        IEdgeChallengeManager challengeManager = IEdgeChallengeManager(
            address(
                new TransparentUpgradeableProxy(
                    address(challengeManagerTemplate),
                    proxyAdminAddr,
                    ""
                )
            )
        );

        challengeManager.initialize({
            _assertionChain: IAssertionChain(rollupAddr),
            _challengePeriodBlocks: config.confirmPeriodBlocks,
            _oneStepProofEntry: osp,
            layerZeroBlockEdgeHeight: config.layerZeroBlockEdgeHeight,
            layerZeroBigStepEdgeHeight: config.layerZeroBigStepEdgeHeight,
            layerZeroSmallStepEdgeHeight: config.layerZeroSmallStepEdgeHeight,
            _stakeToken: IERC20(config.stakeToken),
            _stakeAmount: config.miniStakeValue,
            _excessStakeReceiver: config.owner
        });

        return challengeManager;
    }

    // After this setup:
    // Rollup should be the owner of bridge
    // RollupOwner should be the owner of Rollup's ProxyAdmin
    // RollupOwner should be the owner of Rollup
    // Bridge should have a single inbox and outbox
    function createRollup(Config calldata config) external returns (address) {
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        proxyAdmin.transferOwnership(config.owner);

        // Create the rollup proxy to figure out the address and initialize it later
        RollupProxy rollup = new RollupProxy{salt: keccak256(abi.encode(config))}();

        (
            IBridge bridge,
            ISequencerInbox sequencerInbox,
            IInbox inbox,
            IRollupEventInbox rollupEventInbox,
            IOutbox outbox
        ) = bridgeCreator.createBridge(address(proxyAdmin), address(rollup), config.sequencerInboxMaxTimeVariation);

        IEdgeChallengeManager challengeManager = createChallengeManager(address(rollup), address(proxyAdmin), config);

        rollup.initializeProxy(
            config,
            ContractDependencies({
                bridge: bridge,
                sequencerInbox: sequencerInbox,
                inbox: inbox,
                outbox: outbox,
                rollupEventInbox: rollupEventInbox,
                challengeManager: challengeManager,
                rollupAdminLogic: address(rollupAdminLogic),
                rollupUserLogic: rollupUserLogic,
                validatorUtils: validatorUtils,
                validatorWalletCreator: validatorWalletCreator
            })
        );

        emit RollupCreated(
            address(rollup), address(inbox), address(proxyAdmin), address(sequencerInbox), address(bridge)
        );
        return address(rollup);
    }
}
