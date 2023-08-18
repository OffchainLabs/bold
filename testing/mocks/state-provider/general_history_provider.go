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
		var upTo l2stateprovider.Height
		if !upToHeight.IsNone() {
			upTo = upToHeight.Unwrap()
		} else {
			upTo = l2stateprovider.Height(s.levelZeroBlockEdgeHeight)
		}
		from := startHeights[0]
		states, err := s.statesUpTo(uint64(from), uint64(upTo), uint64(batch))
		if err != nil {
			return emptyCommit, err
		}
		return commitments.New(states)
	}

	// Next, computes the exact start point of where we need to execute
	// the machine from the inputs, and figures out in what increments we need to do so.
	fromMessageNumber := uint64(startHeights[0])
	machine, err := s.machineAtBlock(ctx, fromMessageNumber)
	if err != nil {
		return emptyCommit, err
	}
	machineStartIndex, err := s.computeMachineStartIndex()
	if err != nil {
		return emptyCommit, err
	}

	// We compute the stepwise increments we need for stepping through the machine.
	stepBy, err := s.computeStepIncrement(uint64(len(startHeights)))
	if err != nil {
		return emptyCommit, err
	}

	// We advance the machine to the index we need to start from.
	if err = machine.Step(machineStartIndex); err != nil {
		return emptyCommit, err
	}

	// Advance the machine to the requested start point.
	maxHeightForLevel := s.challengeLeafHeights[len(startHeights)-1]

	// Get the start and end points for the machine stepping.
	start := uint64(startHeights[len(startHeights)-1])
	var end uint64
	if upToHeight.IsNone() {
		end = maxHeightForLevel
	} else {
		end = uint64(upToHeight.Unwrap())
		if end > maxHeightForLevel {
			end = maxHeightForLevel
		}
	}

	if end < start {
		return emptyCommit, fmt.Errorf("invalid range: %d > %d", start, end)
	}

	// We step through the machine in our desired increments, and gather the
	// machine hashes along the way for the history commitment.
	leaves := make([]common.Hash, end-start+1)
	for i := start; i <= end; i++ {
		leaves = append(leaves, s.getMachineHash(machine, fromMessageNumber))
		if i >= end {
			// We don't need to step the machine to the next point because it won't be used.
			break
		}
		if err = machine.Step(stepBy); err != nil {
			return emptyCommit, err
		}
	}
	return commitments.New(leaves)
}

func (s *L2StateBackend) computeMachineStartIndex(
	claimHeights []l2stateprovider.ClaimHeight,
) (uint64, error) {
	if len(claimHeights) != len(s.challengeLeafHeights) {
		return 0, fmt.Errorf(
			"challenge heights length %d != challenge leaf heights length %d",
			len(claimHeights),
			len(s.challengeLeafHeights),
		)
	}
	startIndex := uint64(0)
	for i := 0; i < len(claimHeights)-1; i++ {
		startIndex += claimHeights[i].From * s.challengeLeafHeights[i]
	}
	startIndex += claimHeights[len(claimHeights)-1].From
	return startIndex, nil
}

func (s *L2StateBackend) computeStepIncrement(requestedChallengeLevel uint64) (uint64, error) {
	if requestedChallengeLevel >= uint64(len(s.challengeLeafHeights)) {
		return 0, fmt.Errorf(
			"requested challenge level %d >= challenge leaf heights length %d",
			requestedChallengeLevel,
			len(s.challengeLeafHeights),
		)
	}
	return s.challengeLeafHeights[requestedChallengeLevel], nil
}
