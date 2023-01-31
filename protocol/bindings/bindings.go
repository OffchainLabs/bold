// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IAssertionChainAssertion is an auto generated low-level Go binding around an user-defined struct.
type IAssertionChainAssertion struct {
	SeqNum                       *big.Int
	StateCommitment              IAssertionChainStateCommitment
	Status                       *big.Int
	IsFirstChild                 bool
	FirstChildCreationTimestamp  *big.Int
	SecondChildCreationTimestamp *big.Int
	Actor                        common.Address
}

// IAssertionChainChallenge is an auto generated low-level Go binding around an user-defined struct.
type IAssertionChainChallenge struct {
	SeqNum            *big.Int
	NextSeqNum        *big.Int
	Root              IAssertionChainChallengeVertex
	LatestConfirmed   IAssertionChainChallengeVertex
	CreationTimestamp *big.Int
	Actor             common.Address
}

// IAssertionChainChallengeVertex is an auto generated low-level Go binding around an user-defined struct.
type IAssertionChainChallengeVertex struct {
	SeqNum                         *big.Int
	ChallengeParentStateCommitHash [32]byte
	Actor                          bool
	IsLeaf                         bool
	PsTimer                        *big.Int
}

// IAssertionChainHistoryCommitment is an auto generated low-level Go binding around an user-defined struct.
type IAssertionChainHistoryCommitment struct {
	Height     *big.Int
	MerkleRoot [32]byte
}

// IAssertionChainStateCommitment is an auto generated low-level Go binding around an user-defined struct.
type IAssertionChainStateCommitment struct {
	Height    *big.Int
	StateRoot [32]byte
}

// AssertionChainMetaData contains all meta data concerning the AssertionChain contract.
var AssertionChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.StateCommitment\",\"name\":\"stateCommitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFirstChild\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"firstChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Assertion\",\"name\":\"assertion\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.HistoryCommitment\",\"name\":\"history\",\"type\":\"tuple\"}],\"name\":\"addChallengeVertex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"vertex\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"vertex\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.HistoryCommitment\",\"name\":\"history\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"bisect\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"bisectedVertex\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextSeqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"root\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"latestConfirmed\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"creationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Challenge\",\"name\":\"challenge\",\"type\":\"tuple\"}],\"name\":\"challengeCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengePeriodSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextSeqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"root\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"latestConfirmed\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"creationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Challenge\",\"name\":\"challenge\",\"type\":\"tuple\"}],\"name\":\"challengeWinner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.StateCommitment\",\"name\":\"stateCommitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFirstChild\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"firstChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Assertion\",\"name\":\"assertion\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"vertex\",\"type\":\"tuple\"}],\"name\":\"confirmForChallengeDeadline\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"vertex\",\"type\":\"tuple\"}],\"name\":\"confirmForPSTimer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"vertex\",\"type\":\"tuple\"}],\"name\":\"confirmForSubchallengeWin\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.StateCommitment\",\"name\":\"stateCommitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFirstChild\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"firstChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Assertion\",\"name\":\"assertion\",\"type\":\"tuple\"}],\"name\":\"confirmForWin\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.StateCommitment\",\"name\":\"stateCommitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFirstChild\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"firstChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Assertion\",\"name\":\"assertion\",\"type\":\"tuple\"}],\"name\":\"confirmNoRival\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.StateCommitment\",\"name\":\"stateCommitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFirstChild\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"firstChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Assertion\",\"name\":\"prev\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.StateCommitment\",\"name\":\"commit\",\"type\":\"tuple\"}],\"name\":\"createAssertion\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.StateCommitment\",\"name\":\"stateCommitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFirstChild\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"firstChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Assertion\",\"name\":\"assertion\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.StateCommitment\",\"name\":\"stateCommitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFirstChild\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"firstChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Assertion\",\"name\":\"prev\",\"type\":\"tuple\"}],\"name\":\"createChallenge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextSeqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"root\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"latestConfirmed\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"creationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Challenge\",\"name\":\"challenge\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"vertex\",\"type\":\"tuple\"}],\"name\":\"eligibleForNewSuccessor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"}],\"name\":\"getAssertion\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.StateCommitment\",\"name\":\"stateCommitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFirstChild\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"firstChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Assertion\",\"name\":\"assertion\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentStateCommitHash\",\"type\":\"bytes32\"}],\"name\":\"getChallenge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextSeqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"root\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"latestConfirmed\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"creationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Challenge\",\"name\":\"challenge\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"parentStateCommitHash\",\"type\":\"bytes32\"}],\"name\":\"getChallengeVertex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"vertex\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"vertex\",\"type\":\"tuple\"}],\"name\":\"isPresumptiveSuccessor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmedAssertion\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.StateCommitment\",\"name\":\"stateCommitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFirstChild\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"firstChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Assertion\",\"name\":\"assertion\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"mergingFrom\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"mergingTo\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"merge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challengeParentStateCommitHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"actor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLeaf\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"psTimer\",\"type\":\"uint256\"}],\"internalType\":\"structIAssertionChain.ChallengeVertex\",\"name\":\"mergedToVertex\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numAssertions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.StateCommitment\",\"name\":\"stateCommitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFirstChild\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"firstChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Assertion\",\"name\":\"assertion\",\"type\":\"tuple\"}],\"name\":\"rejectForLoss\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssertionChain.StateCommitment\",\"name\":\"stateCommitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFirstChild\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"firstChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondChildCreationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"}],\"internalType\":\"structIAssertionChain.Assertion\",\"name\":\"assertion\",\"type\":\"tuple\"}],\"name\":\"rejectForPrev\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// AssertionChainABI is the input ABI used to generate the binding from.
// Deprecated: Use AssertionChainMetaData.ABI instead.
var AssertionChainABI = AssertionChainMetaData.ABI

