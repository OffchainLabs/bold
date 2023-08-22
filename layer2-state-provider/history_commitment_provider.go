package l2stateprovider

import (
	"context"
	"errors"
	"fmt"

	"github.com/OffchainLabs/bold/containers/option"
	commitments "github.com/OffchainLabs/bold/state-commitments/history"
	"github.com/ethereum/go-ethereum/common"
)

var (
	emptyCommit = commitments.History{}
)

// MachineHashCollector defines a struct which can collect hashes from an Arbitrator machine
// at a block height, starting at a specific opcode index in the machine and stepping through it
// in increments of custom size. Along the way, it computes each machine hash at each step
// and outputs a list of these hashes at the end. This is an computationally expensive process
// that is best performed if machine hashes are cached after runs.
type MachineHashCollector interface {
	CollectMachineMashes(ctx context.Context, cfg *HashCollectorConfig) ([]common.Hash, error)
}

// HashCollectorConfig defines configuration options for a machine hash collector to
// step through an Arbitrator machine at a specific L2 message number, step through it in
// increments of `StepSize`, and collect the machine hashes at those steps into an output slice.
type HashCollectorConfig struct {
	// The WASM module root the machines should be a part of.
	WasmModuleRoot common.Hash
	// The L2 message number the machine corresponds to.
	MessageNumber Height
	StepHeights   []Height
	// The number of desired hashes to be collected.
	NumDesiredHashes uint64
	// The opcode index at which to start stepping through the machine at a message number.
	MachineStartIndex OpcodeIndex
	// The step size for stepping through the machine in order to collect its hashes.
	StepSize StepSize
}

// L2MessageStateCollector defines a struct which can obtain the machine hashes at each L2 message
// in a specified message range for a given batch index on Arbitrum.
type L2MessageStateCollector interface {
	L2MessageStatesUpTo(
		ctx context.Context,
		from Height,
		upTo option.Option[Height],
		batch Batch,
	) ([]common.Hash, error)
}

// HistoryCommitmentProvider computes history commitments from input parameters
// by loading Arbitrator machines for L2 state transitions. It can compute history commitments
// over ranges of opcodes at specified increments used for the BOLD protocol.
type HistoryCommitmentProvider struct {
	l2MessageStateCollector L2MessageStateCollector
	machineHashCollector    MachineHashCollector
	challengeLeafHeights    []Height
}

// NewHistoryCommitmentProvider --
func NewHistoryCommitmentProvider(
	l2MessageStateCollector L2MessageStateCollector,
	machineHashCollector MachineHashCollector,
	challengeLeafHeights []Height,
) *HistoryCommitmentProvider {
	return &HistoryCommitmentProvider{
		l2MessageStateCollector: l2MessageStateCollector,
		machineHashCollector:    machineHashCollector,
		challengeLeafHeights:    challengeLeafHeights,
	}
}

// A list of heights that have been validated to be non-empty
// and to be < the total number of challenge levels in the protocol.
type validatedStartHeights []Height

// HistoryCommitment computes a Merklelized commitment over a set of hashes
// at specified challenge levels. For block challenges, for example, this is a set
// of machine hashes corresponding each message in a range N to M.
//
// Usage:
//
//		historyCommitment, err := provider.HistoryCommitment(
//		   ctx,
//		   wasmModuleRoot, /* the wasm module root for the machine */
//		   Batch(1), /* the commitment is at L2 batch index 1 */
//		   []Height{
//	         0,   /* at L2 message index 0 /*
//	         100, /* at giant step 100 */
//	         20,  /* at big step 20 */
//	       },
//		   option.Some(Height(200)) /* up to small step 200 (opcode index 200) */
//		)
func (p *HistoryCommitmentProvider) HistoryCommitment(
	ctx context.Context,
	wasmModuleRoot common.Hash,
	batch Batch,
	startHeights []Height,
	upToHeight option.Option[Height],
) (commitments.History, error) {
	// Validate the input heights for correctness.
	validatedHeights, err := p.validateStartHeights(startHeights)
	if err != nil {
		return emptyCommit, err
	}
	// If the call is for message number ranges only, we get the hashes for
	// those states and return a commitment for them.
	fromMessageNumber := uint64(validatedHeights[0])
	if len(validatedHeights) == 1 {
		hashes, hashesErr := p.l2MessageStateCollector.L2MessageStatesUpTo(ctx, Height(fromMessageNumber), upToHeight, batch)
		if hashesErr != nil {
			return emptyCommit, hashesErr
		}
		return commitments.New(hashes)
	}

	// Next, computes the exact start point of where we need to execute
	// the machine from the inputs, and figures out in what increments we need to do so.
	machineStartIndex := p.computeMachineStartIndex(validatedHeights)

	// We compute the stepwise increments we need for stepping through the machine.
	stepSize, err := p.computeStepSize(validatedHeights)
	if err != nil {
		return emptyCommit, err
	}

	// Compute how many machine hashes we need to collect.
	numHashes, err := p.computeRequiredNumberOfHashes(validatedHeights, upToHeight)
	if err != nil {
		return emptyCommit, err
	}

	// Collect the machine hashes at the specified challenge level based on the values we computed.
	hashes, err := p.machineHashCollector.CollectMachineMashes(
		ctx,
		&HashCollectorConfig{
			WasmModuleRoot:    wasmModuleRoot,
			MessageNumber:     Height(fromMessageNumber),
			StepHeights:       validatedHeights[1:],
			NumDesiredHashes:  numHashes,
			MachineStartIndex: machineStartIndex,
			StepSize:          stepSize,
		},
	)
	if err != nil {
		return emptyCommit, err
	}
	return commitments.New(hashes)
}

