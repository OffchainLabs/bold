package solimpl

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/rollupgen"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/nitro/util/headerreader"
	"github.com/pkg/errors"
)

var rollupAssertionCreatedId common.Hash

func init() {
	abi, err := rollupgen.RollupCoreMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := abi.Events["AssertionCreated"]
	if !ok {
		panic("Didn't find AssertionCreated event")
	}
	rollupAssertionCreatedId = event.ID
}

func (e *SpecEdge) Id() protocol.EdgeId {
	return e.id
}

func (e *SpecEdge) GetType() protocol.EdgeType {
	return protocol.EdgeType(e.inner.EType)
}

func (e *SpecEdge) MiniStaker() util.Option[common.Address] {
	return e.miniStaker
}

func (e *SpecEdge) StartCommitment() (protocol.Height, common.Hash) {
	return protocol.Height(e.inner.StartHeight.Uint64()), e.inner.StartHistoryRoot
}

func (e *SpecEdge) EndCommitment() (protocol.Height, common.Hash) {
	return protocol.Height(e.inner.EndHeight.Uint64()), e.inner.EndHistoryRoot
}

func (e *SpecEdge) TimeUnrivaled(ctx context.Context) (uint64, error) {
	timer, err := e.manager.caller.TimeUnrivaled(&bind.CallOpts{Context: ctx}, e.id)
	if err != nil {
		return 0, err
	}
	return timer.Uint64(), nil
}

func (e *SpecEdge) HasRival(ctx context.Context) (bool, error) {
	return e.manager.caller.HasRival(&bind.CallOpts{Context: ctx}, e.id)
}

func (e *SpecEdge) Status(ctx context.Context) (protocol.EdgeStatus, error) {
	edge, err := e.manager.caller.GetEdge(&bind.CallOpts{Context: ctx}, e.id)
	if err != nil {
		return 0, err
	}
	return protocol.EdgeStatus(edge.Status), nil
}

func (e *SpecEdge) HasLengthOneRival(ctx context.Context) (bool, error) {
	ok, err := e.manager.caller.HasLengthOneRival(&bind.CallOpts{Context: ctx}, e.id)
	if err != nil {
		errS := err.Error()
		switch {
		case strings.Contains(errS, "not length 1"):
			return false, nil
		case strings.Contains(errS, "is unrivaled"):
			return false, nil
		default:
			return false, err
		}
	}
	return ok, nil
}

func (e *SpecEdge) Bisect(
	ctx context.Context,
	prefixHistoryRoot common.Hash,
	prefixProof []byte,
) (protocol.SpecEdge, protocol.SpecEdge, error) {
	_, err := transact(ctx, e.manager.backend, e.manager.reader, func() (*types.Transaction, error) {
		return e.manager.writer.BisectEdge(e.manager.txOpts, e.id, prefixHistoryRoot, prefixProof)
	})
	if err != nil {
		return nil, nil, err
	}
	someEdge, err := e.manager.GetEdge(ctx, e.id)
	if err != nil {
		return nil, nil, err
	}
	if someEdge.IsNone() {
		return nil, nil, errors.New("could not refresh edge after bisecting, got empty result")
	}
	edge, ok := someEdge.Unwrap().(*SpecEdge)
	if !ok {
		return nil, nil, errors.New("not a *SpecEdge")
	}
	// Refresh the edge.
	e = edge
	someLowerChild, err := e.manager.GetEdge(ctx, e.inner.LowerChildId)
	if err != nil {
		return nil, nil, err
	}
	someUpperChild, err := e.manager.GetEdge(ctx, e.inner.UpperChildId)
	if err != nil {
		return nil, nil, err
	}
	if someLowerChild.IsNone() || someUpperChild.IsNone() {
		return nil, nil, errors.New("expected edge to have children post-bisection, but has none")
	}
	return someLowerChild.Unwrap(), someUpperChild.Unwrap(), nil
}

func (e *SpecEdge) ConfirmByTimer(ctx context.Context, ancestorIds []protocol.EdgeId) error {
	ancestors := make([][32]byte, len(ancestorIds))
	for i, r := range ancestorIds {
		ancestors[i] = r
	}
	_, err := transact(ctx, e.manager.backend, e.manager.reader, func() (*types.Transaction, error) {
		return e.manager.writer.ConfirmEdgeByTime(e.manager.txOpts, e.id, ancestors)
	})
	return err
}

func (e *SpecEdge) ConfirmByChildren(ctx context.Context) error {
	_, err := transact(ctx, e.manager.backend, e.manager.reader, func() (*types.Transaction, error) {
		return e.manager.writer.ConfirmEdgeByChildren(e.manager.txOpts, e.id)
	})
	return err
}