// AssertionChain is an auto generated Go binding around an Ethereum contract.
type AssertionChain struct {
	AssertionChainCaller     // Read-only binding to the contract
	AssertionChainTransactor // Write-only binding to the contract
	AssertionChainFilterer   // Log filterer for contract events
}

// AssertionChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssertionChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssertionChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssertionChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssertionChainSession struct {
	Contract     *AssertionChain   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssertionChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssertionChainCallerSession struct {
	Contract *AssertionChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AssertionChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssertionChainTransactorSession struct {
	Contract     *AssertionChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AssertionChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssertionChainRaw struct {
	Contract *AssertionChain // Generic contract binding to access the raw methods on
}

// AssertionChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssertionChainCallerRaw struct {
	Contract *AssertionChainCaller // Generic read-only contract binding to access the raw methods on
}

// AssertionChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssertionChainTransactorRaw struct {
	Contract *AssertionChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssertionChain creates a new instance of AssertionChain, bound to a specific deployed contract.
func NewAssertionChain(address common.Address, backend bind.ContractBackend) (*AssertionChain, error) {
	contract, err := bindAssertionChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssertionChain{AssertionChainCaller: AssertionChainCaller{contract: contract}, AssertionChainTransactor: AssertionChainTransactor{contract: contract}, AssertionChainFilterer: AssertionChainFilterer{contract: contract}}, nil
}

// NewAssertionChainCaller creates a new read-only instance of AssertionChain, bound to a specific deployed contract.
func NewAssertionChainCaller(address common.Address, caller bind.ContractCaller) (*AssertionChainCaller, error) {
	contract, err := bindAssertionChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionChainCaller{contract: contract}, nil
}

// NewAssertionChainTransactor creates a new write-only instance of AssertionChain, bound to a specific deployed contract.
func NewAssertionChainTransactor(address common.Address, transactor bind.ContractTransactor) (*AssertionChainTransactor, error) {
	contract, err := bindAssertionChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionChainTransactor{contract: contract}, nil
}

// NewAssertionChainFilterer creates a new log filterer instance of AssertionChain, bound to a specific deployed contract.
func NewAssertionChainFilterer(address common.Address, filterer bind.ContractFilterer) (*AssertionChainFilterer, error) {
	contract, err := bindAssertionChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssertionChainFilterer{contract: contract}, nil
}

// bindAssertionChain binds a generic wrapper to an already deployed contract.
func bindAssertionChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssertionChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssertionChain *AssertionChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssertionChain.Contract.AssertionChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssertionChain *AssertionChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionChain.Contract.AssertionChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssertionChain *AssertionChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssertionChain.Contract.AssertionChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssertionChain *AssertionChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssertionChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssertionChain *AssertionChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssertionChain *AssertionChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssertionChain.Contract.contract.Transact(opts, method, params...)
}

// ChallengePeriodSeconds is a free data retrieval call binding the contract method 0xfb601294.
//
// Solidity: function challengePeriodSeconds() view returns(uint256)
func (_AssertionChain *AssertionChainCaller) ChallengePeriodSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "challengePeriodSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengePeriodSeconds is a free data retrieval call binding the contract method 0xfb601294.
//
// Solidity: function challengePeriodSeconds() view returns(uint256)
func (_AssertionChain *AssertionChainSession) ChallengePeriodSeconds() (*big.Int, error) {
	return _AssertionChain.Contract.ChallengePeriodSeconds(&_AssertionChain.CallOpts)
}

// ChallengePeriodSeconds is a free data retrieval call binding the contract method 0xfb601294.
//
// Solidity: function challengePeriodSeconds() view returns(uint256)
func (_AssertionChain *AssertionChainCallerSession) ChallengePeriodSeconds() (*big.Int, error) {
	return _AssertionChain.Contract.ChallengePeriodSeconds(&_AssertionChain.CallOpts)
}

