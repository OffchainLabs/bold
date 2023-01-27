package protocol

import (
	"fmt"

	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrWrongChallengeKind        = errors.New("wrong top-level kind for subchallenge creation")
	ErrNoChallenge               = errors.New("no challenge corresponds to vertex")
	ErrChallengeNotRunning       = errors.New("challenge is not ongoing")
	ErrSubchallengeAlreadyExists = errors.New("subchallenge already exists on vertex")
	ErrNotEnoughValidChildren    = errors.New("vertex needs at least two unexpired children")
)

// CreateSubChallenge validates a subchallenge can be created and creates it.
func (v *ChallengeVertex) CreateSubChallenge(
	tx *ActiveTx,
	validator common.Address,
	subchallengeType ChallengeType,
) (*Challenge, error) {
	tx.verifyReadWrite()
	if err := v.canCreateSubChallenge(subchallengeType); err != nil {
		return nil, err
	}
	return v.createSubChallenge(subchallengeType, validator)
}

// AddSubchallengeLeaf adds a leaf vertex to a subchallenge according to
// the rules of the specification.
// A leaf can be created with reference to a vertex V in a
// higher level challenge if:

// - V.predecessor = P, and
// - P.challenge is this challenge type, and
// - V.height = V.predecessor.height + 1, and

// The leaf creation will provide a merkle commitment to a
// sequence of wavm states, where the last state in the sequence must
// be proven to be the end state of the wavm VM corresponding to the state
// transition function completing and producing the last block in V’s merkle commitment.
func (c *Challenge) AddSubchallengeLeaf(
	tx *ActiveTx,
	topLevelVertex *ChallengeVertex,
	subchallengeType ChallengeType,
	history util.HistoryCommitment,
	validator common.Address,
) (*ChallengeVertex, error) {
	tx.verifyReadWrite()

	// We deduct a stake from the validator for creating a leaf vertex.
	if err := c.rootAssertion.Unwrap().chain.DeductFromBalance(tx, validator, ChallengeVertexStake); err != nil {
		return nil, errors.Wrapf(ErrInsufficientBalance, err.Error())
	}

	// The challenge and vertex provided must meet basic integrity checks.
	if c.rootVertex.IsNone() || topLevelVertex.Prev.IsNone() {
		return nil, ErrNoChallenge
	}
	// We must be creating a leaf for the appropriate subchallenge type.
	if subchallengeType == BlockChallenge {
		return nil, errors.Wrap(ErrWrongChallengeKind, "cannot add a subchallenge leaf to a block challenge")
	}
	if c.ChallengeType != subchallengeType {
		return nil, ErrWrongChallengeKind
	}

	// The vertex must be one-step away from its previous vertex.
	prev := topLevelVertex.Prev.Unwrap()
	if topLevelVertex.Commitment.Height != prev.Commitment.Height+1 {
		return nil, ErrInvalidHeight
	}

	// We verify other common invariants of challenge leaf addition.
	if err := c.canAddLeaf(
		tx,
		history,
		topLevelVertex.Commitment.LastLeaf,
		util.Some(topLevelVertex),
		validator,
	); err != nil {
		return nil, err
	}

	timer := topLevelVertex.PsTimer.Clone()
	return c.addLeafToChallenge(validator, history, timer, topLevelVertex.winnerIfConfirmed), nil
}

func (c *Challenge) canAddLeaf(
	tx *ActiveTx,
	history util.HistoryCommitment,
	lastLeafToCheck common.Hash,
	topLevelVertex util.Option[*ChallengeVertex],
	validator common.Address,
) error {
	if c.Completed(tx) {
		return ErrWrongState
	}
	if !c.rootVertex.Unwrap().EligibleForNewSuccessor() {
		return ErrPastDeadline
	}
	// We check if we have already included the vertex in the challenge.
	if c.includedHistories[history.Hash()] {
		return errors.Wrapf(ErrVertexAlreadyExists, fmt.Sprintf("Hash: %s", history.Hash().String()))
	}

	// If we are in a small step challenge, the height must be equal to 2^20 if the top-level
	// vertex is not a leaf or a positive value <= 2^20 if the top-level vertex is a leaf.
	if c.ChallengeType == SmallStepChallenge {
		if topLevelVertex.IsNone() {
			return errors.New("top level vertex cannot be empty in SmallStepChallenge")
		}
		topLevelV := topLevelVertex.Unwrap()
		if topLevelV.isLeaf {
			if history.Height > 1<<20 {
				return ErrInvalidHeight
			}
		} else {
			if history.Height != 1<<20 {
				return ErrInvalidHeight
			}
		}
	}

	// The last leaf claimed in the history commitment must be the
	// state root of the assertion we are adding a leaf for.
	if !historyProvidesLastLeafProof(history) {
		return ErrNoLastLeafProof
	}
	if lastLeafToCheck != history.LastLeaf {
		return errors.Wrapf(
			ErrWrongLastLeaf,
			"last leaf of history, %#x, does not match expected %#x",
			history.LastLeaf,
			lastLeafToCheck,
		)
	}

	// The validator must provide a history commitment over
	// a series of states where the last state must be proven to be
	// one corresponding to a correct last leaf.
	if err := util.VerifyPrefixProof(
		history.LastLeafPrefix.Unwrap(),
		history.Normalized().Unwrap(),
		history.LastLeafProof,
	); err != nil {
		return ErrProofFailsToVerify
	}
	return nil
}