func (e *SpecEdge) ConfirmByClaim(ctx context.Context, claimId protocol.ClaimId) error {
	// TODO: Add in fields.
	_, err := transact(ctx, e.manager.backend, e.manager.reader, func() (*types.Transaction, error) {
		return e.manager.writer.ConfirmEdgeByClaim(e.manager.txOpts, e.id, claimId)
	})
	return err
}

func (e *SpecEdge) ConfirmByOneStepProof(ctx context.Context) error {
	_, err := transact(ctx, e.manager.backend, e.manager.reader, func() (*types.Transaction, error) {
		return e.manager.writer.ConfirmEdgeByOneStepProof(
			e.manager.txOpts,
			e.id,
			// TODO: Fill in.
			challengeV2gen.OneStepData{},
			// TODO: Add pre/post proofs.
			nil,
			nil,
		)
	})
	return err
}

func (e *SpecEdge) MutualId(ctx context.Context) (protocol.MutualId, error) {
	return e.manager.caller.CalculateMutualId(
		&bind.CallOpts{Context: ctx},
		e.inner.EType,
		e.inner.OriginId,
		e.inner.StartHeight,
		e.inner.StartHistoryRoot,
		e.inner.EndHeight,
	)
}

// TopLevelClaimHeight gets the height at the BlockChallenge level that originated a subchallenge.
// For example, if two validators open a subchallenge S at edge A in a BlockChallenge, the TopLevelClaimHeight of S is the height of A.
// If two validators open a subchallenge S' at edge B in BigStepChallenge, the TopLevelClaimHeight
// is the height of A.
func (e *SpecEdge) TopLevelClaimHeight(ctx context.Context) (*protocol.OriginHeights, error) {
	switch e.GetType() {
	case protocol.BigStepChallengeEdge:
		rivalId, err := e.manager.caller.FirstRival(&bind.CallOpts{Context: ctx}, e.inner.OriginId)
		if err != nil {
			return nil, err
		}
		blockChallengeOneStepForkSource, err := e.manager.GetEdge(ctx, rivalId)
		if err != nil {
			return nil, errors.Wrapf(err, "block challenge one step fork source does not exist for rival id %#x", rivalId)
		}
		if blockChallengeOneStepForkSource.IsNone() {
			return nil, errors.New("source edge is none")
		}
		startHeight, _ := blockChallengeOneStepForkSource.Unwrap().StartCommitment()
		return &protocol.OriginHeights{
			BlockChallengeOriginHeight: startHeight,
		}, nil
	case protocol.SmallStepChallengeEdge:
		rivalId, err := e.manager.caller.FirstRival(&bind.CallOpts{Context: ctx}, e.inner.OriginId)
		if err != nil {
			return nil, err
		}
		bigStepChallengeOneStepForkSource, err := e.manager.GetEdge(ctx, rivalId)
		if err != nil {
			return nil, errors.Wrap(err, "big step challenge one step fork source does not exist")
		}
		if bigStepChallengeOneStepForkSource.IsNone() {
			return nil, errors.New("source edge is none")
		}
		bigStepEdge, ok := bigStepChallengeOneStepForkSource.Unwrap().(*SpecEdge)
		if !ok {
			return nil, errors.New("not *SpecEdge")
		}
		rivalId, err = e.manager.caller.FirstRival(&bind.CallOpts{Context: ctx}, bigStepEdge.inner.OriginId)
		if err != nil {
			return nil, err
		}
		blockChallengeOneStepForkSource, err := e.manager.GetEdge(ctx, rivalId)
		if err != nil {
			return nil, errors.Wrap(err, "block challenge one step fork source does not exist")
		}
		if blockChallengeOneStepForkSource.IsNone() {
			return nil, errors.New("source edge is none")
		}
		bigStepStartHeight, _ := bigStepEdge.StartCommitment()
		blockChallengeStartHeight, _ := blockChallengeOneStepForkSource.Unwrap().StartCommitment()
		return &protocol.OriginHeights{
			BlockChallengeOriginHeight:   blockChallengeStartHeight,
			BigStepChallengeOriginHeight: bigStepStartHeight,
		}, nil
	default:
		startHeight, _ := e.StartCommitment()
		return &protocol.OriginHeights{
			BlockChallengeOriginHeight: startHeight,
		}, nil
	}
}