// GetAssertion is a free data retrieval call binding the contract method 0x1d99e167.
//
// Solidity: function getAssertion(uint256 seqNum) view returns((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion)
func (_AssertionChain *AssertionChainCaller) GetAssertion(opts *bind.CallOpts, seqNum *big.Int) (IAssertionChainAssertion, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "getAssertion", seqNum)

	if err != nil {
		return *new(IAssertionChainAssertion), err
	}

	out0 := *abi.ConvertType(out[0], new(IAssertionChainAssertion)).(*IAssertionChainAssertion)

	return out0, err

}

// GetAssertion is a free data retrieval call binding the contract method 0x1d99e167.
//
// Solidity: function getAssertion(uint256 seqNum) view returns((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion)
func (_AssertionChain *AssertionChainSession) GetAssertion(seqNum *big.Int) (IAssertionChainAssertion, error) {
	return _AssertionChain.Contract.GetAssertion(&_AssertionChain.CallOpts, seqNum)
}

// GetAssertion is a free data retrieval call binding the contract method 0x1d99e167.
//
// Solidity: function getAssertion(uint256 seqNum) view returns((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion)
func (_AssertionChain *AssertionChainCallerSession) GetAssertion(seqNum *big.Int) (IAssertionChainAssertion, error) {
	return _AssertionChain.Contract.GetAssertion(&_AssertionChain.CallOpts, seqNum)
}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(bytes32 parentStateCommitHash) view returns((uint256,uint256,(uint256,bytes32,bool,bool,uint256),(uint256,bytes32,bool,bool,uint256),uint256,address) challenge)
func (_AssertionChain *AssertionChainCaller) GetChallenge(opts *bind.CallOpts, parentStateCommitHash [32]byte) (IAssertionChainChallenge, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "getChallenge", parentStateCommitHash)

	if err != nil {
		return *new(IAssertionChainChallenge), err
	}

	out0 := *abi.ConvertType(out[0], new(IAssertionChainChallenge)).(*IAssertionChainChallenge)

	return out0, err

}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(bytes32 parentStateCommitHash) view returns((uint256,uint256,(uint256,bytes32,bool,bool,uint256),(uint256,bytes32,bool,bool,uint256),uint256,address) challenge)
func (_AssertionChain *AssertionChainSession) GetChallenge(parentStateCommitHash [32]byte) (IAssertionChainChallenge, error) {
	return _AssertionChain.Contract.GetChallenge(&_AssertionChain.CallOpts, parentStateCommitHash)
}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(bytes32 parentStateCommitHash) view returns((uint256,uint256,(uint256,bytes32,bool,bool,uint256),(uint256,bytes32,bool,bool,uint256),uint256,address) challenge)
func (_AssertionChain *AssertionChainCallerSession) GetChallenge(parentStateCommitHash [32]byte) (IAssertionChainChallenge, error) {
	return _AssertionChain.Contract.GetChallenge(&_AssertionChain.CallOpts, parentStateCommitHash)
}

// GetChallengeVertex is a free data retrieval call binding the contract method 0x601e00f2.
//
// Solidity: function getChallengeVertex(uint256 seqNum, bytes32 parentStateCommitHash) view returns((uint256,bytes32,bool,bool,uint256) vertex)
func (_AssertionChain *AssertionChainCaller) GetChallengeVertex(opts *bind.CallOpts, seqNum *big.Int, parentStateCommitHash [32]byte) (IAssertionChainChallengeVertex, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "getChallengeVertex", seqNum, parentStateCommitHash)

	if err != nil {
		return *new(IAssertionChainChallengeVertex), err
	}

	out0 := *abi.ConvertType(out[0], new(IAssertionChainChallengeVertex)).(*IAssertionChainChallengeVertex)

	return out0, err

}

// GetChallengeVertex is a free data retrieval call binding the contract method 0x601e00f2.
//
// Solidity: function getChallengeVertex(uint256 seqNum, bytes32 parentStateCommitHash) view returns((uint256,bytes32,bool,bool,uint256) vertex)
func (_AssertionChain *AssertionChainSession) GetChallengeVertex(seqNum *big.Int, parentStateCommitHash [32]byte) (IAssertionChainChallengeVertex, error) {
	return _AssertionChain.Contract.GetChallengeVertex(&_AssertionChain.CallOpts, seqNum, parentStateCommitHash)
}

// GetChallengeVertex is a free data retrieval call binding the contract method 0x601e00f2.
//
// Solidity: function getChallengeVertex(uint256 seqNum, bytes32 parentStateCommitHash) view returns((uint256,bytes32,bool,bool,uint256) vertex)
func (_AssertionChain *AssertionChainCallerSession) GetChallengeVertex(seqNum *big.Int, parentStateCommitHash [32]byte) (IAssertionChainChallengeVertex, error) {
	return _AssertionChain.Contract.GetChallengeVertex(&_AssertionChain.CallOpts, seqNum, parentStateCommitHash)
}

