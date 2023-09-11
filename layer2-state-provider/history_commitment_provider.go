package l2stateprovider

import (
	"context"
	"errors"
	"fmt"
	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/containers/option"
	"github.com/OffchainLabs/bold/solgen/go/rollupgen"
	commitments "github.com/OffchainLabs/bold/state-commitments/history"
	prefixproofs "github.com/OffchainLabs/bold/state-commitments/prefix-proofs"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	emptyCommit = commitments.History{}
)

// MachineHashCollector defines a struct which can collect hashes from an Arbitrator machine
// at a block height, starting at a specific opcode index in the machine and stepping through it
// in increments of custom size. Along the way, it computes each machine hash at each step
// and outputs a list of these hashes at the end. This is an computationally expensive process
// that is best performed if machine hashes are cached after runs.
type MachineHashCollector interface {
	CollectMachineMashes(ctx context.Context, cfg *HashCollectorConfig) ([]common.Hash, error)
}

// HashCollectorConfig defines configuration options for a machine hash collector to
// step through an Arbitrator machine at a specific L2 message number, step through it in
// increments of `StepSize`, and collect the machine hashes at those steps into an output slice.
type HashCollectorConfig struct {
	// The WASM module root the machines should be a part of.
	WasmModuleRoot common.Hash
	// The L2 message number the machine corresponds to.
	MessageNumber Height
	StepHeights   []Height
	// The number of desired hashes to be collected.
	NumDesiredHashes uint64
	// The opcode index at which to start stepping through the machine at a message number.
	MachineStartIndex OpcodeIndex
	// The step size for stepping through the machine in order to collect its hashes.
	StepSize StepSize
}

// L2MessageStateCollector defines a struct which can obtain the machine hashes at each L2 message
// in a specified message range for a given batch index on Arbitrum.
type L2MessageStateCollector interface {
	L2MessageStatesUpTo(
		ctx context.Context,
		from Height,
		upTo option.Option[Height],
		batch Batch,
	) ([]common.Hash, error)
}

// HistoryCommitmentProvider computes history commitments from input parameters
// by loading Arbitrator machines for L2 state transitions. It can compute history commitments
// over ranges of opcodes at specified increments used for the BOLD protocol.
type HistoryCommitmentProvider struct {
	l2MessageStateCollector L2MessageStateCollector
	machineHashCollector    MachineHashCollector
	challengeLeafHeights    []Height
}

// NewHistoryCommitmentProvider --
func NewHistoryCommitmentProvider(
	l2MessageStateCollector L2MessageStateCollector,
	machineHashCollector MachineHashCollector,
	challengeLeafHeights []Height,
) *HistoryCommitmentProvider {
	return &HistoryCommitmentProvider{
		l2MessageStateCollector: l2MessageStateCollector,
		machineHashCollector:    machineHashCollector,
		challengeLeafHeights:    challengeLeafHeights,
	}
}

// A list of heights that have been validated to be non-empty
// and to be < the total number of challenge levels in the protocol.
type validatedStartHeights []Height

// HistoryCommitment computes a Merklelized commitment over a set of hashes
// at specified challenge levels. For block challenges, for example, this is a set
// of machine hashes corresponding each message in a range N to M.
func (p *HistoryCommitmentProvider) HistoryCommitment(
	ctx context.Context,
	wasmModuleRoot common.Hash,
	batch Batch,
	startHeights []Height,
	upToHeight option.Option[Height],
) (commitments.History, error) {
	hashes, err := p.historyCommitmentImpl(ctx, wasmModuleRoot, batch, startHeights, upToHeight)
	if err != nil {
		return commitments.History{}, err
	}
	return commitments.New(hashes)
}

func (p *HistoryCommitmentProvider) historyCommitmentImpl(
	ctx context.Context,
	wasmModuleRoot common.Hash,
	batch Batch,
	startHeights []Height,
	upToHeight option.Option[Height],
) ([]common.Hash, error) {
	// Validate the input heights for correctness.
	validatedHeights, err := p.validateStartHeights(startHeights)
	if err != nil {
		return nil, err
	}
	// If the call is for message number ranges only, we get the hashes for
	// those states and return a commitment for them.
	fromMessageNumber := uint64(validatedHeights[0])
	if len(validatedHeights) == 1 {
		hashes, hashesErr := p.l2MessageStateCollector.L2MessageStatesUpTo(ctx, Height(fromMessageNumber), upToHeight, batch)
		if hashesErr != nil {
			return nil, hashesErr
		}
		return hashes, nil
	}

	// Next, computes the exact start point of where we need to execute
	// the machine from the inputs, and figures out in what increments we need to do so.
	machineStartIndex := p.computeMachineStartIndex(validatedHeights)

	// We compute the stepwise increments we need for stepping through the machine.
	stepSize, err := p.computeStepSize(validatedHeights)
	if err != nil {
		return nil, err
	}

	// Compute how many machine hashes we need to collect.
	numHashes, err := p.computeRequiredNumberOfHashes(validatedHeights, upToHeight)
	if err != nil {
		return nil, err
	}

	// Collect the machine hashes at the specified challenge level based on the values we computed.
	hashes, err := p.machineHashCollector.CollectMachineMashes(
		ctx,
		&HashCollectorConfig{
			WasmModuleRoot:    wasmModuleRoot,
			MessageNumber:     Height(fromMessageNumber),
			StepHeights:       validatedHeights[1:],
			NumDesiredHashes:  numHashes,
			MachineStartIndex: machineStartIndex,
			StepSize:          stepSize,
		},
	)
	if err != nil {
		return nil, err
	}
	return hashes, nil
}

