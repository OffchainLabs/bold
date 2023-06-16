package toys

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"

	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	l2stateprovider "github.com/OffchainLabs/challenge-protocol-v2/layer2-state-provider"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/rollupgen"
	commitments "github.com/OffchainLabs/challenge-protocol-v2/state-commitments/history"
	prefixproofs "github.com/OffchainLabs/challenge-protocol-v2/state-commitments/prefix-proofs"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Defines the ABI encoding structure for submission of prefix proofs to the protocol contracts
var (
	b32Arr, _ = abi.NewType("bytes32[]", "", nil)
	// ProofArgs for submission to the protocol.
	ProofArgs = abi.Arguments{
		{Type: b32Arr, Name: "prefixExpansion"},
		{Type: b32Arr, Name: "prefixProof"},
	}
)

// L2StateBackend defines a very naive state manager that is initialized from a list of predetermined
// state roots. It can produce state and history commitments from those roots.
type L2StateBackend struct {
	stateRoots              []common.Hash
	executionStates         []*protocol.ExecutionState
	machineAtBlock          func(context.Context, uint64) (Machine, error)
	maxWavmOpcodes          uint64
	numOpcodesPerBigStep    uint64
	blockDivergenceHeight   uint64
	posInBatchDivergence    int64
	machineDivergenceStep   uint64
	forceMachineBlockCompat bool
}

// New simulated manager from a list of predefined state roots, useful for tests and simulations.
func New(stateRoots []common.Hash, opts ...Opt) (*L2StateBackend, error) {
	if len(stateRoots) == 0 {
		return nil, errors.New("no state roots provided")
	}
	s := &L2StateBackend{
		stateRoots: stateRoots,
		machineAtBlock: func(context.Context, uint64) (Machine, error) {
			return nil, errors.New("state manager created with New() cannot provide machines")
		},
	}
	for _, o := range opts {
		o(s)
	}
	return s, nil
}

type Opt func(*L2StateBackend)

func WithMaxWavmOpcodesPerBlock(maxOpcodes uint64) Opt {
	return func(s *L2StateBackend) {
		s.maxWavmOpcodes = maxOpcodes
	}
}

func WithNumOpcodesPerBigStep(numOpcodes uint64) Opt {
	return func(s *L2StateBackend) {
		s.numOpcodesPerBigStep = numOpcodes
	}
}

func WithMachineDivergenceStep(divergenceStep uint64) Opt {
	return func(s *L2StateBackend) {
		s.machineDivergenceStep = divergenceStep
	}
}

func WithBlockDivergenceHeight(divergenceHeight uint64) Opt {
	return func(s *L2StateBackend) {
		s.blockDivergenceHeight = divergenceHeight
	}
}

func WithDivergentBlockHeightOffset(blockHeightOffset int64) Opt {
	return func(s *L2StateBackend) {
		s.posInBatchDivergence = blockHeightOffset * 150
	}
}

func WithMachineAtBlockProvider(machineAtBlock func(ctx context.Context, blockNum uint64) (Machine, error)) Opt {
	return func(s *L2StateBackend) {
		s.machineAtBlock = machineAtBlock
	}
}

// If enabled, forces the machine hash at block boundaries to be the block hash
func WithForceMachineBlockCompat() Opt {
	return func(s *L2StateBackend) {
		s.forceMachineBlockCompat = true
	}
}

// NewWithAssertionStates creates a simulated state manager from a list of predefined state roots for
// the top-level assertion chain, useful for tests and simulation purposes in block challenges.
// This also allows for specifying the honest states for big and small step subchallenges along
// with the point at which the state manager should diverge from the honest computation.
func NewWithAssertionStates(
	assertionChainExecutionStates []*protocol.ExecutionState,
	opts ...Opt,
) (*L2StateBackend, error) {
	if len(assertionChainExecutionStates) == 0 {
		return nil, errors.New("must have execution states")
	}
	stateRoots := make([]common.Hash, len(assertionChainExecutionStates))
	var lastBatch uint64 = math.MaxUint64
	var lastPosInBatch uint64 = math.MaxUint64
	for i := 0; i < len(stateRoots); i++ {
		state := assertionChainExecutionStates[i]
		if state.GlobalState.Batch == lastBatch && state.GlobalState.PosInBatch == lastPosInBatch {
			return nil, fmt.Errorf("execution states %v and %v have the same batch %v and position in batch %v", i-1, i, lastBatch, lastPosInBatch)
		}
		lastBatch = state.GlobalState.Batch
		lastPosInBatch = state.GlobalState.PosInBatch
		stateRoots[i] = protocol.ComputeSimpleMachineChallengeHash(state)
	}
	s := &L2StateBackend{
		stateRoots:      stateRoots,
		executionStates: assertionChainExecutionStates,
		machineAtBlock: func(context.Context, uint64) (Machine, error) {
			return nil, errors.New("state manager created with NewWithAssertionStates() cannot provide machines")
		},
	}
	for _, o := range opts {
		o(s)
	}
	return s, nil
}