func (e *SpecEdge) TopLevelClaimEndBatchCount(ctx context.Context) (uint64, error) {
	// TODO: presumably, this should be stored with the edge in Solidity. I think it'll be needed there anyways.
	// I think we could also in Go pass it down from parent to child but I don't want to refactor that right now.
	callOpts := &bind.CallOpts{Context: ctx}
	// Find the parent assertion of this challenge
	parentId := e.inner.OriginId
	for {
		exists, err := e.manager.caller.EdgeExists(callOpts, parentId)
		if err != nil {
			return 0, err
		}
		if !exists {
			rival, err := e.manager.caller.FirstRival(callOpts, parentId)
			if err != nil {
				return 0, err
			}
			if rival != (common.Hash{}) {
				parentId = rival
				exists = true
			}
		}
		if !exists {
			break
		}
		edge, err := e.manager.caller.GetEdge(callOpts, parentId)
		if err != nil {
			return 0, err
		}
		parentId = edge.OriginId
	}
	rollup := e.manager.assertionChain.rollup
	assertionNum, err := rollup.GetAssertionNum(callOpts, parentId)
	if err != nil {
		return 0, err
	}
	assertion, err := rollup.GetAssertion(callOpts, assertionNum)
	if err != nil {
		return 0, err
	}
	if !assertion.InboxMsgCountSeen.IsUint64() {
		return 0, fmt.Errorf("assertion %v inbox max count seen %v is not a uint64", assertionNum, assertion.InboxMsgCountSeen)
	}
	return assertion.InboxMsgCountSeen.Uint64(), nil
}

// SpecChallengeManager is a wrapper around the challenge manager contract.
type SpecChallengeManager struct {
	addr           common.Address
	backend        ChainBackend
	assertionChain *AssertionChain
	reader         *headerreader.HeaderReader
	txOpts         *bind.TransactOpts
	caller         *challengeV2gen.EdgeChallengeManagerCaller
	writer         *challengeV2gen.EdgeChallengeManagerTransactor
	filterer       *challengeV2gen.EdgeChallengeManagerFilterer
}

// NewSpecChallengeManager returns an instance of the spec challenge manager
// used by the assertion chain.
func NewSpecChallengeManager(
	ctx context.Context,
	addr common.Address,
	assertionChain *AssertionChain,
	backend ChainBackend,
	reader *headerreader.HeaderReader,
	txOpts *bind.TransactOpts,
) (protocol.SpecChallengeManager, error) {
	managerBinding, err := challengeV2gen.NewEdgeChallengeManager(addr, backend)
	if err != nil {
		return nil, err
	}
	return &SpecChallengeManager{
		addr:           addr,
		assertionChain: assertionChain,
		backend:        backend,
		reader:         reader,
		txOpts:         txOpts,
		caller:         &managerBinding.EdgeChallengeManagerCaller,
		writer:         &managerBinding.EdgeChallengeManagerTransactor,
		filterer:       &managerBinding.EdgeChallengeManagerFilterer,
	}, nil
}

func (cm *SpecChallengeManager) Address() common.Address {
	return cm.addr
}

// Duration of the challenge period.
func (cm *SpecChallengeManager) ChallengePeriodSeconds(
	ctx context.Context,
) (time.Duration, error) {
	res, err := cm.caller.ChallengePeriodSec(&bind.CallOpts{Context: ctx})
	if err != nil {
		return time.Second, err
	}
	return time.Second * time.Duration(res.Uint64()), nil
}

// Gets an edge by its hash.
func (cm *SpecChallengeManager) GetEdge(
	ctx context.Context,
	edgeId protocol.EdgeId,
) (util.Option[protocol.SpecEdge], error) {
	edge, err := cm.caller.GetEdge(&bind.CallOpts{Context: ctx}, edgeId)
	if err != nil {
		return util.None[protocol.SpecEdge](), err
	}
	miniStaker := util.None[common.Address]()
	if edge.Staker != (common.Address{}) {
		miniStaker = util.Some(edge.Staker)
	}
	return util.Some(protocol.SpecEdge(&SpecEdge{
		id:         edgeId,
		manager:    cm,
		inner:      edge,
		miniStaker: miniStaker,
	})), nil
}

// Calculates an edge hash given its challenge id, start history, and end history.
func (cm *SpecChallengeManager) CalculateEdgeId(
	ctx context.Context,
	edgeType protocol.EdgeType,
	originId protocol.OriginId,
	startHeight protocol.Height,
	startHistoryRoot common.Hash,
	endHeight protocol.Height,
	endHistoryRoot common.Hash,
) (protocol.EdgeId, error) {
	return cm.caller.CalculateEdgeId(
		&bind.CallOpts{Context: ctx},
		uint8(edgeType),
		originId,
		big.NewInt(int64(startHeight)),
		startHistoryRoot,
		big.NewInt(int64(endHeight)),
		endHistoryRoot,
	)
}

