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
	if len(startHeights) == 0 {
		return emptyCommit, errors.New("must specify at least one start height")
	}
	// If the call is for message number ranges only, we get the hashes for
	// those states and return a commitment for them.
	if len(startHeights) == 1 {
		return s.blockHistoryCommitment(startHeights[0], upToHeight, batch)
	}

	// Loads a machine at a specific message number.
	fromMessageNumber := uint64(startHeights[0])
	machine, err := s.machineAtBlock(ctx, fromMessageNumber)
	if err != nil {
		return emptyCommit, err
	}

	// Next, computes the exact start point of where we need to execute
	// the machine from the inputs, and figures out in what increments we need to do so.
	machineStartIndex, err := s.computeMachineStartIndex(startHeights)
	if err != nil {
		return emptyCommit, err
	}

	// We compute the stepwise increments we need for stepping through the machine.
	stepBy, err := s.computeStepIncrement(startHeights)
	if err != nil {
		return emptyCommit, err
	}

	// We advance the machine to the index we need to start from.
	if err = machine.Step(machineStartIndex); err != nil {
		return emptyCommit, err
	}

	// Compute how many machine hashes we need to collect.
	numHashes, err := s.computeRequiredNumberOfHashes(startHeights, upToHeight)
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
	startHeights []l2stateprovider.Height,
	upToHeight option.Option[l2stateprovider.Height],
) (uint64, error) {
	challengeLevel := len(startHeights) - 1
	// Get the max number of hashes at the specified challenge level.
	// from the protocol constants.
	maxHeightForLevel := s.challengeLeafHeights[challengeLevel]

	// Get the start height we want to use at the challenge level.
	start := uint64(startHeights[challengeLevel])

	var end uint64
	if upToHeight.IsNone() {
		end = maxHeightForLevel
	} else {
		end = uint64(upToHeight.Unwrap())
		// If the end height is more than the allowed max, we truncate.
		if end > maxHeightForLevel {
			end = maxHeightForLevel
		}
	}
	if end < start {
		return 0, fmt.Errorf("invalid range: %d > %d", start, end)
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

func (s *L2StateBackend) computeMachineStartIndex(
	startHeights []l2stateprovider.Height,
) (uint64, error) {
	// if len(claimHeights) != len(s.challengeLeafHeights) {
	// 	return 0, fmt.Errorf(
	// 		"challenge heights length %d != challenge leaf heights length %d",
	// 		len(claimHeights),
	// 		len(s.challengeLeafHeights),
	// 	)
	// }
	// startIndex := uint64(0)
	// for i := 0; i < len(claimHeights)-1; i++ {
	// 	startIndex += claimHeights[i].From * s.challengeLeafHeights[i]
	// }
	// startIndex += claimHeights[len(claimHeights)-1].From
	// return startIndex, nil
	return 0, nil
}

func (s *L2StateBackend) computeStepIncrement(startHeights []l2stateprovider.Height) (uint64, error) {
	// if requestedChallengeLevel >= uint64(len(s.challengeLeafHeights)) {
	// 	return 0, fmt.Errorf(
	// 		"requested challenge level %d >= challenge leaf heights length %d",
	// 		requestedChallengeLevel,
	// 		len(s.challengeLeafHeights),
	// 	)
	// }
	// return s.challengeLeafHeights[requestedChallengeLevel], nil
	return 0, nil
}