func NewForSimpleMachine(
	opts ...Opt,
) (*L2StateBackend, error) {
	s := &L2StateBackend{
		maxWavmOpcodes:       protocol.LevelZeroSmallStepEdgeHeight * protocol.LevelZeroBigStepEdgeHeight,
		numOpcodesPerBigStep: protocol.LevelZeroSmallStepEdgeHeight,
	}
	for _, o := range opts {
		o(s)
	}
	if s.maxWavmOpcodes == 0 {
		return nil, errors.New("maxWavmOpcodes cannot be zero")
	}
	if s.blockDivergenceHeight > 0 && s.machineDivergenceStep == 0 {
		return nil, errors.New("machineDivergenceStep cannot be zero if blockDivergenceHeight is non-zero")
	}
	nextMachineState := &protocol.ExecutionState{
		GlobalState:   protocol.GoGlobalState{},
		MachineStatus: protocol.MachineStatusFinished,
	}
	maxBatchesRead := big.NewInt(1)
	for block := uint64(0); ; block++ {
		machine := NewSimpleMachine(nextMachineState, maxBatchesRead)
		state := machine.GetExecutionState()
		machHash := machine.Hash()
		if machHash != state.GlobalState.Hash() {
			return nil, fmt.Errorf("machine at block %v has hash %v but we expected hash %v", block, machine.Hash(), state.GlobalState.Hash())
		}
		if s.blockDivergenceHeight > 0 {
			if block == s.blockDivergenceHeight {
				// Note: blockHeightOffset might be negative, but two's complement subtraction works regardless
				state.GlobalState.PosInBatch -= uint64(s.posInBatchDivergence)
			}
			if block >= s.blockDivergenceHeight {
				state.GlobalState.BlockHash[0] = 1
			}
			machHash = protocol.ComputeSimpleMachineChallengeHash(state)
		}
		s.executionStates = append(s.executionStates, state)
		s.stateRoots = append(s.stateRoots, machHash)

		if machine.IsStopped() || state.GlobalState.Batch >= 1 {
			break
		}
		err := machine.Step(s.maxWavmOpcodes)
		if err != nil {
			return nil, err
		}
		nextMachineState = machine.GetExecutionState()
	}
	s.machineAtBlock = func(_ context.Context, block uint64) (Machine, error) {
		if block >= uint64(len(s.executionStates)) {
			block = uint64(len(s.executionStates) - 1)
		}
		return NewSimpleMachine(s.executionStates[block], maxBatchesRead), nil
	}
	return s, nil
}

// Produces the latest state to assert to L1 from the local state manager's perspective.
func (s *L2StateBackend) LatestExecutionState(_ context.Context) (*protocol.ExecutionState, error) {
	if len(s.executionStates) == 0 {
		return nil, errors.New("no execution states")
	}
	return s.executionStates[len(s.executionStates)-1], nil
}

// Checks if the execution manager locally has recorded this state
func (s *L2StateBackend) ExecutionStateBlockHeight(_ context.Context, state *protocol.ExecutionState) (uint64, bool) {
	for i, r := range s.executionStates {
		if r.Equals(state) {
			return uint64(i), true
		}
	}
	return 0, false
}

func (s *L2StateBackend) HistoryCommitmentUpTo(_ context.Context, blockChallengeHeight uint64) (commitments.History, error) {
	// The size is the number of elements being committed to. For example, if the height is 7, there will
	// be 8 elements being committed to from [0, 7] inclusive.
	size := blockChallengeHeight + 1
	return commitments.New(
		blockChallengeHeight,
		s.stateRoots[:size],
	)
}