// ConfirmEdgeByOneStepProof checks a one step proof for a tentative winner vertex id
// which will mark it as the winning claim of its associated challenge if correct.
// The edges along the winning branch and the corresponding assertion then need to be confirmed
// through separate transactions, if this succeeds.
func (cm *SpecChallengeManager) ConfirmEdgeByOneStepProof(
	ctx context.Context,
	tentativeWinnerId protocol.EdgeId,
	oneStepData *protocol.OneStepData,
	preHistoryInclusionProof []common.Hash,
	postHistoryInclusionProof []common.Hash,
) error {
	pre := make([][32]byte, len(preHistoryInclusionProof))
	for i, r := range preHistoryInclusionProof {
		pre[i] = r
	}
	post := make([][32]byte, len(postHistoryInclusionProof))
	for i, r := range postHistoryInclusionProof {
		post[i] = r
	}
	_, err := transact(
		ctx,
		cm.assertionChain.backend,
		cm.assertionChain.headerReader,
		func() (*types.Transaction, error) {
			return cm.writer.ConfirmEdgeByOneStepProof(
				cm.assertionChain.txOpts,
				tentativeWinnerId,
				challengeV2gen.OneStepData{
					ExecCtx: challengeV2gen.ExecutionContext{
						MaxInboxMessagesRead: big.NewInt(int64(oneStepData.MaxInboxMessagesRead)),
						Bridge:               oneStepData.BridgeAddr,
					},
					MachineStep: big.NewInt(int64(oneStepData.MachineStep)),
					BeforeHash:  oneStepData.BeforeHash,
					Proof:       oneStepData.Proof,
				},
				pre,
				post,
			)
		})
	// TODO: Handle receipt.
	return err
}

// Like abi.NewType but panics if it fails for use in constants
func newStaticType(t string, internalType string, components []abi.ArgumentMarshaling) abi.Type {
	ty, err := abi.NewType(t, internalType, components)
	if err != nil {
		panic(err)
	}
	return ty
}

var bytes32Type = newStaticType("bytes32", "", nil)
var bytes32ArrayType = newStaticType("bytes32[]", "", []abi.ArgumentMarshaling{{Type: "bytes32"}})

var blockEdgeProofAbi = abi.Arguments{{
	Name: "inclusionProof",
	Type: bytes32ArrayType,
}}

func (cm *SpecChallengeManager) AddBlockChallengeLevelZeroEdge(
	ctx context.Context,
	assertion protocol.Assertion,
	startCommit util.HistoryCommitment,
	endCommit util.HistoryCommitment,
	startEndPrefixProof []byte,
) (protocol.SpecEdge, error) {
	assertionId, err := cm.assertionChain.GetAssertionId(ctx, assertion.SeqNum())
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"could not get id for assertion with sequence num %d",
			assertion.SeqNum(),
		)
	}
	prevSeqNum, err := assertion.PrevSeqNum()
	if err != nil {
		return nil, err
	}
	prevAssertionId, err := cm.assertionChain.GetAssertionId(ctx, prevSeqNum)
	if err != nil {
		return nil, err
	}
	if startCommit.Height != 0 {
		return nil, fmt.Errorf("start commit has unexpected height %v (expected 0)", startCommit.Height)
	}
	if endCommit.Height != protocol.LayerZeroBlockEdgeHeight {
		return nil, fmt.Errorf("end commit has unexpected height %v (expected %v)", startCommit.Height, protocol.LayerZeroBlockEdgeHeight)
	}
	if startCommit.FirstLeaf != endCommit.FirstLeaf {
		return nil, fmt.Errorf("start commit first leaf %v didn't match end commit first leaf %v", startCommit.FirstLeaf, endCommit.FirstLeaf)
	}
	blockEdgeProof, err := blockEdgeProofAbi.Pack(endCommit.LastLeafProof)
	if err != nil {
		return nil, err
	}
	if true { // Lee TODO: remove (this was just for debugging)
		h, err := cm.assertionChain.rollup.GetStateHash(&bind.CallOpts{}, prevAssertionId)
		if err != nil {
			return nil, err
		}
		fmt.Printf("before hash: %v vs %v\n", common.Hash(h), startCommit.FirstLeaf)
		h, err = cm.assertionChain.rollup.GetStateHash(&bind.CallOpts{}, assertionId)
		if err != nil {
			return nil, err
		}
		fmt.Printf("end hash: %v vs %v\n", common.Hash(h), endCommit.LastLeaf)
	}
	_, err = transact(ctx, cm.backend, cm.reader, func() (*types.Transaction, error) {
		return cm.writer.CreateLayerZeroEdge(
			cm.txOpts,
			challengeV2gen.CreateEdgeArgs{
				EdgeType:         uint8(protocol.BlockChallengeEdge),
				StartHistoryRoot: startCommit.Merkle,
				StartHeight:      big.NewInt(int64(startCommit.Height)),
				EndHistoryRoot:   endCommit.Merkle,
				EndHeight:        big.NewInt(int64(endCommit.Height)),
				ClaimId:          assertionId,
			},
			startEndPrefixProof,
			blockEdgeProof,
		)
	})
	if err != nil {
		return nil, err
	}

	edgeId, err := cm.CalculateEdgeId(
		ctx,
		protocol.BlockChallengeEdge,
		protocol.OriginId(prevAssertionId),
		protocol.Height(startCommit.Height),
		startCommit.Merkle,
		protocol.Height(endCommit.Height),
		endCommit.Merkle,
	)
	if err != nil {
		return nil, err
	}
	someLevelZeroEdge, err := cm.GetEdge(ctx, edgeId)
	if err != nil {
		return nil, err
	}
	if someLevelZeroEdge.IsNone() {
		return nil, errors.New("got empty, newly created level zero edge")
	}
	return someLevelZeroEdge.Unwrap(), nil
}