// AgreesWithHistoryCommitment checks if the l2 state provider agrees with a specified start and end
// history commitment for a type of edge under a specified assertion challenge. It returns an agreement struct
// which informs the caller whether (a) we agree with the start commitment, and whether (b) the edge is honest, meaning
// that we also agree with the end commitment.
func (p *HistoryCommitmentProvider) AgreesWithHistoryCommitment(
	ctx context.Context,
	wasmModuleRoot common.Hash,
	assertionInboxMaxCount uint64,
	parentAssertionAfterStateBatch uint64,
	edgeType protocol.ChallengeLevel,
	heights protocol.OriginHeights,
	commit History,
) (bool, error) {
	var localCommit commitments.History
	var err error

	switch edgeType {
	case protocol.NewBlockChallengeLevel():
		localCommit, err = p.HistoryCommitment(
			ctx,
			wasmModuleRoot,
			Batch(assertionInboxMaxCount),
			[]Height{Height(parentAssertionAfterStateBatch)},
			option.Some[Height](Height(parentAssertionAfterStateBatch+commit.Height)))
		if err != nil {
			return false, err
		}
	default:
		challengeOriginHeights := make([]Height, len(heights.ChallengeOriginHeights))
		for index, height := range heights.ChallengeOriginHeights {
			challengeOriginHeights[index] = Height(height)
		}
		localCommit, err = p.HistoryCommitment(
			ctx,
			wasmModuleRoot,
			Batch(assertionInboxMaxCount),
			append(challengeOriginHeights, 0),
			option.Some[Height](Height(commit.Height)))
		if err != nil {
			return false, err
		}
	}
	return localCommit.Height == commit.Height && localCommit.Merkle == commit.MerkleRoot, nil
}

var (
	b32Arr, _ = abi.NewType("bytes32[]", "", nil)
	// ProofArgs for submission to the protocol.
	ProofArgs = abi.Arguments{
		{Type: b32Arr, Name: "prefixExpansion"},
		{Type: b32Arr, Name: "prefixProof"},
	}
)

