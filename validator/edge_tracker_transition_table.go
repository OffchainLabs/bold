package validator

import (
	"github.com/OffchainLabs/challenge-protocol-v2/util/fsm"
)

func newEdgeTrackerFsm(
	startState edgeTrackerState,
	fsmOpts ...fsm.Opt[edgeTrackerAction, edgeTrackerState],
) (*fsm.Fsm[edgeTrackerAction, edgeTrackerState], error) {
	transitions := []*fsm.Event[edgeTrackerAction, edgeTrackerState]{
		{
			// Returns the tracker to the very beginning. Several states can cause
			// this, including challenge moves.
			Typ: edgeBackToStart{},
			From: []edgeTrackerState{
				edgeBisecting,
				edgeStarted,
				edgeAtOneStepProof,
				edgeAddingSubchallengeLeaf,
			},
			To: edgeStarted,
		},
		// One-step-proof states.
		{
			// The tracker will take some action if it has reached a one-step-fork.
			Typ:  edgeHandleOneStepFork{},
			From: []edgeTrackerState{edgeStarted},
			To:   edgeAtOneStepFork,
		},
		{
			// The tracker will take some action if it has reached a one-step-proof
			// in a small step challenge.
			Typ:  edgeHandleOneStepProof{},
			From: []edgeTrackerState{edgeStarted},
			To:   edgeAtOneStepProof,
		},
		{
			// The tracker will add a subchallenge leaf to its edge's subchallenge.
			Typ:  edgeOpenSubchallengeLeaf{},
			From: []edgeTrackerState{edgeAtOneStepFork, edgeAddingSubchallengeLeaf},
			To:   edgeAddingSubchallengeLeaf,
		},
		// Challenge moves.
		{
			Typ:  edgeBisect{},
			From: []edgeTrackerState{edgeStarted},
			To:   edgeBisecting,
		},
		// Terminal state.
		{
			Typ:  edgeConfirm{},
			From: []edgeTrackerState{edgeStarted, edgeConfirmed, edgeAtOneStepProof},
			To:   edgeConfirmed,
		},
	}
	return fsm.New(startState, transitions, fsmOpts...)
}