var subchallengeEdgeProofAbi = abi.Arguments{
	{
		Name: "startState",
		Type: bytes32Type,
	},
	{
		Name: "endState",
		Type: bytes32Type,
	},
	{
		Name: "claimStartInclusionProof",
		Type: bytes32ArrayType,
	},
	{
		Name: "claimEndInclusionProof",
		Type: bytes32ArrayType,
	},
	{
		Name: "edgeInclusionProof",
		Type: bytes32ArrayType,
	},
}

func (cm *SpecChallengeManager) AddSubChallengeLevelZeroEdge(
	ctx context.Context,
	challengedEdge protocol.SpecEdge,
	startCommit util.HistoryCommitment,
	endCommit util.HistoryCommitment,
	startParentInclusionProof []common.Hash,
	endParentInclusionProof []common.Hash,
	startEndPrefixProof []byte,
) (protocol.SpecEdge, error) {
	var subChalTyp protocol.EdgeType
	switch challengedEdge.GetType() {
	case protocol.BlockChallengeEdge:
		subChalTyp = protocol.BigStepChallengeEdge
	case protocol.BigStepChallengeEdge:
		subChalTyp = protocol.SmallStepChallengeEdge
	default:
		return nil, fmt.Errorf("cannot open level zero edge beneath small step challenge: %s", challengedEdge.GetType())
	}
	subchallengeEdgeProof, err := subchallengeEdgeProofAbi.Pack(startCommit.FirstLeaf, endCommit.LastLeaf, startParentInclusionProof, endParentInclusionProof, endCommit.LastLeafProof)
	if err != nil {
		return nil, err
	}
	_, err = transact(ctx, cm.backend, cm.reader, func() (*types.Transaction, error) {
		return cm.writer.CreateLayerZeroEdge(
			cm.txOpts,
			challengeV2gen.CreateEdgeArgs{
				EdgeType:         uint8(subChalTyp),
				StartHistoryRoot: startCommit.Merkle,
				StartHeight:      big.NewInt(int64(startCommit.Height)),
				EndHistoryRoot:   endCommit.Merkle,
				EndHeight:        big.NewInt(int64(endCommit.Height)),
				ClaimId:          challengedEdge.Id(),
			},
			startEndPrefixProof,
			subchallengeEdgeProof,
		)
	})
	if err != nil {
		return nil, err
	}
	challenged, ok := challengedEdge.(*SpecEdge)
	if !ok {
		return nil, errors.New("not a *SpecEdge")
	}
	mutualId, err := challenged.MutualId(ctx)
	if err != nil {
		return nil, err
	}
	edgeId, err := cm.CalculateEdgeId(
		ctx,
		subChalTyp,
		protocol.OriginId(mutualId),
		protocol.Height(startCommit.Height),
		startCommit.Merkle,
		protocol.Height(endCommit.Height),
		endCommit.Merkle,
	)
	if err != nil {
		return nil, err
	}
	someLevelZeroEdge, err := cm.GetEdge(ctx, edgeId)
	if err != nil {
		return nil, err
	}
	if someLevelZeroEdge.IsNone() {
		return nil, errors.New("got empty, newly created level zero edge")
	}
	return someLevelZeroEdge.Unwrap(), nil
}
