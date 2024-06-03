package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/containers/threadsafe"
	"github.com/OffchainLabs/bold/solgen/go/challengeV2gen"
	"github.com/OffchainLabs/bold/solgen/go/rollupgen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	defaultChanBufferSize    = 100
	defaultChainScanInterval = time.Second
)

var (
	assertionCreatedId               common.Hash
	recommendedL1YieldPerBlock       = big.NewInt(1)  // TODO: Tweak and set in wei.
	blockRefinementCreateGasCost     = uint64(530703) // TODO: Set these refinement costs.
	bigStepRefinementCreateGasCost   = uint64(425628)
	smallStepRefinementCreateGasCost = uint64(439739)
	blockBisectGasCost               = uint64(480051)
	bigStepBisectGasCost             = uint64(648411)
	smallStepBisectGasCost           = uint64(661328)
	confirmByOneStepProofGasCost     = uint64(865060)
	confirmEdgeByTimeGasCost         = uint64(113097)
)

func init() {
	rollupAbi, err := rollupgen.RollupCoreMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	assertionCreatedEvent, ok := rollupAbi.Events["AssertionCreated"]
	if !ok {
		panic("RollupCore ABI missing AssertionCreated event")
	}
	assertionCreatedId = assertionCreatedEvent.ID
}

type claimType uint8

const (
	assertionTyp claimType = iota
	edgeTyp
)

func (c claimType) String() string {
	if c == edgeTyp {
		return "edge"
	}
	return "assertion"
}

// TODO: Use a standard block number for rpc requests that need specific timings.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rpcClient, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}
	client := ethclient.NewClient(rpcClient)
	chainId, err := client.ChainID(ctx)
	if err != nil {
		panic(err)
	}
	rollupAddr := common.HexToAddress("0x99d322A49EeAb96fE51e26944D54824F3Ef6dedF")
	rollupBindings, err := rollupgen.NewRollupUserLogic(rollupAddr, client)
	if err != nil {
		panic(err)
	}
	chalManagerAddr, err := rollupBindings.ChallengeManager(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	chalManager, err := challengeV2gen.NewEdgeChallengeManager(chalManagerAddr, client)
	if err != nil {
		panic(err)
	}
	srv := &service{
		client:                           client,
		rollupBindings:                   rollupBindings,
		rollupAddr:                       rollupAddr,
		chalManagerAddr:                  chalManagerAddr,
		chalManager:                      chalManager,
		claimedChallengeAssertions:       make(chan *protocol.AssertionCreatedInfo, defaultChanBufferSize),
		confirmedAssertionsObserved:      make(chan common.Hash, defaultChanBufferSize),
		confirmedEdgesObserved:           make(chan *edge, defaultChainScanInterval),
		watchList:                        threadsafe.NewMap[common.Hash, claimType](),
		protocolGraphsByClaimedAssertion: make(map[common.Hash]*protocolGraph),
		gasPaymentRequests:               make(chan *gasPaymentRequest, defaultChanBufferSize),
		serviceFeePaymentRequests:        make(chan *serviceFeePaymentRequest, defaultChanBufferSize),
		recommendedL1YieldPerBlock:       recommendedL1YieldPerBlock,
		chainId:                          chainId,
		pendingPaymentsOutputDir:         filepath.Join(os.TempDir(), "payments"),
		chainScanInterval:                defaultChainScanInterval,
	}
	srv.run(ctx)
}

type service struct {
	rollupBindings                   *rollupgen.RollupUserLogic
	rollupAddr                       common.Address
	chalManagerAddr                  common.Address
	chalManager                      *challengeV2gen.EdgeChallengeManager
	client                           protocol.ChainBackend
	claimedChallengeAssertions       chan *protocol.AssertionCreatedInfo
	confirmedEdgesObserved           chan *edge
	confirmedAssertionsObserved      chan common.Hash
	watchList                        *threadsafe.Map[common.Hash, claimType]
	protocolGraphLock                sync.RWMutex
	protocolGraphsByClaimedAssertion map[common.Hash]*protocolGraph
	gasPaymentRequests               chan *gasPaymentRequest
	serviceFeePaymentRequests        chan *serviceFeePaymentRequest
	recommendedL1YieldPerBlock       *big.Int
	chainId                          *big.Int
	pendingPaymentsOutputDir         string
	chainScanInterval                time.Duration
}

func (s *service) run(ctx context.Context) {
	go receiveAsync(ctx, s.claimedChallengeAssertions, s.processAssertionsInChallenges)
	go receiveAsync(ctx, s.confirmedAssertionsObserved, s.processAssertionConfirmation)
	go receiveAsync(ctx, s.confirmedEdgesObserved, s.processEdgeConfirmation)
	go receiveAsync(ctx, s.gasPaymentRequests, s.processGasPaymentRequest)
	go receiveAsync(ctx, s.serviceFeePaymentRequests, s.processServiceFeePaymentRequest)
	s.scanChainEvents(ctx)
}