// LatestConfirmedAssertion is a free data retrieval call binding the contract method 0xb2e6ee2c.
//
// Solidity: function latestConfirmedAssertion() view returns((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion)
func (_AssertionChain *AssertionChainCaller) LatestConfirmedAssertion(opts *bind.CallOpts) (IAssertionChainAssertion, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "latestConfirmedAssertion")

	if err != nil {
		return *new(IAssertionChainAssertion), err
	}

	out0 := *abi.ConvertType(out[0], new(IAssertionChainAssertion)).(*IAssertionChainAssertion)

	return out0, err

}

// LatestConfirmedAssertion is a free data retrieval call binding the contract method 0xb2e6ee2c.
//
// Solidity: function latestConfirmedAssertion() view returns((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion)
func (_AssertionChain *AssertionChainSession) LatestConfirmedAssertion() (IAssertionChainAssertion, error) {
	return _AssertionChain.Contract.LatestConfirmedAssertion(&_AssertionChain.CallOpts)
}

// LatestConfirmedAssertion is a free data retrieval call binding the contract method 0xb2e6ee2c.
//
// Solidity: function latestConfirmedAssertion() view returns((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion)
func (_AssertionChain *AssertionChainCallerSession) LatestConfirmedAssertion() (IAssertionChainAssertion, error) {
	return _AssertionChain.Contract.LatestConfirmedAssertion(&_AssertionChain.CallOpts)
}

// NumAssertions is a free data retrieval call binding the contract method 0x8200394b.
//
// Solidity: function numAssertions() view returns(uint256)
func (_AssertionChain *AssertionChainCaller) NumAssertions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "numAssertions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumAssertions is a free data retrieval call binding the contract method 0x8200394b.
//
// Solidity: function numAssertions() view returns(uint256)
func (_AssertionChain *AssertionChainSession) NumAssertions() (*big.Int, error) {
	return _AssertionChain.Contract.NumAssertions(&_AssertionChain.CallOpts)
}

// NumAssertions is a free data retrieval call binding the contract method 0x8200394b.
//
// Solidity: function numAssertions() view returns(uint256)
func (_AssertionChain *AssertionChainCallerSession) NumAssertions() (*big.Int, error) {
	return _AssertionChain.Contract.NumAssertions(&_AssertionChain.CallOpts)
}

// AddChallengeVertex is a paid mutator transaction binding the contract method 0x95cc44f0.
//
// Solidity: function addChallengeVertex((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion, (uint256,bytes32) history) payable returns((uint256,bytes32,bool,bool,uint256) vertex)
func (_AssertionChain *AssertionChainTransactor) AddChallengeVertex(opts *bind.TransactOpts, assertion IAssertionChainAssertion, history IAssertionChainHistoryCommitment) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "addChallengeVertex", assertion, history)
}

// AddChallengeVertex is a paid mutator transaction binding the contract method 0x95cc44f0.
//
// Solidity: function addChallengeVertex((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion, (uint256,bytes32) history) payable returns((uint256,bytes32,bool,bool,uint256) vertex)
func (_AssertionChain *AssertionChainSession) AddChallengeVertex(assertion IAssertionChainAssertion, history IAssertionChainHistoryCommitment) (*types.Transaction, error) {
	return _AssertionChain.Contract.AddChallengeVertex(&_AssertionChain.TransactOpts, assertion, history)
}

// AddChallengeVertex is a paid mutator transaction binding the contract method 0x95cc44f0.
//
// Solidity: function addChallengeVertex((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion, (uint256,bytes32) history) payable returns((uint256,bytes32,bool,bool,uint256) vertex)
func (_AssertionChain *AssertionChainTransactorSession) AddChallengeVertex(assertion IAssertionChainAssertion, history IAssertionChainHistoryCommitment) (*types.Transaction, error) {
	return _AssertionChain.Contract.AddChallengeVertex(&_AssertionChain.TransactOpts, assertion, history)
}

// Bisect is a paid mutator transaction binding the contract method 0x05eb85d8.
//
// Solidity: function bisect((uint256,bytes32,bool,bool,uint256) vertex, (uint256,bytes32) history, bytes32[] proof) payable returns((uint256,bytes32,bool,bool,uint256) bisectedVertex)
func (_AssertionChain *AssertionChainTransactor) Bisect(opts *bind.TransactOpts, vertex IAssertionChainChallengeVertex, history IAssertionChainHistoryCommitment, proof [][32]byte) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "bisect", vertex, history, proof)
}

