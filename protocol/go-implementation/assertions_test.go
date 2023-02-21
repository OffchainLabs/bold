package goimpl

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var (
	_ = protocol.AssertionChain(&AssertionChain{})
	_ = protocol.ChallengeManager(&AssertionChain{})
	_ = protocol.Assertion(&Assertion{})
	_ = protocol.Challenge(&Challenge{})
	_ = protocol.ChallengeVertex(&ChallengeVertex{})
	_ = protocol.ActiveTx(&ActiveTx{})
)

const testChallengePeriod = 100 * time.Second

func TestAssertionChain_ConfirmAndRefund(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeRef := util.NewArtificialTimeReference()
	correctBlockHashes := correctBlockHashesForTest(200)
	staker := common.BytesToAddress([]byte{1})

	assertionsChain := NewAssertionChain(ctx, timeRef, testChallengePeriod)
	require.Equal(t, 1, len(assertionsChain.assertions))
	require.Equal(t, protocol.AssertionSequenceNumber(0), assertionsChain.latestConfirmed)
	err := assertionsChain.Tx(func(tx protocol.ActiveTx) error {
		assertionsChain.SetBalance(tx, staker, AssertionStake)
		comm := util.StateCommitment{Height: 1, StateRoot: correctBlockHashes[99]}
		a1, err := assertionsChain.CreateAssertion(ctx, tx, comm.Height, 0, nil, nil, big.NewInt(0))
		require.NoError(t, err)
		require.Equal(t, uint64(0), assertionsChain.GetBalance(tx, staker).Uint64())

		comm = util.StateCommitment{Height: 2, StateRoot: correctBlockHashes[199]}
		a2, err := assertionsChain.CreateAssertion(ctx, tx, comm.Height, 1, nil, nil, big.NewInt(0))
		require.NoError(t, err)
		require.Equal(t, uint64(0), assertionsChain.GetBalance(tx, staker).Uint64())
		timeRef.Add(testChallengePeriod + time.Second)

		// Parent is confirmed. Staker should not get a refund because it's not a leaf.
		require.NoError(t, a1.(*Assertion).ConfirmNoRival(tx))
		require.Equal(t, uint64(0), assertionsChain.GetBalance(tx, staker).Uint64())

		// Child is confirmed. Staker should get a refund because it's a leaf.
		require.NoError(t, a2.(*Assertion).ConfirmNoRival(tx))
		require.Equal(t, AssertionStake.Uint64(), assertionsChain.GetBalance(tx, staker).Uint64())

		return nil
	})

	require.NoError(t, err)
}

func TestAssertionChain(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeRef := util.NewArtificialTimeReference()
	correctBlockHashes := correctBlockHashesForTest(200)
	wrongBlockHashes := wrongBlockHashesForTest(200)

	// Validators should agree at the 0th hash, but then
	// they diverge.
	wrongBlockHashes[0] = correctBlockHashes[0]

	staker1 := common.BytesToAddress([]byte{1})
	staker2 := common.BytesToAddress([]byte{2})

	assertionsChain := NewAssertionChain(ctx, timeRef, testChallengePeriod)
	require.Equal(t, 1, len(assertionsChain.assertions))
	require.Equal(t, protocol.AssertionSequenceNumber(0), assertionsChain.latestConfirmed)
	eventChan := make(chan AssertionChainEvent)
	err := assertionsChain.Tx(func(tx protocol.ActiveTx) error {
		genesis, err := assertionsChain.LatestConfirmed(ctx, tx)
		require.NoError(t, err)
		require.Equal(t, util.StateCommitment{
			Height:    0,
			StateRoot: common.Hash{},
		}, genesis.StateHash())
		assertionsChain.SetBalance(tx, staker1, big.NewInt(0).Add(AssertionStake, ChallengeVertexStake))
		assertionsChain.SetBalance(tx, staker2, big.NewInt(0).Add(AssertionStake, ChallengeVertexStake))

		assertionsChain.feed.SubscribeWithFilter(ctx, eventChan, func(ev AssertionChainEvent) bool {
			switch ev.(type) {
			case *SetBalanceEvent:
				return false
			default:
				return true
			}
		})

		// add an assertion, then confirm it
		comm := util.StateCommitment{Height: 1, StateRoot: correctBlockHashes[0]}
		newAssertion, err := assertionsChain.CreateAssertion(ctx, tx, comm.Height, 0, nil, nil, nil)
		require.NoError(t, err)
		require.Equal(t, 2, len(assertionsChain.assertions))
		confirmed, err := assertionsChain.LatestConfirmed(ctx, tx)
		require.NoError(t, err)
		require.Equal(t, genesis, confirmed)
		verifyCreateLeafEventInFeed(t, eventChan, 1, 0, staker1, comm)

		err = newAssertion.(*Assertion).ConfirmNoRival(tx)
		require.ErrorIs(t, err, ErrNotYet)
		timeRef.Add(testChallengePeriod + time.Second)
		require.NoError(t, newAssertion.(*Assertion).ConfirmNoRival(tx))

		confirmed, err = assertionsChain.LatestConfirmed(ctx, tx)
		require.NoError(t, err)
		require.Equal(t, newAssertion, confirmed)
		require.Equal(t, ConfirmedAssertionState, int(newAssertion.(*Assertion).status))
		verifyConfirmEventInFeed(t, eventChan, protocol.AssertionSequenceNumber(1))

		// try to create a duplicate assertion
		_, err = assertionsChain.CreateAssertion(ctx, tx, comm.Height, 0, nil, nil, nil)
		require.ErrorIs(t, err, ErrVertexAlreadyExists)

		// create a fork, let first branch win by timeout
		comm = util.StateCommitment{Height: 4, StateRoot: correctBlockHashes[3]}
		branch1, err := assertionsChain.CreateAssertion(ctx, tx, comm.Height, uint64(newAssertion.SeqNum()), nil, nil, nil)
		require.NoError(t, err)

		timeRef.Add(5 * time.Second)
		verifyCreateLeafEventInFeed(t, eventChan, 2, 1, staker1, comm)
		comm = util.StateCommitment{Height: 4, StateRoot: wrongBlockHashes[3]}
		branch2, err := assertionsChain.CreateAssertion(ctx, tx, comm.Height, uint64(newAssertion.SeqNum()), nil, nil, nil)
		require.NoError(t, err)

		// Assert the creation event.
		verifyCreateLeafEventInFeed(t, eventChan, 3, 1, staker2, comm)

		// Create a challenge at the fork.
		challenge, err := assertionsChain.CreateSuccessionChallenge(ctx, tx, newAssertion.SeqNum())
		require.NoError(t, err)
		verifyStartChallengeEventInFeed(t, eventChan, newAssertion.SeqNum())

		// Add two competing challenge leaves.
		// The last hash must be the state root of the assertion
		// we are targeting.
		hashes := correctBlockHashes[:4]
		require.Equal(t, hashes[len(hashes)-1], branch1.StateHash())

		// We commit to a height that is equal to assertion.height - assertion.prev.height.
		// That is, we are committing to a range of heights from the prev
		// assertion to the assertion we are targeting.
		prevHeight := branch1.(*Assertion).Prev.Unwrap().StateCommitment.Height
		height := branch1.Height() - prevHeight

		historyCommit, err := util.NewHistoryCommitment(
			height,
			hashes,
			util.WithLastElementProof(hashes),
		)
		require.NoError(t, err)

		chal1, err := challenge.AddBlockChallengeLeaf(ctx, tx, branch1, historyCommit)
		require.NoError(t, err)

		badCommit, err := util.NewHistoryCommitment(
			height,
			wrongBlockHashes[:height],
			util.WithLastElementProof(wrongBlockHashes[:height+1]),
		)
		require.NoError(t, err)

		_, err = challenge.AddBlockChallengeLeaf(ctx, tx, branch2, badCommit)
		require.NoError(t, err)

		// Cannot be confirmed yet.
		err = chal1.ConfirmForPsTimer(ctx, tx)
		require.ErrorIs(t, err, ErrNotYet)

		// Add a challenge period, and then the leaf can be confirmed.
		timeRef.Add(testChallengePeriod)
		require.NoError(t, chal1.ConfirmForPsTimer(ctx, tx))

		half := big.NewInt(0).Div(ChallengeVertexStake, big.NewInt(2))
		want := big.NewInt(0).Add(half, ChallengeVertexStake)
		chal1Validator, _ := chal1.MiniStaker(ctx, tx)
		require.Equal(t, want, assertionsChain.GetBalance(tx, chal1Validator)) // Should receive own mini stake plus half of others.
		require.NoError(t, branch1.(*Assertion).ConfirmForWin(ctx, tx))
		confirmed, err = assertionsChain.LatestConfirmed(ctx, tx)
		require.NoError(t, err)
		require.Equal(t, branch1, confirmed)

		verifyConfirmEventInFeed(t, eventChan, 2)
		require.NoError(t, branch2.(*Assertion).RejectForLoss(ctx, tx))
		verifyRejectEventInFeed(t, eventChan, 3)
		return nil
	})
	require.NoError(t, err)

	// verify that feed is empty
	time.Sleep(500 * time.Millisecond)
	select {
	case ev := <-eventChan:
		t.Fatal(ev)
	default:
	}
}

