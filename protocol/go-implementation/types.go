package goimpl

import "github.com/OffchainLabs/challenge-protocol-v2/protocol"
import "github.com/ethereum/go-ethereum/common"

func (a *Assertion) Height() uint64 {
	return a.StateCommitment.Height
}

func (a *Assertion) SeqNum() protocol.AssertionSequenceNumber {
	return protocol.AssertionSequenceNumber(a.SequenceNum)
}

func (a *Assertion) PrevSeqNum() protocol.AssertionSequenceNumber {
	return protocol.AssertionSequenceNumber(a.Prev.Unwrap().SequenceNum)
}

func (a *Assertion) StateHash() common.Hash {
	return a.StateCommitment.StateRoot
}
