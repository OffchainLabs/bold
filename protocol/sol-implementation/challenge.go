package solimpl

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Challenge) Id() protocol.ChallengeHash {
	return c.id
}

func (c *Challenge) Challenger(ctx context.Context, tx protocol.ActiveTx) (common.Address, error) {
	challenge, err := c.fetchChallenge(ctx, tx)
	if err != nil {
		return common.Address{}, err
	}
	return challenge.Challenger, nil
}

func (c *Challenge) fetchChallenge(ctx context.Context, tx protocol.ActiveTx) (*challengeV2gen.Challenge, error) {
	manager, err := c.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return nil, err
	}
	challenge, err := manager.GetChallenge(ctx, tx, c.id)
	if err != nil {
		return nil, err
	}
	if challenge.IsNone() {
		return nil, ErrNotFound
	}
	return challenge.Unwrap(), nil
}

func (c *Challenge) RootAssertion(
	ctx context.Context, tx protocol.ActiveTx,
) (protocol.Assertion, error) {
	manager, err := c.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return nil, err
	}
	challenge, err := c.fetchChallenge(ctx, tx)
	if err != nil {
		return nil, err
	}
	rootVertex, err := manager.GetVertex(ctx, tx, challenge.RootId)
	if err != nil {
		return nil, err
	}
	if rootVertex.IsNone() {
		return nil, errors.New("root vertex not found")
	}
	root := rootVertex.Unwrap()
	assertionNum, err := c.assertionChain.GetAssertionNum(ctx, tx, root.ClaimId)
	if err != nil {
		return nil, err
	}
	assertion, err := c.assertionChain.AssertionBySequenceNum(ctx, tx, assertionNum)
	if err != nil {
		return nil, err
	}
	return assertion, nil
}

func (c *Challenge) RootVertex(
	ctx context.Context, tx protocol.ActiveTx,
) (protocol.ChallengeVertex, error) {
	challenge, err := c.fetchChallenge(ctx, tx)
	if err != nil {
		return nil, err
	}
	rootId := challenge.RootId
	return &ChallengeVertex{
		id:             rootId,
		assertionChain: c.assertionChain,
	}, nil
}

func (c *Challenge) WinningClaim(ctx context.Context, tx protocol.ActiveTx) (util.Option[protocol.AssertionHash], error) {
	challenge, err := c.fetchChallenge(ctx, tx)
	if err != nil {
		return util.None[protocol.AssertionHash](), err
	}
	if challenge.WinningClaim == [32]byte{} {
		return util.None[protocol.AssertionHash](), ErrNotFound
	}
	return util.Some[protocol.AssertionHash](challenge.WinningClaim), nil
}

func (c *Challenge) GetType(ctx context.Context, tx protocol.ActiveTx) (util.Option[protocol.ChallengeType], error) {
	challenge, err := c.fetchChallenge(ctx, tx)
	if err != nil {
		return util.None[protocol.ChallengeType](), err
	}
	return util.Some[protocol.ChallengeType](protocol.ChallengeType(challenge.ChallengeType)), nil
}

func (c *Challenge) GetCreationTime(
	ctx context.Context, tx protocol.ActiveTx,
) (time.Time, error) {
	return time.Time{}, errors.New("unimplemented")
}

func (c *Challenge) ParentStateCommitment(
	ctx context.Context, tx protocol.ActiveTx,
) (util.StateCommitment, error) {
	challenge, err := c.fetchChallenge(ctx, tx)
	if err != nil {
		return util.StateCommitment{}, err
	}
	manager, err := c.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return util.StateCommitment{}, err
	}
	v, err := manager.GetVertex(ctx, tx, challenge.RootId)
	if err != nil {
		return util.StateCommitment{}, err
	}
	if v.IsNone() {
		return util.StateCommitment{}, errors.New("no root vertex for challenge")
	}
	concreteV := v.Unwrap()
	assertionSeqNum, err := c.assertionChain.rollup.GetAssertionNum(
		c.assertionChain.callOpts, concreteV.ClaimId,
	)
	if err != nil {
		return util.StateCommitment{}, err
	}
	assertion, err := c.assertionChain.AssertionBySequenceNum(ctx, tx, protocol.AssertionSequenceNumber(assertionSeqNum))
	if err != nil {
		return util.StateCommitment{}, err
	}
	height, err := assertion.Height()
	if err != nil {
		return util.StateCommitment{}, err
	}
	stateHash, err := assertion.StateHash()
	if err != nil {
		return util.StateCommitment{}, err
	}
	return util.StateCommitment{
		Height:    height,
		StateRoot: stateHash,
	}, nil
}

func (c *Challenge) WinnerVertex(
	ctx context.Context, tx protocol.ActiveTx,
) (util.Option[protocol.ChallengeVertex], error) {
	return util.None[protocol.ChallengeVertex](), errors.New("unimplemented")
}

func (c *Challenge) Completed(
	ctx context.Context, tx protocol.ActiveTx,
) (bool, error) {
	return false, errors.New("unimplemented")
}

