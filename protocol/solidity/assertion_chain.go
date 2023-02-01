package solidity

import (
	"github.com/OffchainLabs/challenge-protocol-v2/protocol/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// TODO: Implement a nice wrapper around solgen bindings to the protocol contracts
// that aligns with the AssertionManager interface. This aims to provide a clean API
// that does not expose Ethereum internals to consumers.
// TODO: Implement custom, nice wrapper types such as Assertion, ChallengeVertex, Challenge
// which do not expose Ethereum internals to callers outside of this package.
type AssertionChain struct {
	backend bind.ContractBackend
	chain   *bindings.IAssertionChain
}

func NewAssertionChain() (*AssertionChain, error) {
	chain, err := bindings.NewIAssertionChain(common.Address{}, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionChain{
		chain: chain,
	}, nil
}

func (ac *AssertionChain) ChallengePeriodLength() (uint64, error) {
	res, err := ac.chain.IAssertionChainCaller.ChallengePeriodSeconds(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}
	return res.Uint64(), nil
}

func (ac *AssertionChain) NumAssertions() (uint64, error) {
	res, err := ac.chain.IAssertionChainCaller.NumAssertions(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}
	return res.Uint64(), nil
}

// // AssertionManager allows the creation of new leaves for a Staker with a State Commitment
// // and a previous assertion.
// type AssertionManager interface {
// 	Inbox() *Inbox
// 	NumAssertions(tx *ActiveTx) uint64
// 	AssertionBySequenceNum(tx *ActiveTx, seqNum AssertionSequenceNumber) (*Assertion, error)
// 	ChallengeByCommitHash(tx *ActiveTx, commitHash ChallengeCommitHash) (*Challenge, error)
// 	ChallengeVertexByCommitHash(tx *ActiveTx, challenge ChallengeCommitHash, vertex VertexCommitHash) (*ChallengeVertex, error)
// 	IsAtOneStepFork(
// 		tx *ActiveTx,
// 		challengeCommitHash ChallengeCommitHash,
// 		vertexCommit util.HistoryCommitment,
// 		vertexParentCommit util.HistoryCommitment,
// 	) (bool, error)
// 	ChallengePeriodLength(tx *ActiveTx) time.Duration
// 	LatestConfirmed(*ActiveTx) *Assertion
// 	CreateLeaf(tx *ActiveTx, prev *Assertion, commitment StateCommitment, staker common.Address) (*Assertion, error)
// 	TimeReference() util.TimeReference
// }
