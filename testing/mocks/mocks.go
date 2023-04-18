package mocks

import (
	"context"
	"math/big"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	statemanager "github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
)

var (
	_ = protocol.SpecChallengeManager(&MockSpecChallengeManager{})
	_ = protocol.SpecEdge(&MockSpecEdge{})
	_ = protocol.AssertionChain(&MockProtocol{})
	_ = statemanager.Manager(&MockStateManager{})
)

type MockAssertion struct {
	Prev           util.Option[*MockAssertion]
	MockHeight     uint64
	MockSeqNum     protocol.AssertionSequenceNumber
	MockPrevSeqNum protocol.AssertionSequenceNumber
	MockStateHash  common.Hash
}

func (m *MockAssertion) Height() (uint64, error) {
	return m.MockHeight, nil
}

func (m *MockAssertion) SeqNum() protocol.AssertionSequenceNumber {
	return m.MockSeqNum
}

func (m *MockAssertion) PrevSeqNum() (protocol.AssertionSequenceNumber, error) {
	return m.MockPrevSeqNum, nil
}

func (m *MockAssertion) StateHash() (common.Hash, error) {
	return m.MockStateHash, nil
}

type MockStateManager struct {
	mock.Mock
}

func (m *MockStateManager) LatestAssertionCreationData(ctx context.Context, prevHeight uint64) (*statemanager.AssertionToCreate, error) {
	args := m.Called(ctx, prevHeight)
	return args.Get(0).(*statemanager.AssertionToCreate), args.Error(1)
}

func (m *MockStateManager) HistoryCommitmentUpTo(ctx context.Context, height uint64) (util.HistoryCommitment, error) {
	args := m.Called(ctx, height)
	return args.Get(0).(util.HistoryCommitment), args.Error(1)
}

func (m *MockStateManager) PrefixProof(ctx context.Context, from, to uint64) ([]byte, error) {
	args := m.Called(ctx, from, to)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockStateManager) BigStepPrefixProof(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	fromBigStep,
	toBigStep uint64,
) ([]byte, error) {
	args := m.Called(ctx, fromBlockChallengeHeight, toBlockChallengeHeight, fromBigStep, toBigStep)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockStateManager) SmallStepPrefixProof(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	fromBigStep,
	toBigStep,
	fromSmallStep,
	toSmallStep uint64,
) ([]byte, error) {
	args := m.Called(ctx, fromBlockChallengeHeight, toBlockChallengeHeight, fromBigStep, toBigStep, fromSmallStep, toSmallStep)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockStateManager) HasStateCommitment(ctx context.Context, commit util.StateCommitment) bool {
	args := m.Called(ctx, commit)
	return args.Bool(0)
}

func (m *MockStateManager) BigStepLeafCommitment(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight uint64,
) (util.HistoryCommitment, error) {
	args := m.Called(ctx, fromBlockChallengeHeight, toBlockChallengeHeight)
	return args.Get(0).(util.HistoryCommitment), args.Error(1)
}

func (m *MockStateManager) BigStepCommitmentUpTo(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	toBigStep uint64,
) (util.HistoryCommitment, error) {
	args := m.Called(ctx, fromBlockChallengeHeight, toBlockChallengeHeight, toBigStep)
	return args.Get(0).(util.HistoryCommitment), args.Error(1)
}

func (m *MockStateManager) SmallStepLeafCommitment(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	fromBigStep,
	toBigStep uint64,
) (util.HistoryCommitment, error) {
	args := m.Called(ctx, fromBlockChallengeHeight, toBlockChallengeHeight, fromBigStep, toBigStep)
	return args.Get(0).(util.HistoryCommitment), args.Error(1)
}

func (m *MockStateManager) SmallStepCommitmentUpTo(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	fromBigStep,
	toBigStep,
	toSmallStep uint64,
) (util.HistoryCommitment, error) {
	args := m.Called(ctx, fromBlockChallengeHeight, toBlockChallengeHeight, fromBigStep, toBigStep, toSmallStep)
	return args.Get(0).(util.HistoryCommitment), args.Error(1)
}

func (m *MockStateManager) OneStepProofData(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	fromBigStep,
	toBigStep,
	fromSmallStep,
	toSmallStep uint64,
) (data *protocol.OneStepData, startLeafInclusionProof, endLeafInclusionProof []common.Hash, err error) {
	args := m.Called(ctx, fromBlockChallengeHeight, toBlockChallengeHeight, fromBigStep, toBigStep, fromSmallStep, toSmallStep)
	return args.Get(0).(*protocol.OneStepData), args.Get(1).([]common.Hash), args.Get(2).([]common.Hash), args.Error(3)
}

