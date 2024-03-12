package threadsafe

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	require.NotNil(t, q, "Queue should not be nil")
}

func TestQueueLen(t *testing.T) {
	q := NewQueue[int]()
	q.Push(1)
	q.Push(2)
	require.Equal(t, 2, q.Len())
}

func TestQueuePush(t *testing.T) {
	q := NewQueue[int]()
	q.Push(1)
	require.Equal(t, 1, q.Len())
}

func TestQueuePop(t *testing.T) {
	q := NewQueue[int]()
	q.Push(1)
	q.Push(2)

	val := q.Pop()
	require.True(t, val.IsSome())
	require.Equal(t, 1, val.Unwrap())

	require.Equal(t, 1, q.Len())

	// Test pop on empty queue
	_ = q.Pop() // Pop remaining element
	opt := q.Pop()
	require.False(t, opt.IsNone())
}

func TestQueueConcurrency(t *testing.T) {
	q := NewQueue[int]()
	var wg sync.WaitGroup

	// Concurrent Push
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			q.Push(val)
		}(i)
	}
	wg.Wait()

	require.Equal(t, 1000, q.Len())

	// Concurrent Pop
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			q.Pop()
		}()
	}
	wg.Wait()

	require.Equal(t, 500, q.Len())
}
