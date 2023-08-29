// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package edgetracker

import (
	"fmt"
)

// State defines a finite state machine that aids
// in deciding a challenge edge tracker's actions.
type State uint8

const (
	// Start state of 0 can never happen to avoid silly mistakes with default Go values.
	_ State = iota
	// The start state of the tracker.
	EdgeStarted
	// The edge being tracked is at a one step proof.
	EdgeAtOneStepProof
	// The tracker is adding a subchallenge leaf on the edge's subchallenge.
	EdgeAddingSubchallengeLeaf
	// The tracker is attempting a bisection move.
	EdgeBisecting
	// Terminal state
	EdgeConfirmed
)

// String turns an edge tracker state into a readable string.
func (s State) String() string {
	switch s {
	case EdgeStarted:
		return "started"
	case EdgeAtOneStepProof:
		return "one_step_proof"
	case EdgeAddingSubchallengeLeaf:
		return "adding_subchallenge_leaf"
	case EdgeBisecting:
		return "bisecting"
	case EdgeConfirmed:
		return "confirmed"
	default:
		return "invalid"
	}
}

// Defines structs that characterize actions an edge tracker
// can take to transition between states in its finite state machine.
type EdgeTrackerAction interface {
	fmt.Stringer
	IsEdgeTrackerAction() bool // Sentinel method that marks the interface as an edge tracker action.
}

// Transitions the edge tracker back to a start state.
type EdgeBackToStart struct{}

// Tracker will act if the edge is at a one step proof.
type EdgeHandleOneStepProof struct{}

// Tracker will add a subchallenge on its edge's subchallenge.
type EdgeOpenSubchallengeLeaf struct{}

// Tracker will attempt to bisect its edge.
type EdgeBisect struct{}

type EdgeConfirm struct{}

func (EdgeBackToStart) String() string {
	return "back_to_start"
}
func (EdgeHandleOneStepProof) String() string {
	return "check_one_step_proof"
}
func (EdgeOpenSubchallengeLeaf) String() string {
	return "open_subchallenge_leaf"
}
func (EdgeBisect) String() string {
	return "bisect"
}
func (EdgeConfirm) String() string {
	return "confirm"
}

func (EdgeBackToStart) IsEdgeTrackerAction() bool {
	return true
}
func (EdgeHandleOneStepProof) IsEdgeTrackerAction() bool {
	return true
}
func (EdgeOpenSubchallengeLeaf) IsEdgeTrackerAction() bool {
	return true
}
func (EdgeBisect) IsEdgeTrackerAction() bool {
	return true
}
func (EdgeConfirm) IsEdgeTrackerAction() bool {
	return true
}
