package threadsafe

import (
	"github.com/OffchainLabs/bold/containers/option"
)

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}

func (q *Queue[T]) Push(t T) {
	q.data = append(q.data, t)
}

func (q *Queue[T]) Pop() option.Option[T] {
	if q.Len() == 0 {
		return option.None[T]()
	}
	item := q.data[0]
	q.data = q.data[1:]
	return option.Some(item)
}