func TestAssertionChain_CreateLeaf_MustHaveValidParent(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeRef := util.NewArtificialTimeReference()
	staker := common.BytesToAddress([]byte{1})

	assertionsChain := NewAssertionChain(ctx, timeRef, testChallengePeriod)
	require.Equal(t, 1, len(assertionsChain.assertions))
	require.Equal(t, protocol.AssertionSequenceNumber(0), assertionsChain.latestConfirmed)
	err := assertionsChain.Tx(func(tx protocol.ActiveTx) error {
		genesis, err := assertionsChain.LatestConfirmed(ctx, tx)
		require.Equal(t, util.StateCommitment{
			Height:    0,
			StateRoot: common.Hash{},
		}, genesis.StateHash())

		bigBalance := new(big.Int).Mul(AssertionStake, big.NewInt(1000))
		assertionsChain.SetBalance(tx, staker, bigBalance)

		foo := common.BytesToHash([]byte("foo"))
		bar := common.BytesToHash([]byte("bar"))
		_ = bar
		comm := util.StateCommitment{Height: 1, StateRoot: foo}
		leaf, err := assertionsChain.CreateAssertion(ctx, tx, comm.Height, 0, nil, nil, nil)
		require.NoError(t, err)

		// Trying to create a new leaf with the same commitment as before should fail.
		leaf.(*Assertion).StateCommitment = util.StateCommitment{Height: 0, StateRoot: bar} // Mutate leaf.
		_, err = assertionsChain.CreateAssertion(ctx, tx, comm.Height, 1, nil, nil, nil)
		require.ErrorIs(t, err, ErrVertexAlreadyExists)

		// Trying to create a new leaf on top of a non-existent parent should fail.
		leaf.(*Assertion).StateCommitment = util.StateCommitment{Height: 0, StateRoot: bar} // Mutate leaf.
		comm = util.StateCommitment{Height: 2, StateRoot: foo}
		_, err = assertionsChain.CreateAssertion(ctx, tx, comm.Height, 1, nil, nil, nil)
		require.ErrorIs(t, err, ErrParentDoesNotExist)
		return nil
	})
	require.NoError(t, err)
}

func TestAssertionChain_LeafCreationThroughDiffStakers(t *testing.T) {
	ctx := context.Background()
	assertionsChain := NewAssertionChain(ctx, util.NewArtificialTimeReference(), testChallengePeriod)

	require.NoError(t, assertionsChain.Tx(func(tx protocol.ActiveTx) error {
		oldStaker := common.BytesToAddress([]byte{1})
		staker := common.BytesToAddress([]byte{2})
		require.Equal(t, assertionsChain.GetBalance(tx, oldStaker), big.NewInt(0)) // Old staker has 0 because it's already staked.
		assertionsChain.SetBalance(tx, staker, AssertionStake)
		require.Equal(t, assertionsChain.GetBalance(tx, staker), AssertionStake) // New staker has full balance because it's not yet staked.

		lc, err := assertionsChain.LatestConfirmed(ctx, tx)
		require.NoError(t, err)
		lc.(*Assertion).Staker = util.Some[common.Address](oldStaker)
		_, err = assertionsChain.CreateAssertion(ctx, tx, 1, 0, nil, nil, nil)
		require.NoError(t, err)

		require.Equal(t, assertionsChain.GetBalance(tx, staker), big.NewInt(0))     // New staker has 0 balance after staking.
		require.Equal(t, assertionsChain.GetBalance(tx, oldStaker), AssertionStake) // Old staker has full balance after unstaking.
		return nil
	}))
}

func TestAssertionChain_LeafCreationsInsufficientStakes(t *testing.T) {
	ctx := context.Background()
	assertionsChain := NewAssertionChain(ctx, util.NewArtificialTimeReference(), testChallengePeriod)

	require.NoError(t, assertionsChain.Tx(func(tx protocol.ActiveTx) error {
		lc, err := assertionsChain.LatestConfirmed(ctx, tx)
		require.NoError(t, err)
		lc.(*Assertion).Staker = util.None[common.Address]()
		_, err = assertionsChain.CreateAssertion(ctx, tx, 1, 0, nil, nil, nil)
		require.ErrorIs(t, err, ErrInsufficientBalance)

		diffStaker := common.BytesToAddress([]byte{2})
		lc.(*Assertion).Staker = util.Some[common.Address](diffStaker)
		_, err = assertionsChain.CreateAssertion(ctx, tx, 1, 0, nil, nil, nil)
		require.ErrorIs(t, err, ErrInsufficientBalance)
		return nil
	}))
}