// Adds a leaf vertex to a challenge by mutating it and returns
// the added vertex.
func (c *Challenge) addLeafToChallenge(
	validator common.Address,
	history util.HistoryCommitment,
	timer *util.CountUpTimer,
	winnerIfConfirmed util.Option[*Assertion],
) *ChallengeVertex {
	nextSeqNumber := c.currentVertexSeqNumber + 1
	leaf := &ChallengeVertex{
		Challenge:            util.Some(c),
		SequenceNum:          nextSeqNumber,
		Validator:            validator,
		isLeaf:               true,
		Status:               PendingAssertionState,
		Commitment:           history,
		Prev:                 c.rootVertex,
		PresumptiveSuccessor: util.None[*ChallengeVertex](),
		PsTimer:              timer,
		SubChallenge:         util.None[*Challenge](),
		winnerIfConfirmed:    winnerIfConfirmed,
	}
	c.currentVertexSeqNumber = nextSeqNumber

	tentativeWinner := leaf.winnerIfConfirmed.Unwrap()
	c.rootVertex.Unwrap().maybeNewPresumptiveSuccessor(leaf)
	c.rootAssertion.Unwrap().chain.challengesFeed.Append(&ChallengeLeafEvent{
		ParentSeqNum:      leaf.Prev.Unwrap().SequenceNum,
		SequenceNum:       leaf.SequenceNum,
		WinnerIfConfirmed: tentativeWinner.SequenceNum,
		History:           history,
		BecomesPS:         leaf.Prev.Unwrap().PresumptiveSuccessor.Unwrap() == leaf,
		Validator:         validator,
	})
	c.includedHistories[history.Hash()] = true
	h := c.Hash()
	c.rootAssertion.Unwrap().chain.challengesByHash[h] = c
	c.rootAssertion.Unwrap().chain.challengeVerticesByCommitHash[h][VertexCommitHash(leaf.Commitment.Hash())] = leaf
	c.leafVertexCount++
	return leaf
}

// Verifies the a subchallenge can be created on a challenge vertex
// based on specification validity conditions below:
//
//	A subchallenge can be created at a vertex P in a “parent” BlockChallenge if:
//	  - P’s challenge has not reached its end time
//	  - P’s has at least two children with unexpired chess clocks
//	The end time of the new challenge is set equal to the end time of P’s challenge.
func (v *ChallengeVertex) canCreateSubChallenge(
	subChallengeType ChallengeType,
) error {
	if v.Challenge.IsNone() {
		return ErrNoChallenge
	}
	chal := v.Challenge.Unwrap()
	// Can only create a subchallenge if the vertex is
	// part of a challenge of a specified kind.
	switch subChallengeType {
	case NoChallengeType:
		return ErrWrongChallengeKind
	case BlockChallenge:
		return ErrWrongChallengeKind
	case BigStepChallenge:
		if chal.ChallengeType != BlockChallenge {
			return ErrWrongChallengeKind
		}
	case SmallStepChallenge:
		if chal.ChallengeType != BigStepChallenge {
			return ErrWrongChallengeKind
		}
	}
	// The challenge must be ongoing.
	chain := chal.rootAssertion.Unwrap().chain
	if chal.hasEnded(chain) {
		return ErrChallengeNotRunning
	}
	// There must not exist a subchallenge.
	if !v.SubChallenge.IsNone() {
		return ErrSubchallengeAlreadyExists
	}
	// The vertex must not be confirmed.
	if v.Status == ConfirmedAssertionState {
		return errors.Wrap(ErrWrongState, "vertex already confirmed")
	}
	// The vertex must have at least two children with unexpired
	// chess clocks in order to create a big step challenge.
	ok, err := hasUnexpiredChildren(chain, v)
	if err != nil {
		return err
	}
	if !ok {
		return ErrNotEnoughValidChildren
	}
	return nil
}