func (s *service) processAssertionsInChallenges(ctx context.Context, assertion *protocol.AssertionCreatedInfo) {
	// If the assertion is a not a descendant of the latest confirmed assertion, we ignore it.
	if !s.isDescendantOfLatestConfirmed(ctx, assertion.AssertionHash) {
		return
	}
	as, err := s.rollupBindings.GetAssertion(&bind.CallOpts{}, assertion.AssertionHash)
	if err != nil {
		panic(err)
	}
	status := protocol.AssertionStatus(as.Status)
	if status == protocol.AssertionConfirmed {
		// If the assertion is already confirmed, we notify listeners of this.
		s.confirmedAssertionsObserved <- assertion.AssertionHash
	} else if status == protocol.AssertionPending {
		// Otherwise, we add the assertion to the watchlist of hashes to check once they are confirmed.
		fmt.Println("Saw a pending assertion that is claimed in a challenge to put in the watchlist")
		s.watchList.Put(assertion.AssertionHash, assertionTyp)
	}
	_ = assertion
}

func (s *service) processAssertionConfirmation(ctx context.Context, confirmedAssertion common.Hash) {
	if !s.watchList.Has(confirmedAssertion) {
		// Ignore if the confirmed item is not in the watch list.
		fmt.Println("Ignored confirmed assertion not in the watchlist")
		return
	}
	fmt.Println("Got an assertion confirmed that was in the watchlist")
	s.executeReimbursement(ctx, reimbursementArgs{
		claimedAssertion: confirmedAssertion,
		itemTyp:          assertionTyp,
		challengeLvl:     protocol.NewBlockChallengeLevel(),
	})

	// Prune all items in the watchlist that are not descendants of the latest confirmed assertion.
	for hash, typ := range s.watchList.Copy() {
		if typ == edgeTyp {
			prevAssertionHash, err := s.chalManager.GetPrevAssertionHash(&bind.CallOpts{}, hash)
			if err != nil {
				panic(err)
			}
			if !s.isDescendantOfLatestConfirmed(ctx, prevAssertionHash) {
				s.watchList.Delete(prevAssertionHash)
				s.watchList.Delete(hash)
				if typ == edgeTyp {
					s.watchList.Delete(hash)
				}
			}
		} else {
			if !s.isDescendantOfLatestConfirmed(ctx, hash) {
				s.watchList.Delete(hash)
			}
		}
	}
}

func (s *service) processEdgeConfirmation(ctx context.Context, eg *edge) {
	// Check if we should be tracking the edge confirmation.
	predecessorAssertion, err := s.chalManager.GetPrevAssertionHash(&bind.CallOpts{}, eg.id)
	if err != nil {
		panic(err)
	}
	// Ignore an observed edge that has a predecessor assertion
	// that is NOT a descendant of the latest confirmed assertion.
	if !s.isDescendantOfLatestConfirmed(ctx, predecessorAssertion) {
		return
	}

	// Get the claimed assertion of the edge from the protocol graph. It is impossible to figure out
	// the claimed information of an assertion just from onchain data when given an edge,
	// so we need to loop over all our tracked protocol graphs to see if we have the edge locally.
	lvl := protocol.ChallengeLevel(eg.Level)
	found := false
	var claimedAssertion common.Hash
	s.protocolGraphLock.RLock()
	for claim, graph := range s.protocolGraphsByClaimedAssertion {
		for _, existingEdge := range graph.edgesByLevel[lvl] {
			if existingEdge.id == eg.id {
				found = true
				claimedAssertion = claim
				break
			}
		}
	}
	s.protocolGraphLock.RUnlock()
	if !found {
		fmt.Println("Not found in protocol graph", eg.Level, eg.StartHeight.Uint64(), eg.EndHeight.Uint64())
		return
	}
	s.executeReimbursement(ctx, reimbursementArgs{
		claimedAssertion: claimedAssertion,
		itemTyp:          edgeTyp,
		challengeLvl:     lvl,
		claimedEdge:      eg.id,
	})
}

func (s *service) isDescendantOfLatestConfirmed(ctx context.Context, assertionHash common.Hash) bool {
	latestConf, err := s.rollupBindings.LatestConfirmed(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	latestConfirmed, err := s.readAssertionCreationInfo(ctx, protocol.AssertionHash{Hash: latestConf})
	if err != nil {
		panic(err)
	}
	curr, err := s.readAssertionCreationInfo(ctx, protocol.AssertionHash{Hash: assertionHash})
	if err != nil {
		panic(err)
	}
	for {
		if curr.AssertionHash == latestConf || curr.ParentAssertionHash == latestConf {
			return true
		}
		// If the cursor's inbox max count is <= the latest confirmed's inbox max count,
		// and we still have not reached the "true" condition, then we can exit with false
		// as the assertion is not a descendant of the latest confirmed.
		if curr.InboxMaxCount.Cmp(latestConfirmed.InboxMaxCount) <= 0 {
			return false
		}
		parent, err := s.readAssertionCreationInfo(ctx, protocol.AssertionHash{Hash: curr.ParentAssertionHash})
		if err != nil {
			panic(err)
		}
		curr = parent
	}
}

type retryFunc[T any] func(...T) (T, error)

func retry[T any](f retryFunc[T], ctx context.Context, args ...T) T {
	var zeroVal T
	for {
		if ctx.Err() != nil {
			fmt.Println("Context timed out waiting to retry func, exiting...")
			return zeroVal
		}
		result, err := f(args...)
		if err != nil {
			fmt.Printf("error running func, will retry: %w\n", err)
			continue
		}
		return result
	}
}

func receiveAsync[T any](ctx context.Context, channel chan T, f func(context.Context, T)) {
	for {
		select {
		case item := <-channel:
			go f(ctx, item)
		case <-ctx.Done():
			return
		}
	}
}