// Bisect is a paid mutator transaction binding the contract method 0x05eb85d8.
//
// Solidity: function bisect((uint256,bytes32,bool,bool,uint256) vertex, (uint256,bytes32) history, bytes32[] proof) payable returns((uint256,bytes32,bool,bool,uint256) bisectedVertex)
func (_AssertionChain *AssertionChainSession) Bisect(vertex IAssertionChainChallengeVertex, history IAssertionChainHistoryCommitment, proof [][32]byte) (*types.Transaction, error) {
	return _AssertionChain.Contract.Bisect(&_AssertionChain.TransactOpts, vertex, history, proof)
}

// Bisect is a paid mutator transaction binding the contract method 0x05eb85d8.
//
// Solidity: function bisect((uint256,bytes32,bool,bool,uint256) vertex, (uint256,bytes32) history, bytes32[] proof) payable returns((uint256,bytes32,bool,bool,uint256) bisectedVertex)
func (_AssertionChain *AssertionChainTransactorSession) Bisect(vertex IAssertionChainChallengeVertex, history IAssertionChainHistoryCommitment, proof [][32]byte) (*types.Transaction, error) {
	return _AssertionChain.Contract.Bisect(&_AssertionChain.TransactOpts, vertex, history, proof)
}

// ChallengeCompleted is a paid mutator transaction binding the contract method 0x6845c13f.
//
// Solidity: function challengeCompleted((uint256,uint256,(uint256,bytes32,bool,bool,uint256),(uint256,bytes32,bool,bool,uint256),uint256,address) challenge) returns(bool)
func (_AssertionChain *AssertionChainTransactor) ChallengeCompleted(opts *bind.TransactOpts, challenge IAssertionChainChallenge) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "challengeCompleted", challenge)
}

// ChallengeCompleted is a paid mutator transaction binding the contract method 0x6845c13f.
//
// Solidity: function challengeCompleted((uint256,uint256,(uint256,bytes32,bool,bool,uint256),(uint256,bytes32,bool,bool,uint256),uint256,address) challenge) returns(bool)
func (_AssertionChain *AssertionChainSession) ChallengeCompleted(challenge IAssertionChainChallenge) (*types.Transaction, error) {
	return _AssertionChain.Contract.ChallengeCompleted(&_AssertionChain.TransactOpts, challenge)
}

// ChallengeCompleted is a paid mutator transaction binding the contract method 0x6845c13f.
//
// Solidity: function challengeCompleted((uint256,uint256,(uint256,bytes32,bool,bool,uint256),(uint256,bytes32,bool,bool,uint256),uint256,address) challenge) returns(bool)
func (_AssertionChain *AssertionChainTransactorSession) ChallengeCompleted(challenge IAssertionChainChallenge) (*types.Transaction, error) {
	return _AssertionChain.Contract.ChallengeCompleted(&_AssertionChain.TransactOpts, challenge)
}

// ChallengeWinner is a paid mutator transaction binding the contract method 0x08fb4b75.
//
// Solidity: function challengeWinner((uint256,uint256,(uint256,bytes32,bool,bool,uint256),(uint256,bytes32,bool,bool,uint256),uint256,address) challenge) returns((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion)
func (_AssertionChain *AssertionChainTransactor) ChallengeWinner(opts *bind.TransactOpts, challenge IAssertionChainChallenge) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "challengeWinner", challenge)
}

// ChallengeWinner is a paid mutator transaction binding the contract method 0x08fb4b75.
//
// Solidity: function challengeWinner((uint256,uint256,(uint256,bytes32,bool,bool,uint256),(uint256,bytes32,bool,bool,uint256),uint256,address) challenge) returns((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion)
func (_AssertionChain *AssertionChainSession) ChallengeWinner(challenge IAssertionChainChallenge) (*types.Transaction, error) {
	return _AssertionChain.Contract.ChallengeWinner(&_AssertionChain.TransactOpts, challenge)
}

// ChallengeWinner is a paid mutator transaction binding the contract method 0x08fb4b75.
//
// Solidity: function challengeWinner((uint256,uint256,(uint256,bytes32,bool,bool,uint256),(uint256,bytes32,bool,bool,uint256),uint256,address) challenge) returns((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion)
func (_AssertionChain *AssertionChainTransactorSession) ChallengeWinner(challenge IAssertionChainChallenge) (*types.Transaction, error) {
	return _AssertionChain.Contract.ChallengeWinner(&_AssertionChain.TransactOpts, challenge)
}

// ConfirmForChallengeDeadline is a paid mutator transaction binding the contract method 0x542457ba.
//
// Solidity: function confirmForChallengeDeadline((uint256,bytes32,bool,bool,uint256) vertex) payable returns()
func (_AssertionChain *AssertionChainTransactor) ConfirmForChallengeDeadline(opts *bind.TransactOpts, vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "confirmForChallengeDeadline", vertex)
}

