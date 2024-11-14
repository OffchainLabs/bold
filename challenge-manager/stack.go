package challengemanager

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/bold/api/db"
	"github.com/offchainlabs/bold/assertions"
	protocol "github.com/offchainlabs/bold/chain-abstraction"
	watcher "github.com/offchainlabs/bold/challenge-manager/chain-watcher"
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
	trackChallengeParentAssertionHashes []protocol.AssertionHash
	apiAddr                             string
	apiDBPath                           string
	enableHeadBlockSubscriptions        bool
	enableFastConfirmation              bool
	assertionManagerOverride            *assertions.Manager
}

var defaultStackParams = stackParams{
	mode:                                types.MakeMode,
	name:                                "unnamed-challenge-manager",
	pollInterval:                        time.Minute,
	postInterval:                        time.Hour,
	confInterval:                        time.Second * 10,
	avgBlockTime:                        time.Second * 12,
	trackChallengeParentAssertionHashes: nil,
	apiAddr:                             "",
	apiDBPath:                           "",
	enableHeadBlockSubscriptions:        false,
	enableFastConfirmation:              false,
	assertionManagerOverride:            nil,
}

// Opt is a functional option for configuring the default challenge manager.
type StackOpt func(*stackParams)

// WithMode sets the mode of the default challenge manager.
func StackWithMode(mode types.Mode) StackOpt {
	return func(p *stackParams) {
		p.mode = mode
	}
}

// WithName sets the name of the default challenge manager.
func StackWithName(name string) StackOpt {
	return func(p *stackParams) {
		p.name = name
	}
}

// WithPollingInterval sets the polling interval of the default challenge
// manager.
func StackWithPollingInterval(interval time.Duration) StackOpt {
	return func(p *stackParams) {
		p.pollInterval = interval
	}
}

// WithPostingInterval sets the posting interval of the default challenge
// manager.
func StackWithPostingInterval(interval time.Duration) StackOpt {
	return func(p *stackParams) {
		p.postInterval = interval
	}
}

// WithConfirmationInterval sets the confirmation interval of the default
// challenge manager.
func StackWithConfirmationInterval(interval time.Duration) StackOpt {
	return func(p *stackParams) {
		p.confInterval = interval
	}
}

// WithAverageBlockCreationTime sets the average block creation time of the
// default challenge manager.
func StackWithAverageBlockCreationTime(interval time.Duration) StackOpt {
	return func(p *stackParams) {
		p.avgBlockTime = interval
	}
}

// WithTrackChallengeParentAssertionHashes sets the track challenge parent
// assertion hashes of the default challenge manager.
func StackWithTrackChallengeParentAssertionHashes(hashes []string) StackOpt {
	return func(p *stackParams) {
		p.trackChallengeParentAssertionHashes = make([]protocol.AssertionHash, len(hashes))
		for i, h := range hashes {
			p.trackChallengeParentAssertionHashes[i] = protocol.AssertionHash{Hash: common.HexToHash(h)}
		}
	}
}

// WithAPIEnabled sets the API address and database path of the default
// challenge manager.
func StackWithAPIEnabled(apiAddr, apiDBPath string) StackOpt {
	return func(p *stackParams) {
		p.apiAddr = apiAddr
		p.apiDBPath = apiDBPath
	}
}

// WithHeadBlockSubscriptionsEnabled sets the enable head block subscriptions of
// the default challenge manager.
func StackWithHeadBlockSubscriptionsEnabled() StackOpt {
	return func(p *stackParams) {
		p.enableHeadBlockSubscriptions = true
	}
}

// WithFastConfirmationEnabled
func StackWithFastConfirmationEnabled() StackOpt {
	return func(p *stackParams) {
		p.enableFastConfirmation = true
	}
}

// OverrideAssertionManger can be used in tests to override the default
// assertion manager.
func OverrideAssertionManager(asm *assertions.Manager) StackOpt {
	return func(p *stackParams) {
		p.assertionManagerOverride = asm
	}
}

// NewDefaultChallengeManager creates a new ChallengeManager and
// all of the dependencies wiring them together.
func NewDefaultChallengeManager(
	ctx context.Context,
	chain protocol.AssertionChain,
	provider l2stateprovider.Provider,
	opts ...StackOpt,
) (*Manager, error) {
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
	// - api/server.Server : apbmaiAddr, apibackend.Backend
	// - assertions.Manager : protocol.AssertionChain, l2stateprovider.Provider, name,
	//                        db.Database, mode, avgBlockTime, confInterval, pollInterval,
	//	                      postInterval, rivalHandler!
	// - cm.Manger : protocol.AssertionChain, l2stateprovider.Provider, assertions.Manager,
	//               rollupAddress, mode, trackChallengeParentAssertionHashes, name, apiAddr

	// Create the api database.
	var apiDB db.Database
	if params.apiDBPath != "" {
		adb, err := db.NewDatabase(params.apiDBPath)
		if err != nil {
			return nil, err
		}
		apiDB = adb
	}

	// NOTE: This is effectively a costant that was being set unconditionally
	//       in the challenge manager's construtor. Now, I'm not sure what to
	//       do with it. It is sometimes overridden in tests.
	chainWatcherInterval := time.Millisecond * 500
	// Create the chain watcher.
	watcher, err := watcher.New(
		chain,
		provider,
		chainWatcherInterval,
		params.name,
		apiDB,
		params.confInterval,
		params.avgBlockTime,
		params.trackChallengeParentAssertionHashes,
	)
	if err != nil {
		return nil, err
	}

	// Create the assertions manager.
	var asm *assertions.Manager
	if params.assertionManagerOverride == nil {
		// Create the assertions manager.
		amOpts := []assertions.Opt{
			assertions.WithAverageBlockCreationTime(params.avgBlockTime),
			assertions.WithConfirmationInterval(params.confInterval),
			assertions.WithPollingInterval(params.pollInterval),
			assertions.WithPostingInterval(params.postInterval),
		}
		if apiDB != nil {
			amOpts = append(amOpts, assertions.WithAPIDB(apiDB))
		}
		if params.enableFastConfirmation {
			amOpts = append(amOpts, assertions.WithFastConfirmation())
		}
		asm, err = assertions.NewManager(
			chain,
			provider,
			params.name,
			params.mode,
			amOpts...,
		)
		if err != nil {
			return nil, err
		}
	} else {
		asm = params.assertionManagerOverride
	}

	// Create the challenge manager.
	cmOpts := []Opt{
		WithMode(params.mode),
		WithName(params.name),
	}
	if params.enableHeadBlockSubscriptions {
		cmOpts = append(cmOpts, WithHeadBlockSubscriptions())
	}
	if params.apiAddr != "" {
		cmOpts = append(cmOpts, WithAPIEnabled(params.apiAddr, apiDB))
	}
	return New(ctx, chain, provider, watcher, asm, cmOpts...)
}