func verifyCreateLeafEventInFeed(t *testing.T, c <-chan AssertionChainEvent, seqNum, prevSeqNum protocol.AssertionSequenceNumber, staker common.Address, comm util.StateCommitment) {
	t.Helper()
	ev := <-c
	switch e := ev.(type) {
	case *CreateLeafEvent:
		if e.SeqNum != seqNum || e.PrevSeqNum != prevSeqNum || e.Validator != staker || e.StateCommitment != comm {
			t.Fatal(e)
		}
	default:
		t.Fatal(e)
	}
}

func verifyConfirmEventInFeed(t *testing.T, c <-chan AssertionChainEvent, seqNum protocol.AssertionSequenceNumber) {
	t.Helper()
	ev := <-c
	switch e := ev.(type) {
	case *ConfirmEvent:
		require.Equal(t, seqNum, e.SeqNum)
	default:
		t.Fatal()
	}
}

func verifyRejectEventInFeed(t *testing.T, c <-chan AssertionChainEvent, seqNum protocol.AssertionSequenceNumber) {
	t.Helper()
	ev := <-c
	switch e := ev.(type) {
	case *RejectEvent:
		require.Equal(t, seqNum, e.SeqNum)
	default:
		t.Fatal()
	}
}

func verifyStartChallengeEventInFeed(t *testing.T, c <-chan AssertionChainEvent, parentSeqNum protocol.AssertionSequenceNumber) {
	t.Helper()
	ev := <-c
	switch e := ev.(type) {
	case *StartChallengeEvent:
		require.Equal(t, parentSeqNum, e.ParentSeqNum)
	default:
		t.Fatal()
	}
}

