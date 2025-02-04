package endtoend

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	protocol "github.com/offchainlabs/bold/chain-abstraction"
	solimpl "github.com/offchainlabs/bold/chain-abstraction/sol-implementation"
	cm "github.com/offchainlabs/bold/challenge-manager"
	"github.com/offchainlabs/bold/challenge-manager/types"
	retry "github.com/offchainlabs/bold/runtime"
	"github.com/offchainlabs/bold/solgen/go/challengeV2gen"
	"github.com/offchainlabs/bold/solgen/go/mocksgen"
	"github.com/offchainlabs/bold/solgen/go/rollupgen"
	challenge_testing "github.com/offchainlabs/bold/testing"
	"github.com/offchainlabs/bold/testing/endtoend/backend"
	statemanager "github.com/offchainlabs/bold/testing/mocks/state-provider"
	"github.com/offchainlabs/bold/testing/setup"
	"github.com/stretchr/testify/require"
)

func TestEndToEnd_HonestValidatorIsDelegatedStaker(t *testing.T) {
	neutralCtx, neutralCancel := context.WithCancel(context.Background())
	defer neutralCancel()
	evilCtx, evilCancel := context.WithCancel(context.Background())
	defer evilCancel()
	honestCtx, honestCancel := context.WithCancel(context.Background())
	defer honestCancel()

	protocolCfg := defaultProtocolParams()
	protocolCfg.challengePeriodBlocks = 40
	timeCfg := defaultTimeParams()
	timeCfg.blockTime = time.Second
	inboxCfg := defaultInboxParams()

	challengeTestingOpts := []challenge_testing.Opt{
		challenge_testing.WithConfirmPeriodBlocks(protocolCfg.challengePeriodBlocks),
		challenge_testing.WithLayerZeroHeights(&protocolCfg.layerZeroHeights),
		challenge_testing.WithNumBigStepLevels(protocolCfg.numBigStepLevels),
	}
	deployOpts := []setup.Opt{
		setup.WithMockBridge(),
		setup.WithMockOneStepProver(),
		setup.WithNumAccounts(4),
		setup.WithChallengeTestingOpts(challengeTestingOpts...),
		setup.WithNumFundedAccounts(3),
	}

	simBackend, err := backend.NewSimulated(timeCfg.blockTime, deployOpts...)
	require.NoError(t, err)
	bk := simBackend

	rollupAddr, err := bk.DeployRollup(neutralCtx, challengeTestingOpts...)
	require.NoError(t, err)

	require.NoError(t, bk.Start(neutralCtx))

	accounts := bk.Accounts()
	bk.Commit()

	rollupUserBindings, err := rollupgen.NewRollupUserLogic(rollupAddr.Rollup, bk.Client())
	require.NoError(t, err)
	bridgeAddr, err := rollupUserBindings.Bridge(&bind.CallOpts{})
	require.NoError(t, err)
	dataHash := common.Hash{1}
	enqueueSequencerMessageAsExecutor(
		t, accounts[0], rollupAddr.UpgradeExecutor, bk.Client(), bridgeAddr, seqMessage{
			dataHash:                 dataHash,
			afterDelayedMessagesRead: big.NewInt(1),
			prevMessageCount:         big.NewInt(1),
			newMessageCount:          big.NewInt(2),
		},
	)

	baseStateManagerOpts := []statemanager.Opt{
		statemanager.WithNumBatchesRead(inboxCfg.numBatchesPosted),
		statemanager.WithLayerZeroHeights(&protocolCfg.layerZeroHeights, protocolCfg.numBigStepLevels),
	}
	honestStateManager, err := statemanager.NewForSimpleMachine(t, baseStateManagerOpts...)
	require.NoError(t, err)

	shp := &simpleHeaderProvider{b: bk, chs: make([]chan<- *gethtypes.Header, 0)}
	shp.Start(neutralCtx)

	baseStackOpts := []cm.StackOpt{
		cm.StackWithMode(types.MakeMode),
		cm.StackWithPollingInterval(timeCfg.assertionScanningInterval),
		cm.StackWithPostingInterval(timeCfg.assertionPostingInterval),
		cm.StackWithAverageBlockCreationTime(timeCfg.blockTime),
		cm.StackWithConfirmationInterval(timeCfg.assertionConfirmationAttemptInterval),
		cm.StackWithMinimumGapToParentAssertion(0),
		cm.StackWithHeaderProvider(shp),
	}

	name := "honest"

	// Ensure the honest validator is a generated account that has no erc20 token balance,
	// but has some ETH to pay for gas costs of BoLD. We ensure that the honest validator
	// is not initially staked, and that the actual address that will be funding the honest
	// validator has enough funds.
	fundsCustodianOpts := accounts[1]         // The 1st account should be the funds custodian.
	honestTxOpts := accounts[len(accounts)-1] // The last account should not be funded with stake token.

	txOpts := accounts[1]
	//nolint:gocritic
	honestOpts := append(
		baseStackOpts,
		cm.StackWithName(name),
		cm.StackWithDelegatedStaking(), // Enable delegated staking for the honest validator only.
		cm.StackWithoutAutoAllowanceApproval(),
		cm.StackWithoutAutoDeposit(),
	)
	// Ensure the funds custodian is the withdrawal address for the honest validator.
	honestChain := setupAssertionChain(
		t,
		honestCtx,
		bk.Client(),
		rollupAddr.Rollup,
		txOpts,
		solimpl.WithCustomWithdrawalAddress(fundsCustodianOpts.From),
	)

	// Ensure the honest validator is not yet staked.
	isStaked, err := honestChain.IsStaked(context.Background())
	require.NoError(t, err)
	require.False(t, isStaked)

	chalManagerAddr := honestChain.SpecChallengeManager().Address()
	cmBindings, err := challengeV2gen.NewEdgeChallengeManager(chalManagerAddr, bk.Client())
	require.NoError(t, err)
	stakeToken, err := cmBindings.StakeToken(&bind.CallOpts{})
	require.NoError(t, err)
	requiredStake, err := honestChain.RollupCore().BaseStake(&bind.CallOpts{})
	require.NoError(t, err)

	tokenBindings, err := mocksgen.NewTestWETH9(stakeToken, bk.Client())
	require.NoError(t, err)

	bal, err := tokenBindings.BalanceOf(&bind.CallOpts{}, honestTxOpts.From)
	require.NoError(t, err)
	require.True(t, bal.Cmp(requiredStake) < 0) // Ensure honest validator does not have enough stake token balance.

	balCustodian, err := tokenBindings.BalanceOf(&bind.CallOpts{}, fundsCustodianOpts.From)
	require.NoError(t, err)
	require.True(t, balCustodian.Cmp(requiredStake) >= 0) // Ensure funds custodian DOES have enough stake token balance.

	assertionDivergenceHeight := uint64(1)
	assertionBlockHeightDifference := int64(1)

	honestManager, err := cm.NewChallengeStack(honestChain, honestStateManager, honestOpts...)
	require.NoError(t, err)

	machineDivergenceStep := uint64(1)
	//nolint:gocritic
	evilStateManagerOpts := append(
		baseStateManagerOpts,
		statemanager.WithMachineDivergenceStep(machineDivergenceStep),
		statemanager.WithBlockDivergenceHeight(assertionDivergenceHeight),
		statemanager.WithDivergentBlockHeightOffset(assertionBlockHeightDifference),
	)
	evilStateManager, err := statemanager.NewForSimpleMachine(t, evilStateManagerOpts...)
	require.NoError(t, err)

	// Honest validator has index 1 in the accounts slice, as 0 is admin, so
	// evil ones should start at 2.
	evilTxOpts := accounts[2]
	//nolint:gocritic
	evilOpts := append(
		baseStackOpts,
		cm.StackWithName("evil"),
	)
	evilChain := setupAssertionChain(t, evilCtx, bk.Client(), rollupAddr.Rollup, evilTxOpts)
	evilManager, err := cm.NewChallengeStack(evilChain, evilStateManager, evilOpts...)
	require.NoError(t, err)

	honestManager.Start(honestCtx)
	evilManager.Start(evilCtx)

	t.Run("expects honest validator to win challenge", func(t *testing.T) {
		chainId, err := bk.Client().ChainID(honestCtx)
		require.NoError(t, err)
		// Wait until a challenged assertion is confirmed by time.
		var confirmed bool
		for neutralCtx.Err() == nil && !confirmed {
			var i *rollupgen.RollupCoreAssertionConfirmedIterator
			i, err = retry.UntilSucceeds(neutralCtx, func() (*rollupgen.RollupCoreAssertionConfirmedIterator, error) {
				return honestChain.RollupCore().FilterAssertionConfirmed(nil, nil)
			})
			require.NoError(t, err)
			for i.Next() {
				creationInfo, err2 := evilChain.ReadAssertionCreationInfo(evilCtx, protocol.AssertionHash{Hash: i.Event.AssertionHash})
				require.NoError(t, err2)

				var parent rollupgen.AssertionNode
				parent, err = retry.UntilSucceeds(neutralCtx, func() (rollupgen.AssertionNode, error) {
					return honestChain.RollupCore().GetAssertion(&bind.CallOpts{Context: neutralCtx}, creationInfo.ParentAssertionHash.Hash)
				})
				require.NoError(t, err)

				tx, _, err2 := bk.Client().TransactionByHash(neutralCtx, creationInfo.TransactionHash)
				require.NoError(t, err2)
				sender, err2 := gethtypes.Sender(gethtypes.NewCancunSigner(chainId), tx)
				require.NoError(t, err2)
				honestConfirmed := sender == txOpts.From

				isChallengeChild := parent.FirstChildBlock > 0 && parent.SecondChildBlock > 0
				if !isChallengeChild {
					// Assertion must be a challenge child.
					continue
				}
				// We expect the honest party to have confirmed it.
				if !honestConfirmed {
					t.Fatal("Evil party confirmed the assertion by challenge win")
				}
				confirmed = true
				break
			}
			time.Sleep(500 * time.Millisecond) // Don't spam the backend.
		}
		// Once the honest, claimed assertion in the challenge is confirmed by time, we win the test.
		t.Log("Assertion was confirmed by time")
	})
}