func (p *HistoryCommitmentProvider) PrefixProof(
	ctx context.Context,
	wasmModuleRoot common.Hash,
	batch Batch,
	startHeights []Height,
	fromMessageNumber Height,
	upToHeight option.Option[Height],
) ([]byte, error) {
	prefixLeaves, err := p.historyCommitmentImpl(
		ctx,
		wasmModuleRoot,
		batch,
		startHeights,
		upToHeight)
	if err != nil {
		return nil, err
	}
	loSize := uint64(fromMessageNumber + 1)
	hiSize := uint64(upToHeight.Unwrap() + 1)
	if len(startHeights) == 1 {
		loSize -= uint64(fromMessageNumber)
		hiSize -= uint64(fromMessageNumber)
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

func (p *HistoryCommitmentProvider) OneStepProofData(
	ctx context.Context,
	wasmModuleRoot common.Hash,
	postState rollupgen.ExecutionState,
	startHeights []Height,
	upToHeight option.Option[Height],
) (data *protocol.OneStepData, startLeafInclusionProof, endLeafInclusionProof []common.Hash, err error) {
	startCommit, commitErr := p.SmallStepCommitmentUpTo(
		ctx,
		wasmModuleRoot,
		messageNumber,
		bigStep,
		smallStep,
	)
	if commitErr != nil {
		err = commitErr
		return
	}
	endCommit, commitErr := p.SmallStepCommitmentUpTo(
		ctx,
		wasmModuleRoot,
		messageNumber,
		bigStep,
		smallStep+1,
	)
	if commitErr != nil {
		err = commitErr
		return
	}

	machine, machineErr := s.machineAtBlock(ctx, messageNumber)
	if machineErr != nil {
		err = machineErr
		return
	}
	step := bigStep*s.numOpcodesPerBigStep + smallStep
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
		BeforeHash: startCommit.LastLeaf,
		Proof:      osp,
	}
	startLeafInclusionProof = startCommit.LastLeafProof
	endLeafInclusionProof = endCommit.LastLeafProof
	return
}

// Computes the required number of hashes for a history commitment
// based on the requested heights and challenge level. The required number of hashes
// for a leaf commitment at each challenge level is a constant, so we can determine
// the desired challenge level from the input params and compute the total
// from there.
func (p *HistoryCommitmentProvider) computeRequiredNumberOfHashes(
	startHeights validatedStartHeights,
	upToHeight option.Option[Height],
) (uint64, error) {
	// Get the max number of hashes at the specified challenge level.
	// from the protocol constants.
	challengeLevel := len(startHeights) - 1
	maxHeightForLevel := p.challengeLeafHeights[challengeLevel]

	// Get the start height we want to use at the challenge level.
	start := startHeights[challengeLevel]

	var end Height
	if upToHeight.IsNone() {
		end = maxHeightForLevel
	} else {
		end = upToHeight.Unwrap()
		// If the end height is more than the allowed max, we return an error.
		// This scenario should not happen, and instead of silently truncating,
		// surfacing an error is the safest way of warning the operator
		// they are committing something invalid.
		if end > maxHeightForLevel {
			return 0, fmt.Errorf(
				"end %d was greater than max height for level %d",
				end,
				maxHeightForLevel,
			)
		}
	}
	if end < start {
		return 0, fmt.Errorf("invalid range: end %d was < start %d", end, start)
	}
	// The number of hashes is the difference between the start and end
	// requested heights, plus 1.
	return uint64(end-start) + 1, nil
}

// Figures out the actual opcode index we should move the machine to
// when we compute the history commitment. As there are different levels of challenge
// granularity, we have to do some math to figure out the correct index.
//
// Take, for example, that we have 5 challenge kinds:
//
// block    => over a range of L2 message hashes
// colossal => over a range of giant steps
// giant    => over a range of big steps
// big      => over a range of small steps
// small    => over a range of individual WAVM opcodes
//
// We only directly step through WAVM machines when in a subchallenge (starting at colossal),
// so we can ignore block challenges for this calculation.
//
// Next, take the following constants:
//
//	colossal_steps_per_giant_step = 16
//	giant_steps_per_big_step      = 8
//	big_steps_per_giant_step      = 4
//	small_steps_per_big_step      = 2
//
// Let's say we want to figure out the machine start opcode index for the following inputs:
//
//	colossal_step=4, giant_step=5, big_step=1, small_step=0
//
// We can compute the opcode index using the following algorithm for the example above.
//
//	  4 * (8 * 4 * 2)
//	+ 5 * (4 * 2)
//	+ 1 * (2)
//	+ 0
//	= 298
//
// This generalizes for any number of subchallenge levels into the algorithm below.
// It works by taking the sum of (each input * product of all challenge level height constants beneath its level).
// For the block challenge level, the machine start opcode index is always 0.
func (p *HistoryCommitmentProvider) computeMachineStartIndex(
	startHeights validatedStartHeights,
) OpcodeIndex {
	if len(startHeights) == 1 {
		return 0
	}
	// We ignore the block challenge level here.
	heights := startHeights[1:]
	leafHeights := p.challengeLeafHeights[1:]

	// Next, we compute the opcode index.
	opcodeIndex := uint64(0)
	idx := 1
	for _, height := range heights {
		total := uint64(1)
		for i := idx; i < len(leafHeights); i++ {
			total *= uint64(leafHeights[i])
		}
		opcodeIndex += total * uint64(height)
		idx += 1
	}
	return OpcodeIndex(opcodeIndex)
}

// Computes the step size for stepping through a machine at a block height. Each challenge level has a different
// amount of steps, so the total step size can be computed as a multiplication of all the
// next challenge levels needed.
func (p *HistoryCommitmentProvider) computeStepSize(startHeights validatedStartHeights) (StepSize, error) {
	challengeLevel := len(startHeights) - 1
	totalChallengeLevels := len(p.challengeLeafHeights)

	// The stepwise increment of the last challenge level is always one.
	if challengeLevel+1 == totalChallengeLevels {
		return 1, nil
	}
	// Otherwise, it is the multiplication of all the challenge leaf heights at the next
	// challenge levels.
	levels := p.challengeLeafHeights[challengeLevel+1:]
	total := uint64(1)
	for _, h := range levels {
		total *= uint64(h)
	}
	return StepSize(total), nil
}

// Validates a start heights input must be non-empty and have a max
// equal to the number of challenge levels.
func (p *HistoryCommitmentProvider) validateStartHeights(
	startHeights []Height,
) (validatedStartHeights, error) {
	if len(startHeights) == 0 {
		return nil, errors.New("must provide start heights to compute number of hashes")
	}
	challengeLevel := len(startHeights) - 1
	if challengeLevel >= len(p.challengeLeafHeights) {
		return nil, fmt.Errorf(
			"challenge level %d is out of range for challenge leaf heights %v",
			challengeLevel,
			p.challengeLeafHeights,
		)
	}
	return validatedStartHeights(startHeights), nil
}