// ConfirmForChallengeDeadline is a paid mutator transaction binding the contract method 0x542457ba.
//
// Solidity: function confirmForChallengeDeadline((uint256,bytes32,bool,bool,uint256) vertex) payable returns()
func (_AssertionChain *AssertionChainSession) ConfirmForChallengeDeadline(vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.Contract.ConfirmForChallengeDeadline(&_AssertionChain.TransactOpts, vertex)
}

// ConfirmForChallengeDeadline is a paid mutator transaction binding the contract method 0x542457ba.
//
// Solidity: function confirmForChallengeDeadline((uint256,bytes32,bool,bool,uint256) vertex) payable returns()
func (_AssertionChain *AssertionChainTransactorSession) ConfirmForChallengeDeadline(vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.Contract.ConfirmForChallengeDeadline(&_AssertionChain.TransactOpts, vertex)
}

// ConfirmForPSTimer is a paid mutator transaction binding the contract method 0xe7d3ded8.
//
// Solidity: function confirmForPSTimer((uint256,bytes32,bool,bool,uint256) vertex) payable returns()
func (_AssertionChain *AssertionChainTransactor) ConfirmForPSTimer(opts *bind.TransactOpts, vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "confirmForPSTimer", vertex)
}

// ConfirmForPSTimer is a paid mutator transaction binding the contract method 0xe7d3ded8.
//
// Solidity: function confirmForPSTimer((uint256,bytes32,bool,bool,uint256) vertex) payable returns()
func (_AssertionChain *AssertionChainSession) ConfirmForPSTimer(vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.Contract.ConfirmForPSTimer(&_AssertionChain.TransactOpts, vertex)
}

// ConfirmForPSTimer is a paid mutator transaction binding the contract method 0xe7d3ded8.
//
// Solidity: function confirmForPSTimer((uint256,bytes32,bool,bool,uint256) vertex) payable returns()
func (_AssertionChain *AssertionChainTransactorSession) ConfirmForPSTimer(vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.Contract.ConfirmForPSTimer(&_AssertionChain.TransactOpts, vertex)
}

// ConfirmForSubchallengeWin is a paid mutator transaction binding the contract method 0xecf361b2.
//
// Solidity: function confirmForSubchallengeWin((uint256,bytes32,bool,bool,uint256) vertex) payable returns()
func (_AssertionChain *AssertionChainTransactor) ConfirmForSubchallengeWin(opts *bind.TransactOpts, vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "confirmForSubchallengeWin", vertex)
}

// ConfirmForSubchallengeWin is a paid mutator transaction binding the contract method 0xecf361b2.
//
// Solidity: function confirmForSubchallengeWin((uint256,bytes32,bool,bool,uint256) vertex) payable returns()
func (_AssertionChain *AssertionChainSession) ConfirmForSubchallengeWin(vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.Contract.ConfirmForSubchallengeWin(&_AssertionChain.TransactOpts, vertex)
}

// ConfirmForSubchallengeWin is a paid mutator transaction binding the contract method 0xecf361b2.
//
// Solidity: function confirmForSubchallengeWin((uint256,bytes32,bool,bool,uint256) vertex) payable returns()
func (_AssertionChain *AssertionChainTransactorSession) ConfirmForSubchallengeWin(vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.Contract.ConfirmForSubchallengeWin(&_AssertionChain.TransactOpts, vertex)
}

// ConfirmForWin is a paid mutator transaction binding the contract method 0x14c0bad1.
//
// Solidity: function confirmForWin((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion) payable returns()
func (_AssertionChain *AssertionChainTransactor) ConfirmForWin(opts *bind.TransactOpts, assertion IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "confirmForWin", assertion)
}

// ConfirmForWin is a paid mutator transaction binding the contract method 0x14c0bad1.
//
// Solidity: function confirmForWin((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion) payable returns()
func (_AssertionChain *AssertionChainSession) ConfirmForWin(assertion IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.Contract.ConfirmForWin(&_AssertionChain.TransactOpts, assertion)
}

// ConfirmForWin is a paid mutator transaction binding the contract method 0x14c0bad1.
//
// Solidity: function confirmForWin((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion) payable returns()
func (_AssertionChain *AssertionChainTransactorSession) ConfirmForWin(assertion IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.Contract.ConfirmForWin(&_AssertionChain.TransactOpts, assertion)
}

// ConfirmNoRival is a paid mutator transaction binding the contract method 0xdd4f70c7.
//
// Solidity: function confirmNoRival((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion) payable returns()
func (_AssertionChain *AssertionChainTransactor) ConfirmNoRival(opts *bind.TransactOpts, assertion IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "confirmNoRival", assertion)
}

