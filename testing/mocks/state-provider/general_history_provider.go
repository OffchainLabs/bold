package stateprovider

import (
	"context"
	"errors"
	"fmt"

	"github.com/OffchainLabs/bold/containers/option"
	l2stateprovider "github.com/OffchainLabs/bold/layer2-state-provider"
	commitments "github.com/OffchainLabs/bold/state-commitments/history"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrNotOneStepFork = errors.New("height range is not a one step fork")
	ErrInvalidRange   = errors.New("invalid height range")
	emptyCommit       = commitments.History{}
)

// A list of heights that have been validated to be non-empty
// and to be < the total number of challenge levels in the protocol.
type validatedStartHeights []l2stateprovider.Height

// HistoryCommitment computes a Merklelized commitment over a set of hashes
// at specified challenge levels.
// For block challenges, for example, this is a set of machine hashes corresponding
// each message in a range N to M.
func (s *L2StateBackend) HistoryCommitment(
	ctx context.Context,
	wasmModuleRoot common.Hash,
	batch l2stateprovider.Batch,
	startHeights []l2stateprovider.Height,
	upToHeight option.Option[l2stateprovider.Height],
) (commitments.History, error) {
	validatedHeights, err := s.validateStartHeights(startHeights)
	if err != nil {
		return emptyCommit, err
	}
	// If the call is for message number ranges only, we get the hashes for
	// those states and return a commitment for them.
	if len(validatedHeights) == 1 {
		return s.blockHistoryCommitment(validatedHeights[0], upToHeight, batch)
	}

	// Loads a machine at a specific message number.
	fromMessageNumber := uint64(validatedHeights[0])
	machine, err := s.machineAtBlock(ctx, fromMessageNumber)
	if err != nil {
		return emptyCommit, err
	}

	// Next, computes the exact start point of where we need to execute
	// the machine from the inputs, and figures out in what increments we need to do so.
	machineStartIndex := s.computeMachineStartIndex(validatedHeights)

	// We compute the stepwise increments we need for stepping through the machine.
	stepBy, err := s.computeStepIncrement(validatedHeights)
	if err != nil {
		return emptyCommit, err
	}

	// We advance the machine to the index we need to start from.
	if err = machine.Step(machineStartIndex); err != nil {
		return emptyCommit, err
	}

	// Compute how many machine hashes we need to collect.
	numHashes, err := s.computeRequiredNumberOfHashes(validatedHeights, upToHeight)
	if err != nil {
		return emptyCommit, err
	}

	// We step through the machine in our desired increments, and gather the
	// machine hashes along the way for the history commitment.
	hashes := make([]common.Hash, numHashes)
	hashes = append(hashes, s.getMachineHash(machine, fromMessageNumber))
	for i := uint64(0); i < numHashes; i++ {
		if err = machine.Step(stepBy); err != nil {
			return emptyCommit, err
		}
		hashes = append(hashes, s.getMachineHash(machine, fromMessageNumber))
	}
	return commitments.New(hashes)
}

// Computes the required number of hashes for a history commitment
// based on the requested heights and challenge level.
func (s *L2StateBackend) computeRequiredNumberOfHashes(
	startHeights validatedStartHeights,
	upToHeight option.Option[l2stateprovider.Height],
) (uint64, error) {
	// Get the max number of hashes at the specified challenge level.
	// from the protocol constants.
	challengeLevel := len(startHeights) - 1
	maxHeightForLevel := s.challengeLeafHeights[challengeLevel]

	// Get the start height we want to use at the challenge level.
	start := uint64(startHeights[challengeLevel])

	var end uint64
	if upToHeight.IsNone() {
		end = maxHeightForLevel
	} else {
		end = uint64(upToHeight.Unwrap())
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
	return (end - start) + 1, nil
}

// Computes a block history commitment from a start height to a specified
// height up to the required batch number.
func (s *L2StateBackend) blockHistoryCommitment(
	from l2stateprovider.Height,
	to option.Option[l2stateprovider.Height],
	batch l2stateprovider.Batch,
) (commitments.History, error) {
	var upTo l2stateprovider.Height
	if !to.IsNone() {
		upTo = to.Unwrap()
	} else {
		blockChallengeLeafHeight := s.challengeLeafHeights[0]
		upTo = l2stateprovider.Height(blockChallengeLeafHeight)
	}
	states, err := s.statesUpTo(uint64(from), uint64(upTo), uint64(batch))
	if err != nil {
		return emptyCommit, err
	}
	return commitments.New(states)
}

// Figure out the actual opcode index we should move the machine to
// when we compute the history commitment. As there are different levels of challenge
// granularity, we have to do some math to figure out the correct index.
// Take, for example:
//
// lvl2_items_per_lvl1_step = 2
// lvl3_items_per_lvl2_step = 4
// lvl4_items_per_lvl3_step = 8
//
// This means there are 2 lvl2 items per lvl1 step, 4 lvl3 items per lvl2 step,
// and 8 lvl4 items per lvl3 step in a challenge.
//
// # Let's say we want to compute the actual opcode index for start heights
//
// [lvl1_start=2, lvl2_start=3, lvl3_start=4]
//
// We can compute the opcode index using the following algorithm for the example above.
//
//			2 * (4 * 8) = 64
//		  + 3 * (8)     = 24
//	   + 4           = opcode at index 92
//
// This generalizes for any number of subchallenge levels into the algorithm below.
func (s *L2StateBackend) computeMachineStartIndex(
	startHeights validatedStartHeights,
) uint64 {
	if len(startHeights) == 1 {
		return 0
	}
	// We ignore the block challenge level here.
	heights := startHeights[1:]
	leafHeights := s.challengeLeafHeights[1:]

	// Next, we compute the opcode index.
	opcodeIndex := uint64(0)
	idx := 1
	// TODO: Handle height 0.
	for _, height := range heights {
		total := uint64(1)
		for i := idx; i < len(leafHeights); i++ {
			total *= leafHeights[i]
		}
		opcodeIndex += total * uint64(height)
		idx += 1
	}
	return opcodeIndex
}

func (s *L2StateBackend) computeStepIncrement(startHeights validatedStartHeights) (uint64, error) {
	challengeLevel := len(startHeights) - 1
	totalChallengeLevels := len(s.challengeLeafHeights)

	// The stepwise increment of the last challenge level is always one.
	if challengeLevel+1 == totalChallengeLevels {
		return 1, nil
	}
	return s.challengeLeafHeights[challengeLevel+1], nil
}

// Validates a start heights input must be non-empty and have a max
// equal to the number of challenge levels.
func (s *L2StateBackend) validateStartHeights(
	startHeights []l2stateprovider.Height,
) (validatedStartHeights, error) {
	if len(startHeights) == 0 {
		return nil, errors.New("must provide start heights to compute number of hashes")
	}
	challengeLevel := len(startHeights) - 1
	if challengeLevel >= len(s.challengeLeafHeights) {
		return nil, fmt.Errorf(
			"challenge level %d is out of range for challenge leaf heights %v",
			challengeLevel,
			s.challengeLeafHeights,
		)
	}
	return validatedStartHeights(startHeights), nil
}
