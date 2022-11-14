package execution_layer

import (
	"bytes"
	"encoding/binary"
	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"sync"
)

type ExecutionState interface {
	Root() common.Hash
	MessagesConsumed() uint64

	ExecuteMessage(msg []byte) ExecutionState
	ExecuteNextChainMessage() (ExecutionState, error)

	Prove(message []byte, afterRoot common.Hash) ([]byte, error)
	GetProofChecker() ProofChecker

	Clone() ExecutionState
	Serialize(io.Writer) error
}

type ProofChecker func(beforeRoot common.Hash, afterRoot common.Hash, msgHash common.Hash, proof []byte) bool

type executionStateImpl struct {
	mutex           sync.Mutex
	chain           *protocol.AssertionChain
	vmRoot          common.Hash
	numMsgsConsumed uint64
}

func GenesisExecutionState(chain *protocol.AssertionChain) ExecutionState {
	return &executionStateImpl{
		mutex:           sync.Mutex{},
		chain:           chain,
		vmRoot:          common.Hash{},
		numMsgsConsumed: 0,
	}
}

func (state *executionStateImpl) Root() common.Hash {
	state.mutex.Lock()
	defer state.mutex.Unlock()
	return crypto.Keccak256Hash(binary.BigEndian.AppendUint64(state.vmRoot.Bytes(), state.numMsgsConsumed))
}

func (state *executionStateImpl) MessagesConsumed() uint64 {
	state.mutex.Lock()
	defer state.mutex.Unlock()
	return state.numMsgsConsumed
}

func (state *executionStateImpl) ExecuteNextChainMessage() (ExecutionState, error) {
	state.mutex.Lock()
	defer state.mutex.Unlock()
	var nextMsg []byte
	err := state.chain.Call(func(tx *protocol.ActiveTx, innerChain *protocol.AssertionChain) error {
		msg, err2 := innerChain.Inbox().GetMessage(tx, state.numMsgsConsumed)
		if err2 != nil {
			return err2
		}
		nextMsg = msg
		return nil
	})
	if err != nil {
		return nil, err
	}
	return state.executeMessageLocked(nextMsg), nil
}

func (state *executionStateImpl) ExecuteMessage(msg []byte) ExecutionState {
	state.mutex.Lock()
	defer state.mutex.Unlock()
	return state.executeMessageLocked(msg)
}

func (state *executionStateImpl) executeMessageLocked(msg []byte) ExecutionState {
	return &executionStateImpl{
		chain:           state.chain,
		vmRoot:          crypto.Keccak256Hash(state.vmRoot.Bytes(), msg),
		numMsgsConsumed: state.numMsgsConsumed + 1,
	}
}

func (state *executionStateImpl) Prove(message []byte, afterRoot common.Hash) ([]byte, error) {
	state.mutex.Lock()
	defer state.mutex.Unlock()
	afterState := state.executeMessageLocked(message)
	if afterState.Root() != afterRoot {
		return nil, protocol.ErrWrongState
	}
	buf := bytes.Buffer{}
	if err := state.serializeLocked(&buf); err != nil {
		return nil, err
	}
	return append(buf.Bytes(), message...), nil
}

func (state *executionStateImpl) Serialize(wr io.Writer) error {
	state.mutex.Lock()
	defer state.mutex.Unlock()
	return state.serializeLocked(wr)
}

func (state *executionStateImpl) serializeLocked(wr io.Writer) error {
	_, err := wr.Write(binary.BigEndian.AppendUint64(state.vmRoot.Bytes(), state.numMsgsConsumed))
	return err
}

func deserializeStateImpl(chain *protocol.AssertionChain, rd io.Reader) (ExecutionState, error) {
	var buf [40]byte
	if _, err := io.ReadFull(rd, buf[:]); err != nil {
		return nil, err
	}
	return &executionStateImpl{
		chain:           chain,
		vmRoot:          common.BytesToHash(buf[:32]),
		numMsgsConsumed: binary.BigEndian.Uint64(buf[32:]),
	}, nil
}

func (state *executionStateImpl) GetProofChecker() ProofChecker {
	return func(beforeRoot common.Hash, afterRoot common.Hash, msgHash common.Hash, proof []byte) bool {
		rd := bytes.NewReader(proof)
		beforeState, err := deserializeStateImpl(state.chain, rd)
		if err != nil {
			return false
		}
		msg, err := io.ReadAll(rd)
		if err != nil {
			return false
		}
		if crypto.Keccak256Hash(msg) != msgHash {
			return false
		}
		afterState := beforeState.(*executionStateImpl).executeMessageLocked(msg) // don't need locking because private instance
		return afterState.Root() == afterRoot
	}
}

func (state *executionStateImpl) Clone() ExecutionState {
	state.mutex.Lock()
	defer state.mutex.Unlock()
	return &executionStateImpl{
		mutex:           sync.Mutex{},
		chain:           state.chain,
		vmRoot:          state.vmRoot,
		numMsgsConsumed: state.numMsgsConsumed,
	}
}