func (v *ChallengeVertex) createSubChallenge(challengeType ChallengeType, validator common.Address) (*Challenge, error) {
	if v.Challenge.IsNone() {
		return nil, ErrNoChallenge
	}
	chal := v.Challenge.Unwrap()
	rootAssertion := chal.rootAssertion.Unwrap()
	chain := rootAssertion.chain
	subChal := createChallengeBase(
		chain,
		rootAssertion,
		challengeType,
		v.Challenge.Unwrap().creationTime,
	)
	v.SubChallenge = util.Some(subChal)

	parentStaker := common.Address{}
	if !rootAssertion.Staker.IsNone() {
		parentStaker = rootAssertion.Staker.Unwrap()
	}

	// TODO: Fire a subchallenge event instead?
	chain.feed.Append(&StartChallengeEvent{
		ParentSeqNum:          rootAssertion.SequenceNum,
		ParentStateCommitment: rootAssertion.StateCommitment,
		ParentStaker:          parentStaker,
		Validator:             validator,
	})

	challengeID := subChal.Hash()
	chain.challengesByHash[challengeID] = subChal
	rootVertex := subChal.rootVertex.Unwrap()
	rootCommit := VertexCommitHash(rootVertex.Commitment.Hash())
	chain.challengeVerticesByCommitHash[challengeID] = map[VertexCommitHash]*ChallengeVertex{rootCommit: rootVertex}
	return subChal, nil
}

func createChallengeBase(
	chain *AssertionChain,
	rootAssertion *Assertion,
	challengeType ChallengeType,
	creationTime time.Time,
) *Challenge {
	currSeqNumber := VertexSequenceNumber(0)
	rootVertex := &ChallengeVertex{
		Challenge:   util.None[*Challenge](),
		SequenceNum: currSeqNumber,
		isLeaf:      false,
		Status:      ConfirmedAssertionState,
		Commitment: util.HistoryCommitment{
			Height: 0,
			Merkle: common.Hash{},
		},
		Prev:                 util.None[*ChallengeVertex](),
		PresumptiveSuccessor: util.None[*ChallengeVertex](),
		PsTimer:              util.NewCountUpTimer(chain.timeReference),
		SubChallenge:         util.None[*Challenge](),
	}
	chal := &Challenge{
		rootAssertion:     util.Some(rootAssertion),
		WinnerAssertion:   util.None[*Assertion](),
		WinnerVertex:      util.None[*ChallengeVertex](),
		rootVertex:        util.Some(rootVertex),
		includedHistories: make(map[common.Hash]bool),
		challengePeriod:   chain.challengePeriod,
		creationTime:      creationTime,
		ChallengeType:     challengeType,
	}
	rootVertex.Challenge = util.Some(chal)
	chal.includedHistories[rootVertex.Commitment.Hash()] = true
	return chal
}

// Checks if a challenge vertex has at least two children with
// unexpired chess-clocks. It does this by filtering out vertices from the chain
// that are the specified vertex's children and checking that at least two in this
// filtered list have unexpired chess clocks and are one-step away from the parent.
func hasUnexpiredChildren(chain *AssertionChain, v *ChallengeVertex) (bool, error) {
	if v.Challenge.IsNone() {
		return false, ErrNoChallenge
	}
	chal := v.Challenge.Unwrap()
	challengeHash := chal.Hash()
	vertices, ok := chain.challengeVerticesByCommitHash[challengeHash]
	if !ok {
		return false, fmt.Errorf("vertices not found for challenge with hash: %#x", challengeHash)
	}
	vertexCommitHash := v.Commitment.Hash()
	unexpiredChildrenTotal := 0
	for _, otherVertex := range vertices {
		if otherVertex.Prev.IsNone() {
			continue
		}
		prev := otherVertex.Prev.Unwrap()
		parentCommitHash := prev.Commitment.Hash()
		isOneStepAway := otherVertex.Commitment.Height == prev.Commitment.Height+1
		isChild := parentCommitHash == vertexCommitHash
		if isOneStepAway && isChild && !otherVertex.chessClockExpired(chain.challengePeriod) {
			unexpiredChildrenTotal++
			if unexpiredChildrenTotal > 1 {
				return true, nil
			}
		}
	}
	return false, nil
}

// Checks if a challenge is still ongoing by making sure the current
// timestamp is within the challenge's creation time + challenge period.
func (c *Challenge) hasEnded(chain *AssertionChain) bool {
	challengeEndTime := c.creationTime.Add(chain.challengePeriod).Unix()
	now := chain.timeReference.Get().Unix()
	return now > challengeEndTime
}

// Checks if a vertex's chess-clock has expired according
// to the challenge period length.
func (v *ChallengeVertex) chessClockExpired(challengePeriod time.Duration) bool {
	return v.PsTimer.Get() > challengePeriod
}