type MockChallengeManager struct {
	mock.Mock
	MockAddr common.Address
}

func (m *MockChallengeManager) ChallengePeriodSeconds(ctx context.Context) (time.Duration, error) {
	args := m.Called(ctx)
	return args.Get(0).(time.Duration), args.Error(1)
}

func (m *MockChallengeManager) Address() common.Address {
	return m.MockAddr
}

// MockSpecChallengeManager is a mock implementation of the SpecChallengeManager interface.
type MockSpecChallengeManager struct {
	mock.Mock
	MockAddr common.Address
}

func (m *MockSpecChallengeManager) Address() common.Address {
	return m.MockAddr
}

func (m *MockSpecChallengeManager) ChallengePeriodSeconds(ctx context.Context) (time.Duration, error) {
	args := m.Called(ctx)
	return args.Get(0).(time.Duration), args.Error(1)
}

func (m *MockSpecChallengeManager) GetEdge(
	ctx context.Context,
	edgeId protocol.EdgeId,
) (util.Option[protocol.SpecEdge], error) {
	args := m.Called(ctx, edgeId)
	return args.Get(0).(util.Option[protocol.SpecEdge]), args.Error(1)
}

func (m *MockSpecChallengeManager) CalculateMutualId(
	ctx context.Context,
	edgeType protocol.EdgeType,
	originId protocol.OriginId,
	startHeight protocol.Height,
	startHistoryRoot common.Hash,
	endHeight protocol.Height,
) (protocol.MutualId, error) {
	args := m.Called(ctx, edgeType, originId, startHeight, startHistoryRoot, endHeight)
	return args.Get(0).(protocol.MutualId), args.Error(1)
}

func (m *MockSpecChallengeManager) CalculateEdgeId(
	ctx context.Context,
	edgeType protocol.EdgeType,
	originId protocol.OriginId,
	startHeight protocol.Height,
	startHistoryRoot common.Hash,
	endHeight protocol.Height,
	endHistoryRoot common.Hash,
) (protocol.EdgeId, error) {
	args := m.Called(ctx, edgeType, originId, startHeight, startHistoryRoot, endHeight, endHistoryRoot)
	return args.Get(0).(protocol.EdgeId), args.Error(1)
}

func (m *MockSpecChallengeManager) AddBlockChallengeLevelZeroEdge(
	ctx context.Context,
	assertion protocol.Assertion,
	startCommit util.HistoryCommitment,
	endCommit util.HistoryCommitment,
) (protocol.SpecEdge, error) {
	args := m.Called(ctx, assertion, startCommit, endCommit)
	return args.Get(0).(protocol.SpecEdge), args.Error(1)
}

func (m *MockSpecChallengeManager) AddSubChallengeLevelZeroEdge(
	ctx context.Context,
	challengedEdge protocol.SpecEdge,
	startCommit util.HistoryCommitment,
	endCommit util.HistoryCommitment,
) (protocol.SpecEdge, error) {
	args := m.Called(ctx, challengedEdge, startCommit, endCommit)
	return args.Get(0).(protocol.SpecEdge), args.Error(1)
}
func (m *MockSpecChallengeManager) ConfirmEdgeByOneStepProof(
	ctx context.Context,
	tentativeWinnerId protocol.EdgeId,
	oneStepData *protocol.OneStepData,
	preHistoryInclusionProof []common.Hash,
	postHistoryInclusionProof []common.Hash,
) error {
	args := m.Called(ctx, tentativeWinnerId, oneStepData, preHistoryInclusionProof, postHistoryInclusionProof)
	return args.Error(0)
}

// MockSpecEdge is a mock implementation of the SpecEdge interface.
type MockSpecEdge struct {
	mock.Mock
}