// Computes the required number of hashes for a history commitment
// based on the requested heights and challenge level. The required number of hashes
// for a leaf commitment at each challenge level is a constant, so we can determine
// the desired challenge level from the input params and compute the total
// from there.
func (p *HistoryCommitmentProvider) computeRequiredNumberOfHashes(
	startHeights validatedStartHeights,
	upToHeight option.Option[Height],
) (uint64, error) {
	// Get the max number of hashes at the specified challenge level.
	// from the protocol constants.
	challengeLevel := len(startHeights) - 1
	maxHeightForLevel := p.challengeLeafHeights[challengeLevel]

	// Get the start height we want to use at the challenge level.
	start := startHeights[challengeLevel]

	var end Height
	if upToHeight.IsNone() {
		end = maxHeightForLevel
	} else {
		end = upToHeight.Unwrap()
		// If the end height is more than the allowed max, we return an error.
		// This scenario should not happen, and instead of silently truncating,
		// surfacing an error is the safest way of warning the operator
		// they are committing something invalid.
		if end > maxHeightForLevel {
			return 0, fmt.Errorf(
				"end %d was greater than max height for level %d",
				end,
				maxHeightForLevel,
			)
		}
	}
	if end < start {
		return 0, fmt.Errorf("invalid range: end %d was < start %d", end, start)
	}
	// The number of hashes is the difference between the start and end
	// requested heights, plus 1.
	return uint64(end-start) + 1, nil
}

// Figures out the actual opcode index we should move the machine to
// when we compute the history commitment. As there are different levels of challenge
// granularity, we have to do some math to figure out the correct index.
//
// Take, for example, that we have 5 challenge kinds:
//
// block    => over a range of L2 message hashes
// colossal => over a range of giant steps
// giant    => over a range of big steps
// big      => over a range of small steps
// small    => over a range of individual WAVM opcodes
//
// We only directly step through WAVM machines when in a subchallenge (starting at colossal),
// so we can ignore block challenges for this calculation.
//
// Next, take the following constants:
//
//	colossal_steps_per_giant_step = 16
//	giant_steps_per_big_step      = 8
//	big_steps_per_giant_step      = 4
//	small_steps_per_big_step      = 2
//
// Let's say we want to figure out the machine start opcode index for the following inputs:
//
//	colossal_step=4, giant_step=5, big_step=1, small_step=0
//
// We can compute the opcode index using the following algorithm for the example above.
//
//	  4 * (8 * 4 * 2)
//	+ 5 * (4 * 2)
//	+ 1 * (2)
//	+ 0
//	= 298
//
// This generalizes for any number of subchallenge levels into the algorithm below.
// It works by taking the sum of (each input * product of all challenge level height constants beneath its level).
// For the block challenge level, the machine start opcode index is always 0.
func (p *HistoryCommitmentProvider) computeMachineStartIndex(
	startHeights validatedStartHeights,
) OpcodeIndex {
	if len(startHeights) == 1 {
		return 0
	}
	// We ignore the block challenge level here.
	heights := startHeights[1:]
	leafHeights := p.challengeLeafHeights[1:]

	// Next, we compute the opcode index.
	opcodeIndex := uint64(0)
	idx := 1
	for _, height := range heights {
		total := uint64(1)
		for i := idx; i < len(leafHeights); i++ {
			total *= uint64(leafHeights[i])
		}
		opcodeIndex += total * uint64(height)
		idx += 1
	}
	return OpcodeIndex(opcodeIndex)
}

// Computes the step size for stepping through a machine at a block height. Each challenge level has a different
// amount of steps, so the total step size can be computed as a multiplication of all the
// next challenge levels needed.
func (p *HistoryCommitmentProvider) computeStepSize(startHeights validatedStartHeights) (StepSize, error) {
	challengeLevel := len(startHeights) - 1
	totalChallengeLevels := len(p.challengeLeafHeights)

	// The stepwise increment of the last challenge level is always one.
	if challengeLevel+1 == totalChallengeLevels {
		return 1, nil
	}
	// Otherwise, it is the multiplication of all the challenge leaf heights at the next
	// challenge levels.
	levels := p.challengeLeafHeights[challengeLevel+1:]
	total := uint64(1)
	for _, h := range levels {
		total *= uint64(h)
	}
	return StepSize(total), nil
}

// Validates a start heights input must be non-empty and have a max
// equal to the number of challenge levels.
func (p *HistoryCommitmentProvider) validateStartHeights(
	startHeights []Height,
) (validatedStartHeights, error) {
	if len(startHeights) == 0 {
		return nil, errors.New("must provide start heights to compute number of hashes")
	}
	challengeLevel := len(startHeights) - 1
	if challengeLevel >= len(p.challengeLeafHeights) {
		return nil, fmt.Errorf(
			"challenge level %d is out of range for challenge leaf heights %v",
			challengeLevel,
			p.challengeLeafHeights,
		)
	}
	return validatedStartHeights(startHeights), nil
}