func (s *L2StateBackend) statesUpTo(blockStart, blockEnd, nextBatchCount uint64) ([]common.Hash, error) {
	if blockEnd < blockStart {
		return nil, fmt.Errorf("end block %v is less than start block %v", blockEnd, blockStart)
	}
	// The size is the number of elements being committed to. For example, if the height is 7, there will
	// be 8 elements being committed to from [0, 7] inclusive.
	desiredStatesLen := int(blockEnd - blockStart + 1)
	var states []common.Hash
	var lastState common.Hash
	for i := blockStart; i <= blockEnd; i++ {
		if i >= uint64(len(s.stateRoots)) {
			break
		}
		state := s.stateRoots[i]
		states = append(states, state)
		lastState = state
		if len(s.executionStates) == 0 {
			// should only happen in tests
			continue
		}
		gs := s.executionStates[i].GlobalState
		if gs.Batch >= nextBatchCount {
			if gs.Batch > nextBatchCount || gs.PosInBatch > 0 {
				return nil, fmt.Errorf("overran next batch count %v with global state batch %v position %v", nextBatchCount, gs.Batch, gs.PosInBatch)
			}
			break
		}
	}
	for len(states) < desiredStatesLen {
		states = append(states, lastState)
	}
	return states, nil
}

func (s *L2StateBackend) HistoryCommitmentUpToBatch(_ context.Context, blockStart, blockEnd, nextBatchCount uint64) (commitments.History, error) {
	states, err := s.statesUpTo(blockStart, blockEnd, nextBatchCount)
	if err != nil {
		return commitments.History{}, err
	}
	return commitments.New(
		blockEnd-blockStart,
		states,
	)
}

func (s *L2StateBackend) AgreesWithHistoryCommitment(
	_ context.Context,
	_ protocol.EdgeType,
	_ uint64,
	_ *protocol.OriginHeights,
	_,
	_ commitments.History,
) (protocol.Agreement, error) {
	return protocol.Agreement{}, nil
}

func (s *L2StateBackend) BigStepLeafCommitment(
	ctx context.Context,
	fromAssertionHeight,
	toAssertionHeight uint64,
) (commitments.History, error) {
	// Number of big steps between assertion heights A and B will be
	// fixed in this simulated state manager. It is simply the max number of opcodes
	// per block divided by the size of a big step.
	numBigSteps := s.maxWavmOpcodes / s.numOpcodesPerBigStep
	return s.BigStepCommitmentUpTo(
		ctx,
		fromAssertionHeight,
		toAssertionHeight,
		numBigSteps,
	)
}

func (s *L2StateBackend) BigStepCommitmentUpTo(
	ctx context.Context,
	fromAssertionHeight,
	toAssertionHeight,
	toBigStep uint64,
) (commitments.History, error) {
	if fromAssertionHeight+1 != toAssertionHeight {
		return commitments.History{}, fmt.Errorf(
			"from height %d is not one-step away from to height %d",
			fromAssertionHeight,
			toAssertionHeight,
		)
	}
	leaves, err := s.intermediateBigStepLeaves(
		ctx,
		fromAssertionHeight,
		toAssertionHeight,
		0, // from big step.
		toBigStep,
	)
	if err != nil {
		return commitments.History{}, err
	}
	return commitments.New(toBigStep, leaves)
}

func (s *L2StateBackend) maybeDivergeState(state *protocol.ExecutionState, block uint64, step uint64) {
	if block+1 == s.blockDivergenceHeight && step == s.maxWavmOpcodes {
		*state = *s.executionStates[block+1]
	}
	if block+1 > s.blockDivergenceHeight || step >= s.machineDivergenceStep {
		state.GlobalState.BlockHash[0] = 1
	}
}

// May modify the machine hash if divergence is enabled
func (s *L2StateBackend) getMachineHash(machine Machine, block uint64) common.Hash {
	if s.forceMachineBlockCompat {
		step := machine.CurrentStepNum()
		if step == 0 {
			return s.stateRoots[block]
		}
		if step == s.maxWavmOpcodes {
			return s.stateRoots[block+1]
		}
	}
	if s.blockDivergenceHeight == 0 || block+1 < s.blockDivergenceHeight {
		return machine.Hash()
	}
	state := machine.GetExecutionState()
	s.maybeDivergeState(state, block, machine.CurrentStepNum())
	return protocol.ComputeSimpleMachineChallengeHash(state)
}

