package assertions

import (
	"context"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/containers/option"
	pools "github.com/OffchainLabs/bold/solgen/go/assertionStakingPoolgen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type useAssertionPoolArgs struct {
	parentAssertion protocol.AssertionHash
	assertionHash   protocol.AssertionHash
}

func (m *Manager) useAssertionPool(
	ctx context.Context,
	args useAssertionPoolArgs,
) {
	assertionPool, err := m.getOrCreateAssertionPool(ctx, args)
	if err != nil {
		panic(err)
	}
}

type assertionStakingPool struct{}

// Get an assertion staking pool for the assertion we wish to post.
func (m *Manager) getOrCreateAssertionPool(
	ctx context.Context,
	args useAssertionPoolArgs,
) (*assertionStakingPool, error) {
	var poolAddr common.Address
	poolAddrOpt := m.checkAssertionPoolCreated(ctx, args)
	if poolAddrOpt.IsSome() {
		poolAddr = poolAddrOpt.Unwrap()
	} else {
		m.createAssertionStakingPool(ctx, args)
	}
	return nil, nil
}

func (m *Manager) createAssertionStakingPool(ctx context.Context, args useAssertionPoolArgs) common.Address {
	return common.Address{}
}

// Scan for any pools created since the latest confirmed assertion
// from the staking pool factory.
func (m *Manager) checkAssertionPoolCreated(
	ctx context.Context,
	args useAssertionPoolArgs,
) option.Option[common.Address] {
	poolFactory, err := pools.NewAssertionStakingPoolCreator(common.Address{}, m.backend)
	if err != nil {
		panic(err)
	}
	parent, err := m.chain.ReadAssertionCreationInfo(ctx, args.parentAssertion)
	if err != nil {
		panic(err)
	}
	// Keep filtering until we reach the threshold.
	filterOpts := &bind.FilterOpts{
		Start: parent.CreationBlock,
		End:   nil,
	}
	it, err := poolFactory.FilterNewAssertionPoolCreated(
		filterOpts,
		[]common.Address{m.rollupAddr},
		[][32]byte{
			args.assertionHash.Hash,
		},
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = it.Close(); err != nil {
			panic(err)
		}
	}()
	for it.Next() {
		if it.Error() != nil {
			panic(err)
		}
		return option.Some(it.Event.AssertionPool)
	}
	return option.None[common.Address]()
}

type monitorPoolCreatorArgs struct {
	parentAssertion protocol.AssertionHash
	assertionHash   protocol.AssertionHash
}

// Monitor any staking pools that we care about which reach the threshold.
// Should we block until the pool threshold is reached and not post more?
// There is only one canonical assertion branch, so yes we likely do have to wait.
func (m *Manager) monitorAssertionPool(
	ctx context.Context,
	args monitorPoolCreatorArgs,
) {
}