// AddBlockChallengeLeaf vertex to a BlockChallenge using an assertion and a history commitment.
func (c *Challenge) AddBlockChallengeLeaf(
	ctx context.Context,
	tx protocol.ActiveTx,
	assertion protocol.Assertion,
	history util.HistoryCommitment,
) (protocol.ChallengeVertex, error) {
	// Flatten the last leaf proof for submission to the chain.
	lastLeafProof := make([]byte, 0)
	for _, h := range history.LastLeafProof {
		lastLeafProof = append(lastLeafProof, h[:]...)
	}
	callOpts := c.assertionChain.callOpts
	assertionId, err := c.assertionChain.rollup.GetAssertionId(callOpts, uint64(assertion.SeqNum()))
	if err != nil {
		return nil, err
	}
	prevSeqNum, err := assertion.PrevSeqNum()
	if err != nil {
		return nil, err
	}
	prevAssertion, err := c.assertionChain.AssertionBySequenceNum(ctx, tx, prevSeqNum)
	if err != nil {
		return nil, err
	}
	prevAssertionStateHash, err := prevAssertion.StateHash()
	if err != nil {
		return nil, err
	}
	leafData := challengeV2gen.AddLeafArgs{
		ChallengeId:            c.id,
		ClaimId:                assertionId,
		Height:                 big.NewInt(int64(history.Height)),
		HistoryRoot:            history.Merkle,
		FirstState:             prevAssertionStateHash,
		FirstStatehistoryProof: make([]byte, 0), // TODO: Add in.
		LastState:              history.LastLeaf,
		LastStatehistoryProof:  lastLeafProof,
	}

	manager, err := c.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return nil, err
	}
	caller, err := manager.GetCaller(ctx, tx)
	if err != nil {
		return nil, err
	}
	// Check the current mini-stake amount and transact using that as the value.
	miniStake, err := caller.MiniStakeValue(c.assertionChain.callOpts)
	if err != nil {
		return nil, err
	}
	opts := copyTxOpts(c.assertionChain.txOpts)
	opts.Value = miniStake

	writer, err := manager.GetWriter(ctx, tx)
	if err != nil {
		return nil, err
	}
	_, err = transact(ctx, c.assertionChain.backend, c.assertionChain.headerReader, func() (*types.Transaction, error) {
		return writer.AddLeaf(
			opts,
			leafData,
			lastLeafProof,
			make([]byte, 0), // Inbox proof
		)
	})
	if err != nil {
		return nil, err
	}

	vertexId, err := caller.CalculateChallengeVertexId(
		c.assertionChain.callOpts,
		c.id,
		history.Merkle,
		big.NewInt(int64(history.Height)),
	)
	if err != nil {
		return nil, err
	}
	return &ChallengeVertex{
		id:             vertexId,
		assertionChain: c.assertionChain,
	}, nil
}

// AddSubChallengeLeaf adds the appropriate leaf to the challenge based on a vertex and history commitment.
func (c *Challenge) AddSubChallengeLeaf(
	ctx context.Context,
	tx protocol.ActiveTx,
	vertex protocol.ChallengeVertex,
	history util.HistoryCommitment,
) (protocol.ChallengeVertex, error) {
	// Flatten the last leaf proof for submission to the chain.
	lastLeafProof := make([]byte, 0)
	for _, h := range history.LastLeafProof {
		lastLeafProof = append(lastLeafProof, h[:]...)
	}

	prev, err := vertex.Prev(ctx, tx)
	if err != nil {
		return nil, err
	}
	if prev.IsNone() {
		return nil, errors.New("no prev vertex")
	}
	manager, err := c.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return nil, err
	}
	caller, err := manager.GetCaller(ctx, tx)
	if err != nil {
		return nil, err
	}
	prevVertexId, err := caller.CalculateChallengeVertexId(
		c.assertionChain.callOpts,
		prev.Unwrap().ChallengeId,
		history.Merkle,
		big.NewInt(int64(history.Height)),
	)
	parentVertex, err := caller.GetVertex(
		c.assertionChain.callOpts,
		prevVertexId,
	)
	if err != nil {
		return nil, err
	}
	leafData := challengeV2gen.AddLeafArgs{
		ChallengeId:            c.id,
		ClaimId:                vertex.Id(),
		Height:                 big.NewInt(int64(history.Height)),
		HistoryRoot:            history.Merkle,
		FirstState:             parentVertex.HistoryRoot,
		FirstStatehistoryProof: make([]byte, 0), // TODO: Add in.
		LastState:              history.LastLeaf,
		LastStatehistoryProof:  lastLeafProof,
	}

	// Check the current mini-stake amount and transact using that as the value.
	miniStake, err := caller.MiniStakeValue(c.assertionChain.callOpts)
	if err != nil {
		return nil, err
	}
	opts := copyTxOpts(c.assertionChain.txOpts)
	opts.Value = miniStake

	writer, err := manager.GetWriter(ctx, tx)
	if err != nil {
		return nil, err
	}
	_, err = transact(ctx, c.assertionChain.backend, c.assertionChain.headerReader, func() (*types.Transaction, error) {
		return writer.AddLeaf(
			opts,
			leafData,
			lastLeafProof,
			lastLeafProof, // TODO(RJ): Should be different for big and small step.
		)
	})
	if err != nil {
		return nil, err
	}

	vertexId, err := caller.CalculateChallengeVertexId(
		c.assertionChain.callOpts,
		c.id,
		history.Merkle,
		big.NewInt(int64(history.Height)),
	)
	if err != nil {
		return nil, err
	}
	return &ChallengeVertex{
		id:             vertexId,
		assertionChain: c.assertionChain,
	}, nil
}
