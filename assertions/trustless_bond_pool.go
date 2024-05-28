package assertions

import (
	"context"
	"math/big"
	"time"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/containers/option"
	pools "github.com/OffchainLabs/bold/solgen/go/assertionStakingPoolgen"
	"github.com/OffchainLabs/bold/solgen/go/mocksgen"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type useAssertionPoolArgs struct {
	parentAssertion protocol.AssertionHash
	assertionHash   protocol.AssertionHash
}

func (m *Manager) useAssertionPool(
	ctx context.Context,
	args useAssertionPoolArgs,
) protocol.Assertion {
	poolFactory, err := pools.NewAssertionStakingPoolCreator(
		m.poolingConfig.AssertionPoolCreatorFactoryAddr,
		m.backend,
	)
	if err != nil {
		panic(err)
	}
	assertionPool, err := m.getOrCreateAssertionPool(ctx, poolFactory, args)
	if err != nil {
		panic(err)
	}
	// Max parameter as a config into how much to deposit into the pool.
	assertionPool.depositIntoPool(ctx, args)

	// After this, we monitor the pool until it is ready to be posted.
	assertionPool.waitUntilFunded(ctx, poolFactory, monitorPoolCreatorArgs{})

	// Then, we can trigger the posting of the assertion.
	return assertionPool.postAssertionToPool(ctx, poolFactory, monitorPoolCreatorArgs{})
}

// Get an assertion staking pool for the assertion we wish to post.
func (m *Manager) getOrCreateAssertionPool(
	ctx context.Context,
	factory *pools.AssertionStakingPoolCreator,
	args useAssertionPoolArgs,
) (*assertionStakingPool, error) {
	var poolCreation createdPool
	poolOpt := m.checkAssertionPoolCreated(ctx, factory, args)
	if poolOpt.IsSome() {
		poolCreation = poolOpt.Unwrap()
	} else {
		poolCreation = m.createAssertionStakingPool(ctx, factory, args)
	}
	pool, err := pools.NewAssertionStakingPool(poolCreation.address, m.backend)
	if err != nil {
		return nil, err
	}
	return &assertionStakingPool{
		cfg:            m.poolingConfig,
		addr:           poolCreation.address,
		assertionHash:  poolCreation.assertionHash,
		pool:           pool,
		createdAtBlock: poolCreation.createdAtBlock,
	}, nil
}

type createdPool struct {
	assertionHash  common.Hash
	address        common.Address
	createdAtBlock uint64
}

// Scan for any pools created since the latest confirmed assertion
// from the staking pool factory.
func (m *Manager) checkAssertionPoolCreated(
	ctx context.Context,
	factory *pools.AssertionStakingPoolCreator,
	args useAssertionPoolArgs,
) option.Option[createdPool] {
	parent, err := m.chain.ReadAssertionCreationInfo(ctx, args.parentAssertion)
	if err != nil {
		panic(err)
	}
	var query = ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(parent.CreationBlock),
		ToBlock:   nil,
		Addresses: []common.Address{m.rollupAddr},
		Topics: [][]common.Hash{
			{args.assertionHash.Hash},
		},
	}
	logs, err := m.backend.FilterLogs(ctx, query)
	if err != nil {
		panic(err)
	}
	if len(logs) != 1 {
		return option.None[createdPool]()
	}
	poolDetails, err := factory.ParseNewAssertionPoolCreated(logs[0])
	if err != nil {
		panic(err)
	}
	return option.Some(createdPool{
		assertionHash:  poolDetails.AssertionHash,
		address:        poolDetails.AssertionPool,
		createdAtBlock: poolDetails.Raw.BlockNumber,
	})
}

func (m *Manager) createAssertionStakingPool(
	ctx context.Context,
	factory *pools.AssertionStakingPoolCreator,
	args useAssertionPoolArgs,
) createdPool {
	// TODO: Do this through the chain abstraction...
	// Get the receipt and the address of the new pool.
	// factory.CreatePool(m.poolingConfig.PoolingTxOpts, m.rollupAddr, args.assertionHash.Hash)
	return createdPool{
		assertionHash: args.assertionHash.Hash,
		// address: ,
	}
}

type assertionStakingPool struct {
	cfg            *AssertionPoolingConfig
	backend        protocol.ChainBackend
	addr           common.Address
	assertionHash  common.Hash
	pool           *pools.AssertionStakingPool
	createdAtBlock uint64
}

type monitorPoolCreatorArgs struct {
	parentAssertion protocol.AssertionHash
	assertionHash   protocol.AssertionHash
}

func (p *assertionStakingPool) depositIntoPool(
	ctx context.Context,
	args useAssertionPoolArgs,
) {
	gweiToDeposit := new(big.Int).SetUint64(p.cfg.MaxGweiToPool)
	gweiToWei := big.NewInt(1e9) // 10^9
	weiToDeposit := new(big.Int).Mul(gweiToDeposit, gweiToWei)
	_, _ = p.pool.DepositIntoPool(p.cfg.PoolingTxOpts, weiToDeposit)
}

// Monitor any staking pools that we care about which reach the threshold.
// Should we block until the pool threshold is reached and not post more?
// There is only one canonical assertion branch, so yes we likely do have to wait.
func (p *assertionStakingPool) waitUntilFunded(
	ctx context.Context,
	factory *pools.AssertionStakingPoolCreator,
	args monitorPoolCreatorArgs,
) {
	fromBlock := p.createdAtBlock
	latestBlock, err := p.backend.HeaderByNumber(ctx, nil) // TODO: Get desired block number.
	if err != nil {
		panic(err)
	}
	tokenAddr, err := p.pool.StakeToken(&bind.CallOpts{Context: ctx}) // TODO: Get desired block number.
	if err != nil {
		panic(err)
	}
	stakeToken, err := mocksgen.NewTestWETH9(tokenAddr, p.backend) // TODO: Do not use the mock here, just use an ierc20 binding.
	if err != nil {
		panic(err)
	}
	bal, err := stakeToken.BalanceOf(&bind.CallOpts{Context: ctx}, p.addr) // TODO: Get desired block number.
	if err != nil {
		panic(err)
	}

	// If balance is already enough, return.
	_ = bal

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if !latestBlock.Number.IsUint64() {
				log.Error("latest block header number is not a uint64")
				continue
			}
			toBlock := latestBlock.Number.Uint64()
			if fromBlock == toBlock {
				continue
			}
			// Store the parent assertion details and use that to determine if funded.
			// Scan for all deposits into the pool until the base stake threshold is reached.
			// Use the base stake from the parent assertion as the factor here.
			filterOpts := &bind.FilterOpts{
				Context: ctx,
				Start:   p.createdAtBlock,
				End:     nil,
			}
			it, err := p.pool.FilterStakeDeposited(filterOpts, nil)
			if err != nil {
				panic(err)
			}
			for it.Next() {
				// Get the balance after the deposit at that block number.
				// If balance reached...
				if true {
					return
				}
			}
		case <-ctx.Done():
			// TODO: Return an error.
			return
		}
	}
}

func (p *assertionStakingPool) postAssertionToPool(
	ctx context.Context,
	factory *pools.AssertionStakingPoolCreator,
	args monitorPoolCreatorArgs,
) protocol.Assertion {
	// p.pool.CreateAssertion(p.cfg.PoolingTxOpts, args.Hash)
	return nil
}