func TestIsAtOneStepFork(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeRef := util.NewArtificialTimeReference()
	assertionsChain := NewAssertionChain(ctx, timeRef, testChallengePeriod)
	genesisCommitHash := protocol.ChallengeHash((util.StateCommitment{}).Hash())

	tests := []struct {
		name         string
		vertexHeight uint64
		parentHeight uint64
		vertices     map[protocol.VertexHash]*ChallengeVertex
		want         bool
	}{
		{
			name:         "height difference != 1 in inputs",
			vertexHeight: 2,
			parentHeight: 0,
			want:         false,
			vertices:     nil,
		},
		{
			name:         "empty list of vertices despite height difference == 1 in inputs",
			vertexHeight: 1,
			parentHeight: 0,
			want:         false,
			vertices:     nil,
		},
		{
			name:         "only one vertex",
			vertexHeight: 1,
			parentHeight: 0,
			want:         false,
			vertices: map[protocol.VertexHash]*ChallengeVertex{
				protocol.VertexHash{1}: &ChallengeVertex{
					PrevV: util.Some(&ChallengeVertex{
						Commitment: util.HistoryCommitment{},
					}),
					Commitment: util.HistoryCommitment{
						Height: 1,
						Merkle: common.BytesToHash([]byte{1}),
					},
				},
			},
		},
		{
			name:         "no vertices with matching parent commitment",
			vertexHeight: 1,
			parentHeight: 0,
			want:         false,
			vertices: map[protocol.VertexHash]*ChallengeVertex{
				protocol.VertexHash{1}: &ChallengeVertex{
					Prev: util.Some(&ChallengeVertex{
						Commitment: util.HistoryCommitment{
							Height: 5,
							Merkle: common.BytesToHash([]byte{5}),
						},
					}),
					Commitment: util.HistoryCommitment{
						Height: 6,
						Merkle: common.BytesToHash([]byte{6}),
					},
				},
			},
		},
		{
			name:         "two vertices but only one is has height difference == 1",
			vertexHeight: 1,
			parentHeight: 0,
			want:         false,
			vertices: map[protocol.VertexHash]*ChallengeVertex{
				protocol.VertexHash{1}: &ChallengeVertex{
					Prev: util.Some(&ChallengeVertex{
						Commitment: util.HistoryCommitment{},
					}),
					Commitment: util.HistoryCommitment{
						Height: 1,
						Merkle: common.BytesToHash([]byte{1}),
					},
				},
				protocol.VertexHash{2}: &ChallengeVertex{
					Prev: util.Some(&ChallengeVertex{
						Commitment: util.HistoryCommitment{},
					}),
					Commitment: util.HistoryCommitment{
						Height: 2,
						Merkle: common.BytesToHash([]byte{1}),
					},
				},
			},
		},
		{
			name:         "two vertices at one-step-fork",
			vertexHeight: 1,
			parentHeight: 0,
			want:         true,
			vertices: map[protocol.VertexHash]*ChallengeVertex{
				protocol.VertexHash{1}: &ChallengeVertex{
					Prev: util.Some(ChallengeVertexInterface(&ChallengeVertex{
						Commitment: util.HistoryCommitment{},
					})),
					Commitment: util.HistoryCommitment{
						Height: 1,
						Merkle: common.BytesToHash([]byte{1}),
					},
				},
				protocol.VertexHash{2}: &ChallengeVertex{
					Prev: util.Some(ChallengeVertexInterface(&ChallengeVertex{
						Commitment: util.HistoryCommitment{},
					})),
					Commitment: util.HistoryCommitment{
						Height: 1,
						Merkle: common.BytesToHash([]byte{2}),
					},
				},
			},
		},
		{
			name:         "three vertices with only two at one-step-fork",
			vertexHeight: 1,
			parentHeight: 0,
			want:         false,
			vertices: map[protocol.VertexHash]*ChallengeVertex{
				protocol.VertexHash{1}: &ChallengeVertex{
					Prev: util.Some(&ChallengeVertex{
						Commitment: util.HistoryCommitment{},
					}),
					Commitment: util.HistoryCommitment{
						Height: 1,
						Merkle: common.BytesToHash([]byte{1}),
					},
				},
				protocol.VertexHash{2}: &ChallengeVertex{
					Prev: util.Some(&ChallengeVertex{
						Commitment: util.HistoryCommitment{},
					}),
					Commitment: util.HistoryCommitment{
						Height: 1,
						Merkle: common.BytesToHash([]byte{2}),
					},
				},
				protocol.VertexHash{3}: &ChallengeVertex{
					Prev: util.Some(&ChallengeVertex{
						Commitment: util.HistoryCommitment{},
					}),
					Commitment: util.HistoryCommitment{
						Height: 2,
						Merkle: common.BytesToHash([]byte{3}),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := assertionsChain.Tx(func(tx *ActiveTx) error {
				vertexCommit := util.HistoryCommitment{
					Height: tt.vertexHeight,
				}
				parentCommit := util.HistoryCommitment{
					Height: tt.parentHeight,
				}
				_ = parentCommit
				_ = vertexCommit
				assertionsChain.challengeVerticesByCommitHash = make(map[protocol.ChallengeHash]map[protocol.VertexHash]*ChallengeVertex)
				assertionsChain.challengeVerticesByCommitHash[genesisCommitHash] = tt.vertices
				// ok, err := assertionsChain.IsAtOneStepFork(
				// 	ctx,
				// 	tx,
				// 	genesisCommitHash,
				// 	vertexCommit,
				// 	parentCommit,
				// )
				// require.NoError(t, err)
				// require.Equal(t, tt.want, ok)
				return nil
			})
			require.NoError(t, err)
		})
	}
}

func TestChallengeVertexByHistoryCommit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeRef := util.NewArtificialTimeReference()
	assertionsChain := NewAssertionChain(ctx, timeRef, testChallengePeriod)

	err := assertionsChain.Tx(func(tx *ActiveTx) error {
		genesisCommitHash := protocol.ChallengeHash((util.StateCommitment{}).Hash())
		t.Run("vertices not found for challenge", func(t *testing.T) {
			vertexCommit := util.HistoryCommitment{
				Height: 1,
			}
			_, err := assertionsChain.ChallengeVertexByCommitHash(
				tx,
				genesisCommitHash,
				protocol.VertexHash(vertexCommit.Merkle),
			)
			require.ErrorContains(t, err, "challenge vertices not found")
		})
		t.Run("vertex with commit not found", func(t *testing.T) {
			vertexCommit := util.HistoryCommitment{
				Height: 1,
			}
			vertices := map[protocol.VertexHash]*ChallengeVertex{}
			assertionsChain.challengeVerticesByCommitHash[genesisCommitHash] = vertices
			_, err := assertionsChain.ChallengeVertexByCommitHash(
				tx,
				genesisCommitHash,
				protocol.VertexHash(vertexCommit.Merkle),
			)
			require.ErrorContains(t, err, "not found")
		})
		t.Run("vertex found", func(t *testing.T) {
			vertexCommit := util.HistoryCommitment{
				Height: 1,
				Merkle: common.Hash(protocol.VertexHash{10}),
			}
			want := &ChallengeVertex{
				Commitment: vertexCommit,
			}
			vertices := map[protocol.VertexHash]*ChallengeVertex{
				protocol.VertexHash{10}: want,
			}
			assertionsChain.challengeVerticesByCommitHash[genesisCommitHash] = vertices
			got, err := assertionsChain.ChallengeVertexByCommitHash(
				tx,
				genesisCommitHash,
				protocol.VertexHash(vertexCommit.Merkle),
			)
			require.NoError(t, err)
			require.Equal(t, want, got)
		})
		return nil
	})
	require.NoError(t, err)
}

func TestAssertionChain_BlockChallenge_CreateLeafInvariants(t *testing.T) {
	ctx := context.Background()
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	validator := common.BytesToAddress([]byte("foo"))
	t.Run("prev does not match root assertion", func(t *testing.T) {
		c := &Challenge{}
		assertion := &Assertion{
			Prev: util.None[*Assertion](),
		}
		_, err := c.AddLeaf(
			ctx,
			tx,
			assertion,
			util.HistoryCommitment{},
			validator,
		)
		require.ErrorIs(t, err, ErrInvalidOp)

		c = &Challenge{
			rootAssertion: util.Some(&Assertion{
				SequenceNum: 1,
			}),
		}
		assertion = &Assertion{
			Prev: util.Some(&Assertion{
				SequenceNum: 2,
			}),
		}
		_, err = c.AddLeaf(
			ctx,
			tx,
			assertion,
			util.HistoryCommitment{},
			validator,
		)
		require.ErrorIs(t, err, ErrInvalidOp)
	})
	t.Run("challenge already complete", func(t *testing.T) {
		c := &Challenge{
			rootAssertion: util.Some(&Assertion{
				SequenceNum: 1,
			}),
			WinnerAssertion: util.Some(&Assertion{}),
		}
		assertion := &Assertion{
			Prev: c.rootAssertion,
		}
		_, err := c.AddLeaf(
			ctx,
			tx,
			assertion,
			util.HistoryCommitment{},
			validator,
		)
		require.ErrorIs(t, err, ErrWrongState)
	})
	t.Run("ineligible for new successor", func(t *testing.T) {
		ref := util.NewArtificialTimeReference()
		c := &Challenge{
			rootAssertion: util.Some(&Assertion{
				SequenceNum: 1,
				challengeManager: &AssertionChain{
					challengePeriod: time.Minute,
				},
			}),
			WinnerAssertion: util.None[*Assertion](),
		}
		timer := util.NewCountUpTimer(ref)
		timer.Add(2 * time.Minute)
		rootVertex := &ChallengeVertex{
			PresumptiveSuccessor: util.Some(&ChallengeVertex{
				PsTimer: timer,
			}),
			Challenge: util.Some(c),
		}
		c.rootVertex = util.Some(rootVertex)
		assertion := &Assertion{
			Prev: c.rootAssertion,
		}
		_, err := c.AddLeaf(
			ctx,
			tx,
			assertion,
			util.HistoryCommitment{},
			validator,
		)
		require.ErrorIs(t, err, ErrPastDeadline)
	})
	t.Run("vertex already exists", func(t *testing.T) {
		history := util.HistoryCommitment{}
		c := &Challenge{
			rootAssertion: util.Some(&Assertion{
				SequenceNum: 1,
				challengeManager: &AssertionChain{
					challengePeriod: time.Minute,
				},
			}),
			includedHistories: map[common.Hash]bool{
				history.Hash(): true,
			},
		}
		c.rootVertex = util.Some(&ChallengeVertex{})
		assertion := &Assertion{
			Prev: c.rootAssertion,
		}
		_, err := c.AddLeaf(
			ctx,
			tx,
			assertion,
			history,
			validator,
		)
		require.ErrorIs(t, err, ErrVertexAlreadyExists)
	})
	t.Run("insufficient balance", func(t *testing.T) {
		history := util.HistoryCommitment{}
		c := &Challenge{
			rootAssertion: util.Some(&Assertion{
				SequenceNum: 1,
				challengeManager: &AssertionChain{
					challengePeriod: time.Minute,
					balances: util.NewMapWithDefaultAdvanced[common.Address](
						common.Big0,
						func(x *big.Int) bool { return x.Sign() == 0 },
					),
				},
			}),
			includedHistories: make(map[common.Hash]bool),
		}
		c.rootVertex = util.Some(&ChallengeVertex{})
		assertion := &Assertion{
			Prev: c.rootAssertion,
		}
		_, err := c.AddLeaf(
			ctx,
			tx,
			assertion,
			history,
			validator,
		)
		require.ErrorIs(t, err, ErrInsufficientBalance)
	})
	t.Run("no proof of last leaf provided", func(t *testing.T) {
		balances := util.NewMapWithDefaultAdvanced[common.Address](
			common.Big0,
			func(x *big.Int) bool { return x.Sign() == 0 },
		)
		balances.Set(validator, ChallengeVertexStake)

		history := util.HistoryCommitment{}
		c := &Challenge{
			rootAssertion: util.Some(&Assertion{
				SequenceNum: 1,
				challengeManager: &AssertionChain{
					challengePeriod: time.Minute,
					balances:        balances,
					feed:            NewEventFeed[AssertionChainEvent](ctx),
				},
			}),
			includedHistories: make(map[common.Hash]bool),
		}
		c.rootVertex = util.Some(&ChallengeVertex{})
		assertion := &Assertion{
			Prev: c.rootAssertion,
		}
		_, err := c.AddLeaf(
			ctx,
			tx,
			assertion,
			history,
			validator,
		)
		require.ErrorIs(t, err, ErrNoLastLeafProof)
	})
	t.Run("last leaf does not match assertion state root", func(t *testing.T) {
		balances := util.NewMapWithDefaultAdvanced[common.Address](
			common.Big0,
			func(x *big.Int) bool { return x.Sign() == 0 },
		)
		balances.Set(validator, ChallengeVertexStake)

		c := &Challenge{
			rootAssertion: util.Some(&Assertion{
				SequenceNum: 1,
				challengeManager: &AssertionChain{
					challengePeriod: time.Minute,
					balances:        balances,
					feed:            NewEventFeed[AssertionChainEvent](ctx),
				},
			}),
			includedHistories: make(map[common.Hash]bool),
		}
		c.rootVertex = util.Some(&ChallengeVertex{})
		assertion := &Assertion{
			Prev: c.rootAssertion,
		}

		hashes := correctBlockHashesForTest(10)
		history, err := util.NewHistoryCommitment(
			5,
			hashes[:5],
			util.WithLastElementProof(hashes[:5]),
		)
		require.NoError(t, err)
		_, err = c.AddLeaf(
			ctx,
			tx,
			assertion,
			history,
			validator,
		)
		require.ErrorIs(t, err, ErrWrongLastLeaf)
	})
	t.Run("first leaf must be the previous assertions state root", func(t *testing.T) {
		balances := util.NewMapWithDefaultAdvanced[common.Address](
			common.Big0,
			func(x *big.Int) bool { return x.Sign() == 0 },
		)
		balances.Set(validator, ChallengeVertexStake)

		hashes := correctBlockHashesForTest(10)
		c := &Challenge{
			rootAssertion: util.Some(&Assertion{
				SequenceNum: 1,
				StateCommitment: util.StateCommitment{
					Height:    5,
					StateRoot: hashes[5],
				},
				challengeManager: &AssertionChain{
					challengePeriod: time.Minute,
					balances:        balances,
					feed:            NewEventFeed[AssertionChainEvent](ctx),
				},
			}),
			includedHistories: make(map[common.Hash]bool),
		}
		c.rootVertex = util.Some(&ChallengeVertex{})
		assertion := &Assertion{
			Prev: c.rootAssertion,
			StateCommitment: util.StateCommitment{
				Height:    3,
				StateRoot: hashes[5],
			},
		}

		history, err := util.NewHistoryCommitment(
			5,
			hashes[:5],
			util.WithLastElementProof(hashes[:6]),
		)
		require.NoError(t, err)
		_, err = c.AddLeaf(
			ctx,
			tx,
			assertion,
			history,
			validator,
		)
		require.ErrorIs(t, err, ErrWrongFirstLeaf)
	})
	t.Run("prev height must be less than current height", func(t *testing.T) {
		balances := util.NewMapWithDefaultAdvanced[common.Address](
			common.Big0,
			func(x *big.Int) bool { return x.Sign() == 0 },
		)
		balances.Set(validator, ChallengeVertexStake)

		hashes := correctBlockHashesForTest(10)
		c := &Challenge{
			rootAssertion: util.Some(&Assertion{
				SequenceNum: 1,
				StateCommitment: util.StateCommitment{
					Height:    5,
					StateRoot: hashes[0],
				},
				challengeManager: &AssertionChain{
					challengePeriod: time.Minute,
					balances:        balances,
					feed:            NewEventFeed[AssertionChainEvent](ctx),
				},
			}),
			includedHistories: make(map[common.Hash]bool),
		}
		c.rootVertex = util.Some(&ChallengeVertex{})
		assertion := &Assertion{
			Prev: c.rootAssertion,
			StateCommitment: util.StateCommitment{
				Height:    3,
				StateRoot: hashes[5],
			},
		}

		history, err := util.NewHistoryCommitment(
			5,
			hashes[:5],
			util.WithLastElementProof(hashes[:6]),
		)
		require.NoError(t, err)
		_, err = c.AddLeaf(
			ctx,
			tx,
			assertion,
			history,
			validator,
		)
		require.ErrorIs(t, err, ErrInvalidHeight)
	})
	t.Run("claimed height must be range of curr - prev's heights", func(t *testing.T) {
		balances := util.NewMapWithDefaultAdvanced[common.Address](
			common.Big0,
			func(x *big.Int) bool { return x.Sign() == 0 },
		)
		balances.Set(validator, ChallengeVertexStake)

		hashes := correctBlockHashesForTest(10)
		c := &Challenge{
			rootAssertion: util.Some(&Assertion{
				SequenceNum: 1,
				StateCommitment: util.StateCommitment{
					Height:    5,
					StateRoot: hashes[0],
				},
				challengeManager: &AssertionChain{
					challengePeriod: time.Minute,
					balances:        balances,
					feed:            NewEventFeed[AssertionChainEvent](ctx),
				},
			}),
			includedHistories: make(map[common.Hash]bool),
		}
		c.rootVertex = util.Some(&ChallengeVertex{})
		assertion := &Assertion{
			Prev: c.rootAssertion,
			StateCommitment: util.StateCommitment{
				Height:    8,
				StateRoot: hashes[8],
			},
		}

		history, err := util.NewHistoryCommitment(
			4,
			hashes[:8],
			util.WithLastElementProof(hashes[:9]),
		)
		require.NoError(t, err)
		_, err = c.AddLeaf(
			ctx,
			tx,
			assertion,
			history,
			validator,
		)
		require.ErrorIs(t, err, ErrInvalidHeight)
	})
	t.Run("commitment should prove the last element in the Merkleization is the last leaf", func(t *testing.T) {
		balances := util.NewMapWithDefaultAdvanced[common.Address](
			common.Big0,
			func(x *big.Int) bool { return x.Sign() == 0 },
		)
		balances.Set(validator, ChallengeVertexStake)

		hashes := correctBlockHashesForTest(10)
		c := &Challenge{
			rootAssertion: util.Some(&Assertion{
				SequenceNum: 1,
				StateCommitment: util.StateCommitment{
					Height:    5,
					StateRoot: hashes[5],
				},
				challengeManager: &AssertionChain{
					challengePeriod: time.Minute,
					balances:        balances,
					feed:            NewEventFeed[AssertionChainEvent](ctx),
				},
			}),
			includedHistories: make(map[common.Hash]bool),
		}
		c.rootVertex = util.Some(&ChallengeVertex{})
		assertion := &Assertion{
			Prev: c.rootAssertion,
			StateCommitment: util.StateCommitment{
				Height:    8,
				StateRoot: hashes[8],
			},
		}

		history, err := util.NewHistoryCommitment(
			3,
			hashes[5:8],
			util.WithLastElementProof(hashes[5:9]),
		)
		require.NoError(t, err)

		// Corrupt the Merkle proof.
		history.LastLeafProof[0] = common.BytesToHash([]byte("nyan"))

		_, err = c.AddLeaf(
			ctx,
			tx,
			assertion,
			history,
			validator,
		)
		require.ErrorIs(t, err, ErrProofFailsToVerify)
	})
	t.Run("OK", func(t *testing.T) {
		ref := util.NewArtificialTimeReference()
		balances := util.NewMapWithDefaultAdvanced[common.Address](
			common.Big0,
			func(x *big.Int) bool { return x.Sign() == 0 },
		)
		balances.Set(validator, ChallengeVertexStake)
		chain := &AssertionChain{
			timeReference:                 ref,
			challengePeriod:               time.Minute,
			balances:                      balances,
			feed:                          NewEventFeed[AssertionChainEvent](ctx),
			challengesFeed:                NewEventFeed[ChallengeEvent](ctx),
			challengesByCommitHash:        make(map[protocol.ChallengeHash]*Challenge),
			challengeVerticesByCommitHash: make(map[protocol.ChallengeHash]map[protocol.VertexHash]*ChallengeVertex),
		}

		hashes := correctBlockHashesForTest(10)
		c := &Challenge{
			rootAssertion: util.Some(&Assertion{
				SequenceNum: 1,
				StateCommitment: util.StateCommitment{
					Height:    5,
					StateRoot: hashes[5],
				},
				challengeManager: chain,
			}),
			includedHistories: make(map[common.Hash]bool),
		}

		chalHash := protocol.ChallengeHash(c.rootAssertion.Unwrap().StateCommitment.Hash())
		chain.challengeVerticesByCommitHash[chalHash] = make(map[protocol.VertexHash]*ChallengeVertex)

		c.rootVertex = util.Some(&ChallengeVertex{})
		assertion := &Assertion{
			Prev:             c.rootAssertion,
			challengeManager: chain,
			StateCommitment: util.StateCommitment{
				Height:    8,
				StateRoot: hashes[8],
			},
		}

		history, err := util.NewHistoryCommitment(
			3,
			hashes[5:8],
			util.WithLastElementProof(hashes[5:9]),
		)
		require.NoError(t, err)
		_, err = c.AddLeaf(
			ctx,
			tx,
			assertion,
			history,
			validator,
		)
		require.NoError(t, err)
	})
}

func TestAssertionChain_Bisect(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeRef := util.NewArtificialTimeReference()
	correctBlockHashes := correctBlockHashesForTest(10)
	wrongBlockHashes := wrongBlockHashesForTest(10)
	staker1 := common.BytesToAddress([]byte{1})
	staker2 := common.BytesToAddress([]byte{2})

	assertionsChain := NewAssertionChain(ctx, timeRef, testChallengePeriod)

	err := assertionsChain.Tx(func(tx protocol.ActiveTx) error {
		// We create a fork with genesis as the rootAssertion, where one branch is a higher depth than the other.
		genesis, err := assertionsChain.LatestConfirmed(ctx, tx)
		require.NoError(t, err)
		bigBalance := new(big.Int).Mul(AssertionStake, big.NewInt(1000))
		assertionsChain.SetBalance(tx, staker1, bigBalance)
		assertionsChain.SetBalance(tx, staker2, bigBalance)

		correctBranch, err := assertionsChain.CreateLeaf(tx, genesis, util.StateCommitment{Height: 6, StateRoot: correctBlockHashes[6]}, staker1)
		require.NoError(t, err)
		wrongBranch, err := assertionsChain.CreateLeaf(tx, genesis, util.StateCommitment{Height: 6, StateRoot: wrongBlockHashes[6]}, staker2)
		require.NoError(t, err)

		challenge, err := genesis.CreateChallenge(tx, ctx, staker2)
		require.NoError(t, err)

		// Add some leaves to the mix...
		expectedBisectionHeight := uint64(4)
		lo := expectedBisectionHeight

		hi := uint64(6)
		loExp := util.ExpansionFromLeaves(wrongBlockHashes[:lo])
		badCommit, err := util.NewHistoryCommitment(
			hi,
			wrongBlockHashes[:hi],
			util.WithLastElementProof(wrongBlockHashes[:hi+1]),
		)
		require.NoError(t, err)

		badLeaf, err := challenge.AddLeaf(
			ctx,
			tx,
			wrongBranch,
			badCommit,
			staker1,
		)
		require.NoError(t, err)

		goodCommit, err := util.NewHistoryCommitment(
			hi,
			correctBlockHashes[:hi],
			util.WithLastElementProof(correctBlockHashes[:hi+1]),
		)
		require.NoError(t, err)
		goodLeaf, err := challenge.AddLeaf(
			ctx,
			tx,
			correctBranch,
			goodCommit,
			staker2,
		)
		require.NoError(t, err)

		// Ensure the lower height challenge vertex is the ps.
		badLeafIsPresumptiveSuccessor, _ := badLeaf.IsPresumptiveSuccessor(ctx, tx)
		goodLeafIsPresumptiveSuccessor, _ := goodLeaf.IsPresumptiveSuccessor(ctx, tx)
		require.Equal(t, true, badLeafIsPresumptiveSuccessor)
		require.Equal(t, false, goodLeafIsPresumptiveSuccessor)

		// Next, only the vertex that is not the presumptive successor can start a bisection move.
		bisectionHeight, err := goodLeaf.(*ChallengeVertex).requiredBisectionHeight(ctx, tx)
		require.NoError(t, err)
		require.Equal(t, expectedBisectionHeight, bisectionHeight)

		proof := util.GeneratePrefixProof(lo, loExp, wrongBlockHashes[lo:6])
		_, err = badLeaf.Bisect(
			ctx,
			tx,
			util.HistoryCommitment{
				Height: lo,
				Merkle: loExp.Root(),
			},
			proof,
			staker1,
		)
		require.ErrorIs(t, err, ErrWrongState)

		// Generate a prefix proof for the associated history commitments from the bisection
		// height up to the height of the state commitment for the non-presumptive challenge leaf.
		loExp = util.ExpansionFromLeaves(correctBlockHashes[:lo])
		proof = util.GeneratePrefixProof(lo, loExp, correctBlockHashes[lo:hi])
		bisection, err := goodLeaf.Bisect(
			ctx,
			tx,
			util.HistoryCommitment{
				Height: lo,
				Merkle: loExp.Root(),
			},
			proof,
			staker2,
		)
		require.NoError(t, err)

		// Ensure the prev value of cl2 is set to the vertex we just bisected to.
		goodLeafPrev, _ := goodLeaf.GetPrev(ctx, tx)
		require.Equal(t, bisection, goodLeafPrev.Unwrap())

		// The rootAssertion of the bisectoin should be the rootVertex of this challenge and the bisection
		// should be the new presumptive successor.
		bisectionPrev, _ := bisection.GetPrev(ctx, tx)
		bisectionPrevCommitment, _ := bisectionPrev.Unwrap().GetCommitment(ctx, tx)
		require.Equal(t, challenge.(*Challenge).rootVertex.Unwrap().(*ChallengeVertex).Commitment.Merkle, bisectionPrevCommitment.Merkle)
		bisectionPrevPresumptiveSuccessor, _ := bisectionPrev.Unwrap().IsPresumptiveSuccessor(ctx, tx)
		require.Equal(t, true, bisectionPrevPresumptiveSuccessor)
		return nil
	})

	require.NoError(t, err)
}

func TestAssertionChain_Merge(t *testing.T) {
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	ctx := context.Background()
	t.Run("past deadline", func(t *testing.T) {
		timeRef := util.NewArtificialTimeReference()
		counter := util.NewCountUpTimer(timeRef)
		counter.Add(2 * time.Minute)
		rootAssertion := util.Some(&Assertion{
			challengeManager: &AssertionChain{
				challengePeriod: time.Minute,
			},
		})
		ps := util.Some(ChallengeVertexInterface(&ChallengeVertex{
			PsTimer: counter,
			Commitment: util.HistoryCommitment{
				Height: 1,
			},
		}))
		mergingTo := &ChallengeVertex{
			Challenge: util.Some(ChallengeInterface(&Challenge{
				rootAssertion: rootAssertion,
			})),
			PresumptiveSuccessor: ps,
		}
		mergingFrom := &ChallengeVertex{}
		err := mergingFrom.Merge(
			ctx,
			tx,
			mergingTo,
			[]common.Hash{},
			common.Address{},
		)
		require.ErrorIs(t, err, ErrPastDeadline)
	})
	t.Run("invalid bisection point", func(t *testing.T) {
		mergingTo := &ChallengeVertex{}
		mergingFrom := &ChallengeVertex{
			Prev: util.Some(ChallengeVertexInterface(&ChallengeVertex{
				Commitment: util.HistoryCommitment{
					Height: 3,
				},
			})),
			Commitment: util.HistoryCommitment{
				Height: 4,
			},
		}
		err := mergingFrom.Merge(
			ctx,
			tx,
			mergingTo,
			[]common.Hash{},
			common.Address{},
		)
		require.ErrorIs(t, err, util.ErrUnableToBisect)
	})
	t.Run("invalid height", func(t *testing.T) {
		mergingTo := &ChallengeVertex{
			Commitment: util.HistoryCommitment{
				Height: 2,
			},
		}
		mergingFrom := &ChallengeVertex{
			Prev: util.Some[ChallengeVertexInterface](&ChallengeVertex{
				Commitment: util.HistoryCommitment{
					Height: 2,
				},
			}),
			Commitment: util.HistoryCommitment{
				Height: 4,
			},
		}
		err := mergingFrom.Merge(
			ctx,
			tx,
			mergingTo,
			[]common.Hash{},
			common.Address{},
		)
		require.ErrorIs(t, err, ErrInvalidHeight)
	})
	t.Run("invalid prefix proof", func(t *testing.T) {
		mergingTo := &ChallengeVertex{
			Commitment: util.HistoryCommitment{
				Height: 3,
			},
		}
		mergingFrom := &ChallengeVertex{
			Prev: util.Some[ChallengeVertexInterface](&ChallengeVertex{
				Commitment: util.HistoryCommitment{
					Height: 2,
				},
			}),
			Commitment: util.HistoryCommitment{
				Height: 4,
			},
		}
		err := mergingFrom.Merge(
			ctx,
			tx,
			mergingTo,
			[]common.Hash{},
			common.Address{},
		)
		require.ErrorIs(t, err, util.ErrIncorrectProof)
	})
	t.Run("OK", func(t *testing.T) {
		ctx := context.Background()
		timeRef := util.NewArtificialTimeReference()
		counter := util.NewCountUpTimer(timeRef)
		stateRoots := correctBlockHashesForTest(10)

		loExp := util.ExpansionFromLeaves(stateRoots[:3])
		proof := util.GeneratePrefixProof(
			3,
			loExp,
			stateRoots[3:4],
		)

		exp := util.ExpansionFromLeaves(stateRoots[:3])
		mergingToCommit := util.HistoryCommitment{
			Height: 3,
			Merkle: exp.Root(),
		}
		mergingTo := &ChallengeVertex{
			PsTimer:    counter,
			Commitment: mergingToCommit,
		}
		exp = util.ExpansionFromLeaves(stateRoots[:4])
		mergingFromCommit := util.HistoryCommitment{
			Height: 4,
			Merkle: exp.Root(),
		}
		mergingFrom := &ChallengeVertex{
			PsTimer: counter,
			Challenge: util.Some[ChallengeInterface](&Challenge{
				rootAssertion: util.Some[*Assertion](&Assertion{
					challengeManager: &AssertionChain{
						challengesFeed: NewEventFeed[ChallengeEvent](ctx),
					},
				}),
			}),
			Prev: util.Some[ChallengeVertexInterface](&ChallengeVertex{
				Commitment: util.HistoryCommitment{
					Height: 2,
				},
			}),
			Commitment: mergingFromCommit,
		}
		err := mergingFrom.Merge(
			ctx,
			tx,
			mergingTo,
			proof,
			common.Address{},
		)
		require.NoError(t, err)
	})
}

func correctBlockHashesForTest(numBlocks uint64) []common.Hash {
	var ret []common.Hash
	for i := uint64(0); i < numBlocks; i++ {
		ret = append(ret, util.HashForUint(i))
	}
	return ret
}

func wrongBlockHashesForTest(numBlocks uint64) []common.Hash {
	var ret []common.Hash
	for i := uint64(0); i < numBlocks; i++ {
		ret = append(ret, util.HashForUint(71285937102384-i))
	}
	return ret
}

func TestAssertionChain_StakerInsufficientBalance(t *testing.T) {
	ctx := context.Background()
	chain := NewAssertionChain(ctx, util.NewArtificialTimeReference(), testChallengePeriod)
	require.ErrorContains(t, chain.DeductFromBalance(
		&ActiveTx{TxStatus: ReadWriteTxStatus},
		common.BytesToAddress([]byte{1}),
		AssertionStake,
	), "0 < 100000000000000000000: insufficient balance")
}

func TestAssertionChain_ChallengePeriodLength(t *testing.T) {
	ctx := context.Background()
	cp := 123 * time.Second
	tx := &ActiveTx{TxStatus: ReadOnlyTxStatus}
	chain := NewAssertionChain(ctx, util.NewArtificialTimeReference(), cp)
	require.Equal(t, chain.ChallengePeriodLength(tx), cp)
}

func TestAssertionChain_Inbox(t *testing.T) {
	ctx := context.Background()
	chain := NewAssertionChain(ctx, util.NewArtificialTimeReference(), testChallengePeriod)
	require.Equal(t, chain.Inbox().messages, NewInbox(ctx).messages)
}

func TestAssertionChain_RetrieveAssertions(t *testing.T) {
	ctx := context.Background()
	chain := NewAssertionChain(ctx, util.NewArtificialTimeReference(), testChallengePeriod)
	require.Equal(t, chain.Inbox().messages, NewInbox(ctx).messages)
	staker := common.BytesToAddress([]byte{1})
	bigBalance := new(big.Int).Mul(AssertionStake, big.NewInt(1000))
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	chain.SetBalance(tx, staker, bigBalance)
	p := chain.LatestConfirmed(tx)
	a, err := chain.CreateLeaf(tx, p, util.StateCommitment{Height: 1}, staker)
	require.NoError(t, err)
	require.Equal(t, chain.NumAssertions(tx), uint64(2))
	got, err := chain.AssertionBySequenceNum(tx, 0)
	require.NoError(t, err)
	require.Equal(t, got, p)
	got, err = chain.AssertionBySequenceNum(tx, 1)
	require.NoError(t, err)
	require.Equal(t, got, a)
}

func TestAssertionChain_LeafCreationErrors(t *testing.T) {
	ctx := context.Background()
	chain := NewAssertionChainWithChainId(ctx, util.NewArtificialTimeReference(), testChallengePeriod, 0)
	badChain := NewAssertionChainWithChainId(ctx, util.NewArtificialTimeReference(), testChallengePeriod+1, 1)
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	lc := chain.LatestConfirmed(tx)
	_, err := badChain.CreateLeaf(tx, lc, util.StateCommitment{}, common.BytesToAddress([]byte{}))
	require.ErrorIs(t, err, ErrWrongChain)
	_, err = chain.CreateLeaf(tx, lc, util.StateCommitment{}, common.BytesToAddress([]byte{}))
	require.ErrorIs(t, err, ErrInvalidOp)
}

func TestAssertion_ErrWrongState(t *testing.T) {
	ctx := context.Background()
	chain := NewAssertionChain(ctx, util.NewArtificialTimeReference(), testChallengePeriod)
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	a := chain.LatestConfirmed(tx)
	require.ErrorIs(t, a.RejectForPrev(tx), ErrWrongState)
	require.ErrorIs(t, a.RejectForLoss(ctx, tx), ErrWrongState)
	require.ErrorIs(t, a.ConfirmForWin(ctx, tx), ErrWrongState)
}

func TestAssertion_ErrWrongPredecessorState(t *testing.T) {
	ctx := context.Background()
	chain := NewAssertionChain(ctx, util.NewArtificialTimeReference(), testChallengePeriod)
	staker := common.BytesToAddress([]byte{1})
	bigBalance := new(big.Int).Mul(AssertionStake, big.NewInt(1000))
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	chain.SetBalance(tx, staker, bigBalance)
	newA, err := chain.CreateLeaf(tx, chain.LatestConfirmed(tx), util.StateCommitment{Height: 1}, staker)
	require.NoError(t, err)
	require.ErrorIs(t, newA.RejectForPrev(tx), ErrWrongPredecessorState)
	require.ErrorIs(t, newA.ConfirmForWin(ctx, tx), ErrWrongPredecessorState)
}

func TestAssertion_ErrNotYet(t *testing.T) {
	ctx := context.Background()
	chain := NewAssertionChain(ctx, util.NewArtificialTimeReference(), testChallengePeriod)
	staker := common.BytesToAddress([]byte{1})
	bigBalance := new(big.Int).Mul(AssertionStake, big.NewInt(1000))
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	chain.SetBalance(tx, staker, bigBalance)
	newA, err := chain.CreateLeaf(tx, chain.LatestConfirmed(tx), util.StateCommitment{Height: 1}, staker)
	require.NoError(t, err)
	require.ErrorIs(t, newA.ConfirmNoRival(tx), ErrNotYet)
}

func TestAssertion_ErrInvalid(t *testing.T) {
	ctx := context.Background()
	chain := NewAssertionChain(ctx, util.NewArtificialTimeReference(), testChallengePeriod)
	staker := common.BytesToAddress([]byte{1})
	bigBalance := new(big.Int).Mul(AssertionStake, big.NewInt(1000))
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	chain.SetBalance(tx, staker, bigBalance)
	newA, err := chain.CreateLeaf(tx, chain.LatestConfirmed(tx), util.StateCommitment{Height: 1}, staker)
	require.NoError(t, err)
	newA.Prev = util.None[*Assertion]()
	require.ErrorIs(t, newA.RejectForPrev(tx), ErrInvalidOp)
	require.ErrorIs(t, newA.RejectForLoss(ctx, tx), ErrInvalidOp)
	require.ErrorIs(t, newA.ConfirmNoRival(tx), ErrInvalidOp)
	require.ErrorIs(t, newA.ConfirmForWin(ctx, tx), ErrInvalidOp)
}

func TestAssertion_HasConfirmedSibling(t *testing.T) {
	ctx := context.Background()
	c := &Challenge{}
	tx := &ActiveTx{TxStatus: ReadOnlyTxStatus}
	a := util.Some(&Assertion{
		challengeManager: &AssertionChain{
			challengeVerticesByCommitHash: make(map[ChallengeCommitHash]map[VertexCommitHash]ChallengeVertexInterface),
		}})
	c.rootAssertion = a

	parentStateCommitment, _ := c.ParentStateCommitment(ctx, tx)
	h := parentStateCommitment.Hash()
	parent := &ChallengeVertex{}
	c.rootAssertion.Unwrap().challengeManager.(*AssertionChain).challengeVerticesByCommitHash[ChallengeCommitHash(h)] = map[VertexCommitHash]ChallengeVertexInterface{
		VertexCommitHash(h): &ChallengeVertex{SequenceNum: 100, Status: ConfirmedAssertionState, Prev: util.Some(ChallengeVertexInterface(parent))},
	}

	child := &ChallengeVertex{SequenceNum: 101, Prev: util.Some(ChallengeVertexInterface(parent))}

	hasConfirmedSibling, _ := c.HasConfirmedSibling(ctx, tx, child)
	require.True(t, hasConfirmedSibling)
}