// ConfirmNoRival is a paid mutator transaction binding the contract method 0xdd4f70c7.
//
// Solidity: function confirmNoRival((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion) payable returns()
func (_AssertionChain *AssertionChainSession) ConfirmNoRival(assertion IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.Contract.ConfirmNoRival(&_AssertionChain.TransactOpts, assertion)
}

// ConfirmNoRival is a paid mutator transaction binding the contract method 0xdd4f70c7.
//
// Solidity: function confirmNoRival((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion) payable returns()
func (_AssertionChain *AssertionChainTransactorSession) ConfirmNoRival(assertion IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.Contract.ConfirmNoRival(&_AssertionChain.TransactOpts, assertion)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0x94898d79.
//
// Solidity: function createAssertion((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) prev, (uint256,bytes32) commit) payable returns((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion)
func (_AssertionChain *AssertionChainTransactor) CreateAssertion(opts *bind.TransactOpts, prev IAssertionChainAssertion, commit IAssertionChainStateCommitment) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "createAssertion", prev, commit)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0x94898d79.
//
// Solidity: function createAssertion((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) prev, (uint256,bytes32) commit) payable returns((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion)
func (_AssertionChain *AssertionChainSession) CreateAssertion(prev IAssertionChainAssertion, commit IAssertionChainStateCommitment) (*types.Transaction, error) {
	return _AssertionChain.Contract.CreateAssertion(&_AssertionChain.TransactOpts, prev, commit)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0x94898d79.
//
// Solidity: function createAssertion((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) prev, (uint256,bytes32) commit) payable returns((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion)
func (_AssertionChain *AssertionChainTransactorSession) CreateAssertion(prev IAssertionChainAssertion, commit IAssertionChainStateCommitment) (*types.Transaction, error) {
	return _AssertionChain.Contract.CreateAssertion(&_AssertionChain.TransactOpts, prev, commit)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xff3a6e56.
//
// Solidity: function createChallenge((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) prev) payable returns((uint256,uint256,(uint256,bytes32,bool,bool,uint256),(uint256,bytes32,bool,bool,uint256),uint256,address) challenge)
func (_AssertionChain *AssertionChainTransactor) CreateChallenge(opts *bind.TransactOpts, prev IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "createChallenge", prev)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xff3a6e56.
//
// Solidity: function createChallenge((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) prev) payable returns((uint256,uint256,(uint256,bytes32,bool,bool,uint256),(uint256,bytes32,bool,bool,uint256),uint256,address) challenge)
func (_AssertionChain *AssertionChainSession) CreateChallenge(prev IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.Contract.CreateChallenge(&_AssertionChain.TransactOpts, prev)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xff3a6e56.
//
// Solidity: function createChallenge((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) prev) payable returns((uint256,uint256,(uint256,bytes32,bool,bool,uint256),(uint256,bytes32,bool,bool,uint256),uint256,address) challenge)
func (_AssertionChain *AssertionChainTransactorSession) CreateChallenge(prev IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.Contract.CreateChallenge(&_AssertionChain.TransactOpts, prev)
}

// EligibleForNewSuccessor is a paid mutator transaction binding the contract method 0x38703628.
//
// Solidity: function eligibleForNewSuccessor((uint256,bytes32,bool,bool,uint256) vertex) returns(bool)
func (_AssertionChain *AssertionChainTransactor) EligibleForNewSuccessor(opts *bind.TransactOpts, vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "eligibleForNewSuccessor", vertex)
}

// EligibleForNewSuccessor is a paid mutator transaction binding the contract method 0x38703628.
//
// Solidity: function eligibleForNewSuccessor((uint256,bytes32,bool,bool,uint256) vertex) returns(bool)
func (_AssertionChain *AssertionChainSession) EligibleForNewSuccessor(vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.Contract.EligibleForNewSuccessor(&_AssertionChain.TransactOpts, vertex)
}

// EligibleForNewSuccessor is a paid mutator transaction binding the contract method 0x38703628.
//
// Solidity: function eligibleForNewSuccessor((uint256,bytes32,bool,bool,uint256) vertex) returns(bool)
func (_AssertionChain *AssertionChainTransactorSession) EligibleForNewSuccessor(vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.Contract.EligibleForNewSuccessor(&_AssertionChain.TransactOpts, vertex)
}

// IsPresumptiveSuccessor is a paid mutator transaction binding the contract method 0x9e4c5ab7.
//
// Solidity: function isPresumptiveSuccessor((uint256,bytes32,bool,bool,uint256) vertex) returns(bool)
func (_AssertionChain *AssertionChainTransactor) IsPresumptiveSuccessor(opts *bind.TransactOpts, vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "isPresumptiveSuccessor", vertex)
}

// IsPresumptiveSuccessor is a paid mutator transaction binding the contract method 0x9e4c5ab7.
//
// Solidity: function isPresumptiveSuccessor((uint256,bytes32,bool,bool,uint256) vertex) returns(bool)
func (_AssertionChain *AssertionChainSession) IsPresumptiveSuccessor(vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.Contract.IsPresumptiveSuccessor(&_AssertionChain.TransactOpts, vertex)
}

// IsPresumptiveSuccessor is a paid mutator transaction binding the contract method 0x9e4c5ab7.
//
// Solidity: function isPresumptiveSuccessor((uint256,bytes32,bool,bool,uint256) vertex) returns(bool)
func (_AssertionChain *AssertionChainTransactorSession) IsPresumptiveSuccessor(vertex IAssertionChainChallengeVertex) (*types.Transaction, error) {
	return _AssertionChain.Contract.IsPresumptiveSuccessor(&_AssertionChain.TransactOpts, vertex)
}

// Merge is a paid mutator transaction binding the contract method 0xeb9ca78b.
//
// Solidity: function merge((uint256,bytes32,bool,bool,uint256) mergingFrom, (uint256,bytes32,bool,bool,uint256) mergingTo, bytes32[] proof) payable returns((uint256,bytes32,bool,bool,uint256) mergedToVertex)
func (_AssertionChain *AssertionChainTransactor) Merge(opts *bind.TransactOpts, mergingFrom IAssertionChainChallengeVertex, mergingTo IAssertionChainChallengeVertex, proof [][32]byte) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "merge", mergingFrom, mergingTo, proof)
}