func (s *L2StateBackend) intermediateBigStepLeaves(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	fromBigStep,
	toBigStep uint64,
) ([]common.Hash, error) {
	if toBlockChallengeHeight != fromBlockChallengeHeight+1 {
		return nil, fmt.Errorf("attempting to get big step leaves from block %v to %v", fromBlockChallengeHeight, toBlockChallengeHeight)
	}
	leaves := make([]common.Hash, 0)
	machine, err := s.machineAtBlock(ctx, fromBlockChallengeHeight)
	if err != nil {
		return nil, err
	}
	// Up to and including the specified step.
	for i := fromBigStep; i <= toBigStep; i++ {
		leaves = append(leaves, s.getMachineHash(machine, fromBlockChallengeHeight))
		if i >= toBigStep {
			// We don't need to step the machine to the next point because it won't be used
			break
		}
		err = machine.Step(s.numOpcodesPerBigStep)
		if err != nil {
			return nil, err
		}
	}
	return leaves, nil
}

func (s *L2StateBackend) SmallStepLeafCommitment(
	ctx context.Context,
	fromAssertionHeight,
	toAssertionHeight,
	fromBigStep,
	toBigStep uint64,
) (commitments.History, error) {
	return s.SmallStepCommitmentUpTo(
		ctx,
		fromAssertionHeight,
		toAssertionHeight,
		fromBigStep,
		toBigStep,
		s.numOpcodesPerBigStep,
	)
}

func (s *L2StateBackend) SmallStepCommitmentUpTo(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	fromBigStep,
	toBigStep,
	toSmallStep uint64,
) (commitments.History, error) {
	if fromBlockChallengeHeight+1 != toBlockChallengeHeight {
		return commitments.History{}, fmt.Errorf(
			"from height %d is not one-step away from to height %d",
			fromBlockChallengeHeight,
			toBlockChallengeHeight,
		)
	}
	if fromBigStep+1 != toBigStep {
		return commitments.History{}, fmt.Errorf(
			"from height %d is not one-step away from to height %d",
			fromBigStep,
			toBigStep,
		)
	}

	fromSmall := fromBigStep * s.numOpcodesPerBigStep
	toSmall := fromSmall + toSmallStep
	leaves, err := s.intermediateSmallStepLeaves(
		ctx,
		fromBlockChallengeHeight,
		toBlockChallengeHeight,
		fromSmall,
		toSmall,
	)
	if err != nil {
		return commitments.History{}, err
	}
	return commitments.New(toSmallStep, leaves)
}

