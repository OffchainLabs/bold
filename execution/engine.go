package execution

import (
	"errors"
	"math/big"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrOutOfBounds = errors.New("instruction number out of bounds")
)

type Machine interface {
	CurrentStepNum() uint64
	Hash() common.Hash
	IsStopped() bool
	Clone() Machine
	Step(steps uint64) error
	OneStepProof() ([]byte, error)
}

type SimpleMachine struct {
	step  uint64
	state *big.Int
}

func NewSimpleMachine(startingState *big.Int) *SimpleMachine {
	return &SimpleMachine{
		step:  0,
		state: new(big.Int).Set(startingState),
	}
}

func (m *SimpleMachine) CurrentStepNum() uint64 {
	return m.step
}

func (m *SimpleMachine) stateBytes() []byte {
	return math.U256Bytes(m.state)
}

func (m *SimpleMachine) Hash() common.Hash {
	var blockHash common.Hash
	if m.state.Sign() > 0 {
		blockHash = crypto.Keccak256Hash(m.stateBytes())
	}
	return protocol.GoGlobalState{
		BlockHash: blockHash,
	}.Hash()
}

func (m *SimpleMachine) IsStopped() bool {
	return m.step > 0 && m.Hash()[0] == 0
}

func (m *SimpleMachine) Clone() Machine {
	newMachine := *m
	newMachine.state = new(big.Int).Set(m.state)
	return &newMachine
}

func (m *SimpleMachine) Step(steps uint64) error {
	for ; steps > 0; steps-- {
		if m.IsStopped() {
			m.step += steps
			return nil
		}
		m.step++
		m.state.Add(m.state, common.Big1)
	}
	return nil
}

func (m *SimpleMachine) OneStepProof() ([]byte, error) {
	return m.stateBytes(), nil
}

func (m *SimpleMachine) GetState() *big.Int {
	return new(big.Int).Set(m.state)
}

// VerifySimpleMachineOneStepProof checks the claimed post-state root results from executing
// a specified pre-state hash.
func VerifySimpleMachineOneStepProof(beforeStateRoot common.Hash, claimedAfterStateRoot common.Hash, step uint64, proof []byte) bool {
	if len(proof) != 32 {
		return false
	}
	state := new(big.Int).SetBytes(proof)
	mach := NewSimpleMachine(state)
	mach.step = step
	if mach.Hash() != beforeStateRoot {
		return false
	}
	err := mach.Step(1)
	if err != nil {
		return false
	}
	return mach.Hash() == claimedAfterStateRoot
}