// Merge is a paid mutator transaction binding the contract method 0xeb9ca78b.
//
// Solidity: function merge((uint256,bytes32,bool,bool,uint256) mergingFrom, (uint256,bytes32,bool,bool,uint256) mergingTo, bytes32[] proof) payable returns((uint256,bytes32,bool,bool,uint256) mergedToVertex)
func (_AssertionChain *AssertionChainSession) Merge(mergingFrom IAssertionChainChallengeVertex, mergingTo IAssertionChainChallengeVertex, proof [][32]byte) (*types.Transaction, error) {
	return _AssertionChain.Contract.Merge(&_AssertionChain.TransactOpts, mergingFrom, mergingTo, proof)
}

// Merge is a paid mutator transaction binding the contract method 0xeb9ca78b.
//
// Solidity: function merge((uint256,bytes32,bool,bool,uint256) mergingFrom, (uint256,bytes32,bool,bool,uint256) mergingTo, bytes32[] proof) payable returns((uint256,bytes32,bool,bool,uint256) mergedToVertex)
func (_AssertionChain *AssertionChainTransactorSession) Merge(mergingFrom IAssertionChainChallengeVertex, mergingTo IAssertionChainChallengeVertex, proof [][32]byte) (*types.Transaction, error) {
	return _AssertionChain.Contract.Merge(&_AssertionChain.TransactOpts, mergingFrom, mergingTo, proof)
}

// RejectForLoss is a paid mutator transaction binding the contract method 0x73080da9.
//
// Solidity: function rejectForLoss((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion) payable returns()
func (_AssertionChain *AssertionChainTransactor) RejectForLoss(opts *bind.TransactOpts, assertion IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "rejectForLoss", assertion)
}

// RejectForLoss is a paid mutator transaction binding the contract method 0x73080da9.
//
// Solidity: function rejectForLoss((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion) payable returns()
func (_AssertionChain *AssertionChainSession) RejectForLoss(assertion IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.Contract.RejectForLoss(&_AssertionChain.TransactOpts, assertion)
}

// RejectForLoss is a paid mutator transaction binding the contract method 0x73080da9.
//
// Solidity: function rejectForLoss((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion) payable returns()
func (_AssertionChain *AssertionChainTransactorSession) RejectForLoss(assertion IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.Contract.RejectForLoss(&_AssertionChain.TransactOpts, assertion)
}

// RejectForPrev is a paid mutator transaction binding the contract method 0x6b97244f.
//
// Solidity: function rejectForPrev((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion) payable returns()
func (_AssertionChain *AssertionChainTransactor) RejectForPrev(opts *bind.TransactOpts, assertion IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "rejectForPrev", assertion)
}

// RejectForPrev is a paid mutator transaction binding the contract method 0x6b97244f.
//
// Solidity: function rejectForPrev((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion) payable returns()
func (_AssertionChain *AssertionChainSession) RejectForPrev(assertion IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.Contract.RejectForPrev(&_AssertionChain.TransactOpts, assertion)
}

// RejectForPrev is a paid mutator transaction binding the contract method 0x6b97244f.
//
// Solidity: function rejectForPrev((uint256,(uint256,bytes32),uint256,bool,uint256,uint256,address) assertion) payable returns()
func (_AssertionChain *AssertionChainTransactorSession) RejectForPrev(assertion IAssertionChainAssertion) (*types.Transaction, error) {
	return _AssertionChain.Contract.RejectForPrev(&_AssertionChain.TransactOpts, assertion)
}
