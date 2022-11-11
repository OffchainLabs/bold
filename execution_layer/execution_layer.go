package execution_layer

import (
	"bytes"
	"encoding/binary"
	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
)

type ExecutionState interface {
	Root() common.Hash
	MessagesConsumed() uint64

	Execute() (ExecutionState, error)

	Prove(message []byte, afterRoot common.Hash) ([]byte, error)
	GetProofChecker() ProofChecker

	Serialize(io.Writer) error
}

type ProofChecker func(beforeRoot common.Hash, afterRoot common.Hash, proof []byte) bool

type executionStateImpl struct {
	chain        *protocol.AssertionChain
	vmRoot       common.Hash
	msgsConsumed uint64
}

func GenesisExecutionState(chain *protocol.AssertionChain) ExecutionState {
	return &executionStateImpl{
		chain:        chain,
		vmRoot:       common.Hash{},
		msgsConsumed: 0,
	}
}

func (state *executionStateImpl) Root() common.Hash {
	return crypto.Keccak256Hash(binary.BigEndian.AppendUint64(state.vmRoot.Bytes(), state.msgsConsumed))
}

func (state *executionStateImpl) MessagesConsumed() uint64 {
	return state.msgsConsumed
}

func (state *executionStateImpl) Execute() (ExecutionState, error) {
	var nextMsg []byte
	err := state.chain.Call(func(tx *protocol.ActiveTx, innerChain *protocol.AssertionChain) error {
		msg, err2 := innerChain.Inbox().GetMessage(tx, state.msgsConsumed)
		if err2 != nil {
			return err2
		}
		nextMsg = msg
		return nil
	})
	if err != nil {
		return nil, err
	}
	return state.executeMessage(nextMsg), nil
}

func (state *executionStateImpl) executeMessage(msg []byte) ExecutionState {
	return &executionStateImpl{
		chain:        state.chain,
		vmRoot:       crypto.Keccak256Hash(state.vmRoot.Bytes(), msg),
		msgsConsumed: state.msgsConsumed + 1,
	}
}

func (state *executionStateImpl) Prove(message []byte, afterRoot common.Hash) ([]byte, error) {
	afterState := state.executeMessage(message)
	if afterState.Root() != afterRoot {
		return nil, protocol.ErrWrongState
	}
	buf := bytes.Buffer{}
	if err := state.Serialize(&buf); err != nil {
		return nil, err
	}
	return append(buf.Bytes(), message...), nil
}

func (state *executionStateImpl) Serialize(wr io.Writer) error {
	_, err := wr.Write(binary.BigEndian.AppendUint64(state.vmRoot.Bytes(), state.msgsConsumed))
	return err
}

func deserializeStateImpl(chain *protocol.AssertionChain, rd io.Reader) (ExecutionState, error) {
	var buf [40]byte
	if _, err := io.ReadFull(rd, buf[:]); err != nil {
		return nil, err
	}
	return &executionStateImpl{
		chain:        chain,
		vmRoot:       common.BytesToHash(buf[:32]),
		msgsConsumed: binary.BigEndian.Uint64(buf[32:]),
	}, nil
}

func (state *executionStateImpl) GetProofChecker() ProofChecker {
	return func(beforeRoot common.Hash, afterRoot common.Hash, proof []byte) bool {
		rd := bytes.NewReader(proof)
		beforeState, err := deserializeStateImpl(state.chain, rd)
		if err != nil {
			return false
		}
		msg, err := io.ReadAll(rd)
		if err != nil {
			return false
		}
		afterState := beforeState.(*executionStateImpl).executeMessage(msg)
		return afterState.Root() == afterRoot
	}
}