func (m *MockSpecEdge) Id() protocol.EdgeId {
	args := m.Called()
	return args.Get(0).(protocol.EdgeId)
}
func (m *MockSpecEdge) GetType() protocol.EdgeType {
	args := m.Called()
	return args.Get(0).(protocol.EdgeType)
}
func (m *MockSpecEdge) MiniStaker() util.Option[common.Address] {
	args := m.Called()
	return args.Get(0).(util.Option[common.Address])
}
func (m *MockSpecEdge) StartCommitment() (protocol.Height, common.Hash) {
	args := m.Called()
	return args.Get(0).(protocol.Height), args.Get(1).(common.Hash)
}
func (m *MockSpecEdge) EndCommitment() (protocol.Height, common.Hash) {
	args := m.Called()
	return args.Get(0).(protocol.Height), args.Get(1).(common.Hash)
}
func (m *MockSpecEdge) TopLevelClaimHeight(ctx context.Context) (*protocol.OriginHeights, error) {
	args := m.Called(ctx)
	return args.Get(0).(*protocol.OriginHeights), args.Error(1)
}
func (m *MockSpecEdge) TimeUnrivaled(ctx context.Context) (uint64, error) {
	args := m.Called(ctx)
	return args.Get(0).(uint64), args.Error(1)
}
func (m *MockSpecEdge) HasRival(ctx context.Context) (bool, error) {
	args := m.Called(ctx)
	return args.Get(0).(bool), args.Error(1)
}
func (m *MockSpecEdge) Status(ctx context.Context) (protocol.EdgeStatus, error) {
	args := m.Called(ctx)
	return args.Get(0).(protocol.EdgeStatus), args.Error(1)
}
func (m *MockSpecEdge) Bisect(
	ctx context.Context,
	prefixHistoryRoot common.Hash,
	prefixProof []byte,
) (protocol.SpecEdge, protocol.SpecEdge, error) {
	args := m.Called(ctx, prefixHistoryRoot, prefixProof)
	return args.Get(0).(protocol.SpecEdge), args.Get(1).(protocol.SpecEdge), args.Error(2)
}
func (m *MockSpecEdge) ConfirmByTimer(ctx context.Context, ancestorIds []protocol.EdgeId) error {
	args := m.Called(ctx, ancestorIds)
	return args.Error(0)
}
func (m *MockSpecEdge) ConfirmByClaim(ctx context.Context, claimId protocol.ClaimId) error {
	args := m.Called(ctx, claimId)
	return args.Error(0)
}
func (m *MockSpecEdge) ConfirmByOneStepProof(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
func (m *MockSpecEdge) ConfirmByChildren(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
func (m *MockSpecEdge) HasLengthOneRival(ctx context.Context) (bool, error) {
	args := m.Called(ctx)
	return args.Get(0).(bool), args.Error(1)
}

type MockProtocol struct {
	mock.Mock
}

// Read-only methods.
func (m *MockProtocol) NumAssertions(ctx context.Context) (uint64, error) {
	args := m.Called(ctx)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockProtocol) AssertionBySequenceNum(ctx context.Context, seqNum protocol.AssertionSequenceNumber) (protocol.Assertion, error) {
	args := m.Called(ctx, seqNum)
	return args.Get(0).(protocol.Assertion), args.Error(1)
}

func (m *MockProtocol) GetAssertionId(ctx context.Context, seqNum protocol.AssertionSequenceNumber) (protocol.AssertionId, error) {
	args := m.Called(ctx, seqNum)
	return args.Get(0).(protocol.AssertionId), args.Error(1)
}

func (m *MockProtocol) GetAssertionNum(ctx context.Context, assertionHash protocol.AssertionId) (protocol.AssertionSequenceNumber, error) {
	args := m.Called(ctx, assertionHash)
	return args.Get(0).(protocol.AssertionSequenceNumber), args.Error(1)
}

func (m *MockProtocol) LatestConfirmed(ctx context.Context) (protocol.Assertion, error) {
	args := m.Called(ctx)
	return args.Get(0).(protocol.Assertion), args.Error(1)
}

// Mutating methods.
func (m *MockProtocol) CreateAssertion(ctx context.Context, height uint64, prevSeqNum protocol.AssertionSequenceNumber, prevAssertionState *protocol.ExecutionState, postState *protocol.ExecutionState, prevInboxMaxCount *big.Int) (protocol.Assertion, error) {
	args := m.Called(ctx, height, prevSeqNum, prevAssertionState, postState, prevInboxMaxCount)
	return args.Get(0).(protocol.Assertion), args.Error(1)
}

func (m *MockProtocol) SpecChallengeManager(ctx context.Context) (protocol.SpecChallengeManager, error) {
	args := m.Called(ctx)
	return args.Get(0).(protocol.SpecChallengeManager), args.Error(1)
}

func (m *MockProtocol) CreateSpecChallenge(ctx context.Context, seqNum protocol.AssertionSequenceNumber) error {
	args := m.Called(ctx, seqNum)
	return args.Error(0)
}

func (m *MockProtocol) Confirm(ctx context.Context, blockHash, sendRoot common.Hash) error {
	args := m.Called(ctx, blockHash, sendRoot)
	return args.Error(0)
}