func (s *L2StateBackend) intermediateSmallStepLeaves(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	fromSmallStep,
	toSmallStep uint64,
) ([]common.Hash, error) {
	if toBlockChallengeHeight != fromBlockChallengeHeight+1 {
		return nil, fmt.Errorf("attempting to get small step leaves from block %v to %v", fromBlockChallengeHeight, toBlockChallengeHeight)
	}
	leaves := make([]common.Hash, 0)
	machine, err := s.machineAtBlock(ctx, fromBlockChallengeHeight)
	if err != nil {
		return nil, err
	}
	err = machine.Step(fromSmallStep)
	if err != nil {
		return nil, err
	}
	for i := fromSmallStep; i <= toSmallStep; i++ {
		leaves = append(leaves, s.getMachineHash(machine, fromBlockChallengeHeight))
		if i >= toSmallStep {
			// We don't need to step the machine to the next point because it won't be used
			break
		}
		err = machine.Step(1)
		if err != nil {
			return nil, err
		}
	}
	return leaves, nil
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
var uint64Type = newStaticType("uint64", "", nil)
var uint8Type = newStaticType("uint8", "", nil)
var addressType = newStaticType("address", "", nil)
var uint256Type = newStaticType("uint256", "", nil)

var WasmModuleProofAbi = abi.Arguments{
	{
		Name: "requiredStake",
		Type: uint256Type,
	},
	{
		Name: "challengeManager",
		Type: addressType,
	},
	{
		Name: "confirmPeriodBlocks",
		Type: uint64Type,
	},
}

var ExecutionStateAbi = abi.Arguments{
	{
		Name: "b1",
		Type: bytes32Type,
	},
	{
		Name: "b2",
		Type: bytes32Type,
	},
	{
		Name: "u1",
		Type: uint64Type,
	},
	{
		Name: "u2",
		Type: uint64Type,
	},
	{
		Name: "status",
		Type: uint8Type,
	},
}

func (s *L2StateBackend) OneStepProofData(
	ctx context.Context,
	cfgSnapshot *l2stateprovider.ConfigSnapshot,
	postState rollupgen.ExecutionState,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	fromBigStep,
	toBigStep,
	fromSmallStep,
	toSmallStep uint64,
) (data *protocol.OneStepData, startLeafInclusionProof, endLeafInclusionProof []common.Hash, err error) {
	inboxMaxCountProof, packErr := ExecutionStateAbi.Pack(
		postState.GlobalState.Bytes32Vals[0],
		postState.GlobalState.Bytes32Vals[1],
		postState.GlobalState.U64Vals[0],
		postState.GlobalState.U64Vals[1],
		postState.MachineStatus,
	)
	if packErr != nil {
		err = packErr
		return
	}

	wasmModuleRootProof, packErr := WasmModuleProofAbi.Pack(
		cfgSnapshot.RequiredStake,
		cfgSnapshot.ChallengeManagerAddress,
		cfgSnapshot.ConfirmPeriodBlocks,
	)
	if packErr != nil {
		err = packErr
		return
	}
	startCommit, commitErr := s.SmallStepCommitmentUpTo(
		ctx,
		fromBlockChallengeHeight,
		toBlockChallengeHeight,
		fromBigStep,
		toBigStep,
		fromSmallStep,
	)
	if commitErr != nil {
		err = commitErr
		return
	}
	endCommit, commitErr := s.SmallStepCommitmentUpTo(
		ctx,
		fromBlockChallengeHeight,
		toBlockChallengeHeight,
		fromBigStep,
		toBigStep,
		toSmallStep,
	)
	if commitErr != nil {
		err = commitErr
		return
	}

	machine, machineErr := s.machineAtBlock(ctx, fromBlockChallengeHeight)
	if machineErr != nil {
		err = machineErr
		return
	}
	step := fromBigStep*s.numOpcodesPerBigStep + fromSmallStep
	err = machine.Step(step)
	if err != nil {
		return
	}
	beforeHash := machine.Hash()
	if beforeHash != startCommit.LastLeaf {
		err = fmt.Errorf("machine executed to start step %v hash %v but expected %v", step, beforeHash, startCommit.LastLeaf)
		return
	}
	osp, ospErr := machine.OneStepProof()
	if ospErr != nil {
		err = ospErr
		return
	}
	err = machine.Step(1)
	if err != nil {
		return
	}
	afterHash := machine.Hash()
	if afterHash != endCommit.LastLeaf {
		err = fmt.Errorf("machine executed to end step %v hash %v but expected %v", step+1, beforeHash, endCommit.LastLeaf)
		return
	}

	data = &protocol.OneStepData{
		BeforeHash:             startCommit.LastLeaf,
		Proof:                  osp,
		InboxMsgCountSeen:      cfgSnapshot.InboxMaxCount,
		InboxMsgCountSeenProof: inboxMaxCountProof,
		WasmModuleRoot:         cfgSnapshot.WasmModuleRoot,
		WasmModuleRootProof:    wasmModuleRootProof,
	}
	startLeafInclusionProof = startCommit.LastLeafProof
	endLeafInclusionProof = endCommit.LastLeafProof
	return
}

func (s *L2StateBackend) prefixProofImpl(_ context.Context, start, lo, hi, batchCount uint64) ([]byte, error) {
	states, err := s.statesUpTo(start, hi, batchCount)
	if err != nil {
		return nil, err
	}
	loSize := lo + 1 - start
	hiSize := hi + 1 - start
	prefixExpansion, err := prefixproofs.ExpansionFromLeaves(states[:loSize])
	if err != nil {
		return nil, err
	}
	prefixProof, err := prefixproofs.GeneratePrefixProof(
		loSize,
		prefixExpansion,
		states[loSize:hiSize],
		prefixproofs.RootFetcherFromExpansion,
	)
	if err != nil {
		return nil, err
	}
	_, numRead := prefixproofs.MerkleExpansionFromCompact(prefixProof, loSize)
	onlyProof := prefixProof[numRead:]
	return ProofArgs.Pack(&prefixExpansion, &onlyProof)
}

func (s *L2StateBackend) PrefixProof(ctx context.Context, lo, hi uint64) ([]byte, error) {
	return s.prefixProofImpl(ctx, 0, lo, hi, math.MaxUint64)
}

func (s *L2StateBackend) PrefixProofUpToBatch(ctx context.Context, start, lo, hi, batchCount uint64) ([]byte, error) {
	return s.prefixProofImpl(ctx, start, lo, hi, batchCount)
}

func (s *L2StateBackend) BigStepPrefixProof(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	fromBigStep,
	toBigStep uint64,
) ([]byte, error) {
	if fromBlockChallengeHeight+1 != toBlockChallengeHeight {
		return nil, fmt.Errorf(
			"fromAssertionHeight=%d is not 1 height apart from toAssertionHeight=%d",
			fromBlockChallengeHeight,
			toBlockChallengeHeight,
		)
	}
	return s.bigStepPrefixProofCalculation(
		ctx,
		fromBlockChallengeHeight,
		toBlockChallengeHeight,
		fromBigStep,
		toBigStep,
	)
}

func (s *L2StateBackend) bigStepPrefixProofCalculation(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	fromBigStep,
	toBigStep uint64,
) ([]byte, error) {
	loSize := fromBigStep + 1
	hiSize := toBigStep + 1
	prefixLeaves, err := s.intermediateBigStepLeaves(
		ctx,
		fromBlockChallengeHeight,
		toBlockChallengeHeight,
		0,
		toBigStep,
	)
	if err != nil {
		return nil, err
	}
	prefixExpansion, err := prefixproofs.ExpansionFromLeaves(prefixLeaves[:loSize])
	if err != nil {
		return nil, err
	}
	prefixProof, err := prefixproofs.GeneratePrefixProof(
		loSize,
		prefixExpansion,
		prefixLeaves[loSize:hiSize],
		prefixproofs.RootFetcherFromExpansion,
	)
	if err != nil {
		return nil, err
	}
	_, numRead := prefixproofs.MerkleExpansionFromCompact(prefixProof, loSize)
	onlyProof := prefixProof[numRead:]
	return ProofArgs.Pack(&prefixExpansion, &onlyProof)
}

func (s *L2StateBackend) SmallStepPrefixProof(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	fromBigStep,
	toBigStep,
	fromSmallStep,
	toSmallStep uint64,
) ([]byte, error) {
	if fromBlockChallengeHeight+1 != toBlockChallengeHeight {
		return nil, fmt.Errorf(
			"fromAssertionHeight=%d is not 1 height apart from toAssertionHeight=%d",
			fromBlockChallengeHeight,
			toBlockChallengeHeight,
		)
	}
	if fromBigStep+1 != toBigStep {
		return nil, fmt.Errorf(
			"fromBigStep=%d is not 1 height apart from toBigStep=%d",
			fromBigStep,
			toBigStep,
		)
	}
	return s.smallStepPrefixProofCalculation(
		ctx,
		fromBlockChallengeHeight,
		toBlockChallengeHeight,
		fromBigStep,
		fromSmallStep,
		toSmallStep,
	)
}

func (s *L2StateBackend) smallStepPrefixProofCalculation(
	ctx context.Context,
	fromBlockChallengeHeight,
	toBlockChallengeHeight,
	fromBigStep,
	fromSmallStep,
	toSmallStep uint64,
) ([]byte, error) {
	fromSmall := fromBigStep * s.numOpcodesPerBigStep
	toSmall := fromSmall + toSmallStep
	prefixLeaves, err := s.intermediateSmallStepLeaves(
		ctx,
		fromBlockChallengeHeight,
		toBlockChallengeHeight,
		fromSmall,
		toSmall,
	)
	if err != nil {
		return nil, err
	}
	loSize := fromSmallStep + 1
	hiSize := toSmallStep + 1
	prefixExpansion, err := prefixproofs.ExpansionFromLeaves(prefixLeaves[:loSize])
	if err != nil {
		return nil, err
	}
	prefixProof, err := prefixproofs.GeneratePrefixProof(
		loSize,
		prefixExpansion,
		prefixLeaves[loSize:hiSize],
		prefixproofs.RootFetcherFromExpansion,
	)
	if err != nil {
		return nil, err
	}
	_, numRead := prefixproofs.MerkleExpansionFromCompact(prefixProof, loSize)
	onlyProof := prefixProof[numRead:]
	return ProofArgs.Pack(&prefixExpansion, &onlyProof)
}
