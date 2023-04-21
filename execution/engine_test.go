package execution

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestExecutionEngine(t *testing.T) {
	machine := NewSimpleMachine(big.NewInt(123))
	for i := uint64(0); i < 100; i++ {
		require.Equal(t, i, machine.CurrentStepNum())

		thisHash := machine.Hash()
		stopped := machine.IsStopped()
		osp, err := machine.OneStepProof()
		require.NoError(t, err)
		err = machine.Step(1)
		require.NoError(t, err)
		nextHash := machine.Hash()

		require.Equal(t, thisHash == nextHash, stopped)

		if !VerifySimpleMachineOneStepProof(thisHash, nextHash, i, osp) {
			t.Fatal(i)
		}

		// verify that bad proofs get rejected
		fakeProof := append([]byte{}, osp...)
		fakeProof = append(fakeProof, 0)
		if VerifySimpleMachineOneStepProof(thisHash, nextHash, i, fakeProof) {
			t.Fatal(i)
		}

		fakeProof = append([]byte{}, osp...)
		fakeProof[19] ^= 1
		if VerifySimpleMachineOneStepProof(thisHash, nextHash, i, fakeProof) {
			t.Fatal(i)
		}

		fakeProof = append([]byte{}, osp...)
		fakeProof = fakeProof[len(fakeProof)-1:]
		if VerifySimpleMachineOneStepProof(thisHash, nextHash, i, fakeProof) {
			t.Fatal(i)
		}

		if thisHash != nextHash && VerifySimpleMachineOneStepProof(thisHash, thisHash, i, osp) {
			t.Fatal(i)
		}
		if VerifySimpleMachineOneStepProof(thisHash, common.Hash{}, i, osp) {
			t.Fatal(i)
		}
	}
}
