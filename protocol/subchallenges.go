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
	// The challenge and vertex provided must meet basic integrity checks.
	if c.rootVertex.IsNone() || topLevelVertex.Prev.IsNone() {
		return nil, ErrNoChallenge
	}

	// We must be creating a leaf for the appropriate subchallenge type.
	if subchallengeType == BlockChallenge {
		return nil, ErrWrongChallengeKind
	}
	if c.ChallengeType != subchallengeType {
		return nil, ErrWrongChallengeKind
	}
	challengedVertex := c.rootVertex.Unwrap()
	prev := topLevelVertex.Prev.Unwrap()

	// The previous vertex's challenge must be this one.
	if prev.Challenge.Unwrap() != c {
		return nil, ErrInvalidOp
	}

	// The previous vertex must match the root vertex of this challenge.
	if prev != challengedVertex {
		return nil, ErrInvalidOp
	}

	// The vertex must be one-step away from its previous vertex.
	// TODO: Should we check if the previous vertex is at a one-step-fork?
	if prev.Commitment.Height != topLevelVertex.Commitment.Height+1 {
		return nil, ErrInvalidOp
	}

	// We check if we have already included the vertex in the challenge.
	if c.includedHistories[history.Hash()] {
		return nil, errors.Wrapf(ErrVertexAlreadyExists, fmt.Sprintf("Hash: %s", history.Hash().String()))
	}

	// We deduct a stake from the validator for creating a leaf vertex.
	if err := c.rootAssertion.Unwrap().chain.DeductFromBalance(tx, validator, ChallengeVertexStake); err != nil {
		return nil, errors.Wrapf(ErrInsufficientBalance, err.Error())
	}

	// If we are in a small step challenge, the height must be equal to 2^20 if the top-level
	// vertex is not a leaf or a positive value <= 2^20 if the top-level vertex is a leaf.
	if c.ChallengeType == SmallStepChallenge {
		if topLevelVertex.isLeaf {
			if history.Height > 1<<20 {
				return nil, ErrInvalidHeight
			}
		} else {
			if history.Height != 1<<20 {
				return nil, ErrInvalidHeight
			}
		}
	}

	// The last leaf claimed in the history commitment must be the
	// state root of the assertion we are adding a leaf for.
	if history.LastLeaf == (common.Hash{}) ||
		len(history.LastLeafProof) == 0 ||
		history.LastLeafPrefix.IsNone() ||
		history.Normalized().IsNone() {
		return nil, errors.New("history commitment must provide a last leaf proof")
	}
	if topLevelVertex.Commitment.LastLeaf != history.LastLeaf {
		return nil, errors.Wrapf(
			ErrInvalidOp,
			"last leaf of history does not match top-level vertex's last leaf %#x != %#x",
			topLevelVertex.Commitment.LastLeaf,
			history.LastLeaf,
		)
	}

	// The validator must provide a history commitment over
	// a series of states where the last state must be proven to be
	// one corresponding to the top-level vertex.
	if err := util.VerifyPrefixProof(
		history.LastLeafPrefix.Unwrap(),
		history.Normalized().Unwrap(),
		history.LastLeafProof,
	); err != nil {
		return nil, errors.New(
			"merkle proof fails to verify for last state of history commitment",
		)
	}

	timer := topLevelVertex.PsTimer.Clone()
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
		winnerIfConfirmed:    topLevelVertex.winnerIfConfirmed,
	}
	c.currentVertexSeqNumber = nextSeqNumber

	winnerIfConfirmed := leaf.winnerIfConfirmed.Unwrap()
	c.rootVertex.Unwrap().maybeNewPresumptiveSuccessor(leaf)
	c.rootAssertion.Unwrap().chain.challengesFeed.Append(&ChallengeLeafEvent{
		ParentSeqNum:      leaf.Prev.Unwrap().SequenceNum,
		SequenceNum:       leaf.SequenceNum,
		WinnerIfConfirmed: winnerIfConfirmed.SequenceNum,
		History:           history,
		BecomesPS:         leaf.Prev.Unwrap().PresumptiveSuccessor.Unwrap() == leaf,
		Validator:         validator,
	})
	c.includedHistories[history.Hash()] = true
	h := c.Hash()
	c.rootAssertion.Unwrap().chain.challengesByHash[h] = c
	c.rootAssertion.Unwrap().chain.challengeVerticesByCommitHash[h][VertexCommitHash(leaf.Commitment.Hash())] = leaf
	c.leafVertexCount++
	return nil, nil
}

// CreateBigStepChallenge creates a BigStep subchallenge on a vertex.
func (v *ChallengeVertex) CreateBigStepChallenge(
	tx *ActiveTx,
	validator common.Address,
) (*Challenge, error) {
	tx.verifyReadWrite()
	if err := v.canCreateSubChallenge(BigStepChallenge); err != nil {
		return nil, err
	}
	// TODO: Add the challenge to the chain under a key that does not
	// collide with top-level challenges and fire events.
	rootAssertion := v.Challenge.Unwrap().rootAssertion.Unwrap()
	chain := rootAssertion.chain
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

	subChal := &Challenge{
		rootAssertion:     util.Some(rootAssertion),
		WinnerAssertion:   util.None[*Assertion](),
		WinnerVertex:      util.None[*ChallengeVertex](),
		rootVertex:        util.Some(rootVertex),
		includedHistories: make(map[common.Hash]bool),
		challengePeriod:   chain.challengePeriod,
		// Set the creation time of the subchallenge to be
		// the same as the top-level challenge, as they should
		// expire at the same timestamp.
		creationTime:  v.Challenge.Unwrap().creationTime,
		ChallengeType: BigStepChallenge,
	}

	rootVertex.Challenge = util.Some(subChal)
	subChal.includedHistories[rootVertex.Commitment.Hash()] = true
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
	if _, ok := chain.challengesByHash[challengeID]; ok {
		return nil, errors.New("challenge with id already exists")
	}
	chain.challengesByHash[challengeID] = subChal
	chain.challengeVerticesByCommitHash[challengeID] = map[VertexCommitHash]*ChallengeVertex{VertexCommitHash(rootVertex.Commitment.Hash()): rootVertex}

	return subChal, nil
}

// CreateSmallStepChallenge creates a SmallStep subchallenge on a vertex.
func (v *ChallengeVertex) CreateSmallStepChallenge(tx *ActiveTx) error {
	tx.verifyReadWrite()
	if err := v.canCreateSubChallenge(SmallStepChallenge); err != nil {
		return err
	}
	// TODO: Add all other required challenge fields.
	v.SubChallenge = util.Some(&Challenge{
		creationTime:  v.Challenge.Unwrap().creationTime,
		ChallengeType: SmallStepChallenge,
	})
	// TODO: Add the challenge to the chain under a key that does not
	// collide with top-level challenges and fire events.
	return nil
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
