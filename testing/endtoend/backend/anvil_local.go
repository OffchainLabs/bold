// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package backend

import (
	"context"
	"time"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/testing/setup"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

var _ Backend = &AnvilLocal{}

type AnvilLocal struct {
	client    *ethclient.Client
	rpc       *rpc.Client
	setup     *setup.ChainSetup
	accounts  []*setup.TestAccount
	setupOpts []setup.Opt
	blockTime time.Duration
}

func NewAnvilLocal(ctx context.Context, opts ...setup.Opt) (*AnvilLocal, error) {
	a := &AnvilLocal{}
	c, err := rpc.DialContext(ctx, "http://localhost:8545")
	if err != nil {
		return nil, err
	}
	a.rpc = c
	a.blockTime = time.Second * 2
	a.setupOpts = opts
	a.client = ethclient.NewClient(c)
	if err := a.loadAccounts(); err != nil {
		return nil, err
	}
	return a, nil
}

// Load accounts from test mnemonic. These are not real accounts. Don't even try to use them.
func (a *AnvilLocal) loadAccounts() error {
	accounts := make([]*setup.TestAccount, 0)
	chainId, err := a.client.ChainID(context.Background())
	if err != nil {
		return err
	}
	for i := 0; i < 3; i++ {
		privKeyHex := hexutil.MustDecode(anvilPrivKeyHexStrings[i])
		privKey, err := crypto.ToECDSA(privKeyHex)
		if err != nil {
			return err
		}
		txOpts, err := bind.NewKeyedTransactorWithChainID(privKey, chainId)
		if err != nil {
			return err
		}
		accounts = append(accounts, &setup.TestAccount{
			PrivateKey:  privKey,
			AccountAddr: txOpts.From,
			TxOpts:      txOpts,
		})
	}
	a.accounts = accounts
	return nil
}

func (a *AnvilLocal) Start(ctx context.Context) error {
	// Advance blocks in the background.
	go func() {
		ticker := time.NewTicker(a.blockTime)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := a.MineBlocks(ctx, 1); err != nil {
					log.Info("failed to mine block: %v", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

// Client returns the ethclient associated with the backend.
func (a *AnvilLocal) Client() protocol.ChainBackend {
	return a.client
}

func (a *AnvilLocal) Accounts() []*bind.TransactOpts {
	accs := make([]*bind.TransactOpts, 0)
	for _, acc := range a.accounts {
		accs = append(accs, acc.TxOpts)
	}
	return accs
}

func (a *AnvilLocal) Commit() common.Hash {
	return common.Hash{}
}

func (a *AnvilLocal) DeployRollup(ctx context.Context) (*setup.RollupAddresses, error) {
	setp := &setup.ChainSetup{
		NumAccountsToGen: 3,
	}
	for _, opt := range a.setupOpts {
		opt(setp)
	}
	setp, err := setup.SetupStackFromBackend(setp, a.accounts, a.client)
	if err != nil {
		return nil, err
	}
	a.setup = setp
	return setp.Addrs, nil
}

// MineBlocks will call anvil to instantly mine n blocks.
func (a *AnvilLocal) MineBlocks(ctx context.Context, n uint64) error {
	return a.rpc.CallContext(ctx, nil, "anvil_mine", hexutil.EncodeUint64(n))
}

func (a *AnvilLocal) ContractAddresses() *setup.RollupAddresses {
	return a.setup.Addrs
}
