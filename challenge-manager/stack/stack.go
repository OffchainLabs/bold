package stack

import (
	"context"
	"time"

	"github.com/offchainlabs/bold/assertions"
	protocol "github.com/offchainlabs/bold/chain-abstraction"
	cm "github.com/offchainlabs/bold/challenge-manager"
	"github.com/offchainlabs/bold/challenge-manager/types"
	l2stateprovider "github.com/offchainlabs/bold/layer2-state-provider"
)

type stackParams struct {
	mode                                types.Mode
	name                                string
	pollInterval                        time.Duration
	postInterval                        time.Duration
	confInterval                        time.Duration
	avgBlockTime                        time.Duration
	trackChallengeParentAssertionHashes []string
	apiAddr                             string
	apiDBPath                           string
}

var defaultStackParams = stackParams{
	mode:                                types.MakeMode,
	name:                                "unnamed-challenge-manager",
	pollInterval:                        time.Minute,
	confInterval:                        time.Second * 10,
	postInterval:                        time.Hour,
	avgBlockTime:                        time.Second * 12,
	trackChallengeParentAssertionHashes: nil,
	apiAddr:                             "",
	apiDBPath:                           "",
}

// Opt is a functional option for configuring the default challenge manager.
type Opt func(*stackParams)

// WithMode sets the mode of the default challenge manager.
func WithMode(mode types.Mode) Opt {
	return func(p *stackParams) {
		p.mode = mode
	}
}

// WithName sets the name of the default challenge manager.
func WithName(name string) Opt {
	return func(p *stackParams) {
		p.name = name
	}
}

// WithPollingInterval sets the polling interval of the default challenge
// manager.
func WithPollingInterval(interval time.Duration) Opt {
	return func(p *stackParams) {
		p.pollInterval = interval
	}
}

// WithPostingInterval sets the posting interval of the default challenge
// manager.
func WithPostingInterval(interval time.Duration) Opt {
	return func(p *stackParams) {
		p.postInterval = interval
	}
}

// WithConfirmationInterval sets the confirmation interval of the default
// challenge manager.
func WithConfirmationInterval(interval time.Duration) Opt {
	return func(p *stackParams) {
		p.confInterval = interval
	}
}

// WithAverageBlockCreationTime sets the average block creation time of the
// default challenge manager.
func WithAverageBlockCreationTime(interval time.Duration) Opt {
	return func(p *stackParams) {
		p.avgBlockTime = interval
	}
}

// WithTrackChallengeParentAssertionHashes sets the track challenge parent
// assertion hashes of the default challenge manager.
func WithTrackChallengeParentAssertionHashes(hashes []string) Opt {
	return func(p *stackParams) {
		p.trackChallengeParentAssertionHashes = hashes
	}
}

// WithAPIEnabled sets the API address and database path of the default
// challenge manager.
func WithAPIEnabled(apiAddr, apiDBPath string) Opt {
	return func(p *stackParams) {
		p.apiAddr = apiAddr
		p.apiDBPath = apiDBPath
	}
}

// NewDefaultChallengeManager creates a new ChallengeManager and
// all of the dependencies wiring them together.
func NewDefaultChallengeManager(
	ctx context.Context,
	chain protocol.AssertionChain,
	provider l2stateprovider.Provider,
	opts ...Opt,
) (*cm.Manager, error) {
	params := defaultStackParams
	for _, o := range opts {
		o(&params)
	}

	// TODO(eljobe): Remove this comment once it's wired together correctly.
	// The dependencies are like this:
	// - protocol.AssertionChain : No Deps
	// - l2stateprovider.Provider : No Deps
	// - db.Database : apiDBPath
	// - watcher.Watcher : protocol.AssertionChain, cm.Manager, watcherInterval,
	//                     numBigSteps, db.Database, confInterval, avgBlockTime,
	//                     trackChallengeParentAssertionHashes
	// - apibackend.Backend : db.Database, protocol.AssertionChain, watcher.Watcher, cm.Manager!
	// - api/server.Server : apiAddr, apibackend.Backend
	// - assertions.Manager : protocol.AssertionChain, l2stateprovider.Provider, name,
	//                        db.Database, mode, avgBlockTime, confInterval, pollInterval,
	//	                      postInterval, rivalHandler!, cmAddress!
	// - cm.Manger : protocol.AssertionChain, l2stateprovider.Provider, assertions.Manager,
	//               rollupAddress, mode, trackChallengeParentAssertionHashes, name, apiAddr

	var amOpts []assertions.Opt
	amOpts = append(amOpts, assertions.WithAverageBlockCreationTime(params.avgBlockTime))
	amOpts = append(amOpts, assertions.WithConfirmationInterval(params.confInterval))
	amOpts = append(amOpts, assertions.WithPollingInterval(params.pollInterval))
	amOpts = append(amOpts, assertions.WithPostingInterval(params.postInterval))
	asm, err := assertions.NewManager(
		chain,
		provider,
		chain.Backend(),
		chain.RollupAddress(),
		params.name,
		nil,
		types.MakeMode,
		amOpts...,
	)
	if err != nil {
		return nil, err
	}

	var cmOpts []cm.Opt
	cmOpts = append(cmOpts, cm.WithMode(params.mode))
	cmOpts = append(cmOpts, cm.WithTrackChallengeParentAssertionHashes(params.trackChallengeParentAssertionHashes))
	cmOpts = append(cmOpts, cm.WithName(params.name))
	if params.apiAddr != "" {
		cmOpts = append(cmOpts, cm.WithAPIEnabled(params.apiAddr, params.apiDBPath))
	}
	return cm.New(ctx, chain, provider, asm, chain.RollupAddress(), cmOpts...)
}
