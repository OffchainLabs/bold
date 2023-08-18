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

func (s *L2StateBackend) HistoryCommitment(
	ctx context.Context,
	wasmModuleRoot common.Hash,
	inboxMaxCount uint64,
	messageNumberRange l2stateprovider.MessageNumberRange,
	claimHeights ...l2stateprovider.ClaimHeight,
) (commitments.History, error) {
	// If the call is for message number ranges only, we get the hashes for
	// those states and return a commitment for them.
	if len(claimHeights) == 0 {
		states, err := s.statesUpTo(messageNumberRange.From, messageNumberRange.To, inboxMaxCount)
		if err != nil {
			return emptyCommit, err
		}
		return commitments.New(states)
	}
	// Otherwise, the previous range for a claim height should have a length of one
	// for all specified claim heights, and we verify this condition.
	ranges := []l2stateprovider.ClaimHeight{
		{From: messageNumberRange.From, To: option.Some(messageNumberRange.To)},
	}
	// Validates ranges.
	ranges = append(ranges, claimHeights...)
	for i := 1; i < len(ranges)-1; i++ {
		prevRange := ranges[i-1]
		prevFrom := prevRange.From
		prevTo := prevRange.To
		if prevTo.IsNone() {
			return emptyCommit, ErrInvalidRange
		}
		if prevTo.Unwrap() != prevFrom+1 {
			return emptyCommit, fmt.Errorf(
				"%w: %d != %d+1",
				ErrNotOneStepFork,
				prevTo.Unwrap(),
				prevFrom,
			)
		}
	}

	// Next, computes the exact start point of where we need to execute
	// the machine from the inputs, and figures out in what increments we need to do so.
	machine, err := s.machineAtBlock(ctx, messageNumberRange.From)
	if err != nil {
		return emptyCommit, err
	}
	machineStartIndex, err := s.computeMachineStartIndex(claimHeights)
	if err != nil {
		return emptyCommit, err
	}

	// We compute the stepwise increments we need for stepping through the machine.
	stepBy, err := s.computeStepIncrement(uint64(len(claimHeights)))
	if err != nil {
		return emptyCommit, err
	}

	// We advance the machine to the index we need to start from.
	if err = machine.Step(machineStartIndex); err != nil {
		return emptyCommit, err
	}

	requestedRange := claimHeights[len(claimHeights)-1]

	// Advance the machine to the requested start point.
	maxHeightForLevel := s.challengeLeafHeights[len(claimHeights)-1]

	// Get the start and end points for the machine stepping.
	start := requestedRange.From
	var end uint64
	if requestedRange.To.IsNone() {
		end = s.challengeLeafHeights[len(claimHeights)-1]
	} else {
		end = requestedRange.To.Unwrap()
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
		leaves = append(leaves, s.getMachineHash(machine, messageNumberRange.From))
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

func (s *L2StateBackend) HistoryCommitmentV2(
	ctx context.Context,
	wasmModuleRoot common.Hash,
	batch l2stateprovider.Batch,
	startHeights []l2stateprovider.Height,
	upToHeight option.Option[l2stateprovider.Height],
) (commitments.History, error) {
	return commitments.History{}, errors.New("unimplemented")
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
