package execution

import (
	"encoding/binary"
	"errors"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var OutOfBoundsError = errors.New("instruction number out of bounds")

type Engine interface {
	// Gets the state root of the execution state.
	MachineHash() common.Hash
	BlockHash(blockNum uint64) common.Hash
	MachineIterator
	OneStepProver
}

type MachineIterator interface {
	IsStopped() bool
	// NumSteps returns the number of steps of execution to create that block.
	NumSteps() uint64
	// NextState gets the execution state after executing one instruction from an execution state.
	NextState() (*ExecutionState, error)
	// StateAfter returns the ExecutionState after executing num instructions.
	StateAfter(num uint64) (*ExecutionState, error)
}

// OneStepProver generates a one-step proof for executing one instruction from
// an execution state.
type OneStepProver interface {
	OneStepProof() ([]byte, error)
}

// SimulatedBlockGenerator generates the blocks that make up a chain.
type SimulatedBlockGenerator struct {
	stateRoots              []common.Hash
	maxInstructionsPerBlock uint64
}

// NewBlockGenerator(maxInstructionsPerBlock uint64) creates a block generator
// each block will use up to maxInstructionsPerBlock instructions of execution (randomly varying).
func NewBlockGenerator(maxInstructionsPerBlock uint64) *SimulatedBlockGenerator {
	return &SimulatedBlockGenerator{
		stateRoots:              []common.Hash{util.HashForUint(0)},
		maxInstructionsPerBlock: maxInstructionsPerBlock,
	}
}

func (gen *SimulatedBlockGenerator) BlockHash(blockNum uint64) common.Hash {
	for uint64(len(gen.stateRoots)) <= blockNum {
		gen.stateRoots = append(
			gen.stateRoots,
			crypto.Keccak256Hash(gen.stateRoots[len(gen.stateRoots)-1].Bytes()),
		)
	}
	return gen.stateRoots[blockNum]
}

func (gen *SimulatedBlockGenerator) NewExecutionEngine(blockNum uint64) (*SimulatedEngine, error) {
	if blockNum == 0 {
		return nil, errors.New("tried to make execution engine for genesis block")
	}
	startStateRoot := gen.BlockHash(blockNum - 1)
	endStateRoot := gen.BlockHash(blockNum)
	numSteps := binary.BigEndian.Uint64(crypto.Keccak256(startStateRoot.Bytes())[:8]) % (1 + gen.maxInstructionsPerBlock)
	return &SimulatedEngine{
		startStateRoot: startStateRoot,
		endStateRoot:   endStateRoot,
		numSteps:       numSteps,
	}, nil
}

type SimulatedEngine struct {
	startStateRoot common.Hash
	endStateRoot   common.Hash
	numSteps       uint64
}

func (engine *SimulatedEngine) serialize() []byte {
	ret := []byte{}
	ret = append(ret, engine.startStateRoot.Bytes()...)
	ret = append(ret, engine.endStateRoot.Bytes()...)
	ret = append(ret, binary.BigEndian.AppendUint64([]byte{}, engine.numSteps)...)
	return ret
}

func deserializeExecutionEngine(buf []byte) (*SimulatedEngine, error) {
	if len(buf) != 32+32+8 {
		return nil, errors.New("deserialization error")
	}
	return &SimulatedEngine{
		startStateRoot: common.BytesToHash(buf[:32]),
		endStateRoot:   common.BytesToHash(buf[32:64]),
		numSteps:       binary.BigEndian.Uint64(buf[64:]),
	}, nil
}

func (engine *SimulatedEngine) internalHash() common.Hash {
	return crypto.Keccak256Hash(engine.serialize())
}

type ExecutionState struct {
	engine  *SimulatedEngine
	stepNum uint64
}

func (engine *SimulatedEngine) NumSteps() uint64 {
	return engine.numSteps
}

func (engine *SimulatedEngine) StateAfter(num uint64) (*ExecutionState, error) {
	if num > engine.numSteps {
		return nil, OutOfBoundsError
	}
	return &ExecutionState{
		engine:  engine,
		stepNum: num,
	}, nil
}

func (execState *ExecutionState) IsStopped() bool {
	return execState.stepNum == execState.engine.numSteps
}

func (execState *ExecutionState) Hash() common.Hash {
	if execState.IsStopped() {
		return execState.engine.endStateRoot
	}
	return crypto.Keccak256Hash(binary.BigEndian.AppendUint64(execState.engine.internalHash().Bytes(), execState.stepNum))
}

func (execState *ExecutionState) NextState() (*ExecutionState, error) {
	if execState.IsStopped() {
		return nil, OutOfBoundsError
	}
	return &ExecutionState{
		engine:  execState.engine,
		stepNum: execState.stepNum + 1,
	}, nil
}

func (execState *ExecutionState) OneStepProof() ([]byte, error) {
	if execState.IsStopped() {
		return nil, OutOfBoundsError
	}
	ret := execState.engine.serialize()
	ret = append(ret, binary.BigEndian.AppendUint64([]byte{}, execState.stepNum)...)
	return ret, nil
}

// VerifyOneStepProof for a machine pre and post states.
func VerifyOneStepProof(beforeStateRoot common.Hash, claimedAfterStateRoot common.Hash, proof []byte) bool {
	if len(proof) < 8 {
		return false
	}
	engine, err := deserializeExecutionEngine(proof[:len(proof)-8])
	if err != nil {
		return false
	}
	beforeState := ExecutionState{
		engine:  engine,
		stepNum: binary.BigEndian.Uint64(proof[len(proof)-8:]),
	}
	if beforeState.Hash() != beforeStateRoot {
		return false
	}
	afterState, err := beforeState.NextState()
	if err != nil {
		return false
	}
	return afterState.Hash() == claimedAfterStateRoot
}
