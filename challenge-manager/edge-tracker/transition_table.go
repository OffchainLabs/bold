// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package edgetracker

import (
	"github.com/OffchainLabs/bold/containers/fsm"
)

func newEdgeTrackerFsm(
	startState State,
	fsmOpts ...fsm.Opt[EdgeTrackerAction, State],
) (*fsm.Fsm[EdgeTrackerAction, State], error) {
	transitions := []*fsm.Event[EdgeTrackerAction, State]{
		{
			// Returns the tracker to the very beginning. Several states can cause
			// this, including challenge moves.
			Typ: EdgeBackToStart{},
			From: []State{
				EdgeBisecting,
				EdgeStarted,
				EdgeAtOneStepProof,
				EdgeAddingSubchallengeLeaf,
			},
			To: EdgeStarted,
		},
		{
			// The tracker will take some action if it has reached a one-step-proof
			// in a small step challenge.
			Typ:  EdgeHandleOneStepProof{},
			From: []State{EdgeStarted},
			To:   EdgeAtOneStepProof,
		},
		{
			// The tracker will add a subchallenge leaf to its edge's subchallenge.
			Typ:  EdgeOpenSubchallengeLeaf{},
			From: []State{EdgeStarted, EdgeAddingSubchallengeLeaf},
			To:   EdgeAddingSubchallengeLeaf,
		},
		// Challenge moves.
		{
			Typ:  EdgeBisect{},
			From: []State{EdgeStarted},
			To:   EdgeBisecting,
		},
		// Terminal state.
		{
			Typ:  EdgeConfirm{},
			From: []State{EdgeStarted, EdgeConfirmed, EdgeAtOneStepProof},
			To:   EdgeConfirmed,
		},
	}
	return fsm.New(startState, transitions, fsmOpts...)
}
