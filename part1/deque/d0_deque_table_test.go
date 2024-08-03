package deque_test

import "github.com/deemson/princeton-algorithms/part1/deque"

type testDequeMaker[T any] func() (string, deque.Deque[T])

type testTableBuilder[T any] struct {
	capacity int
}

func (b testTableBuilder[T]) build() []testDequeMaker[T] {
	return []testDequeMaker[T]{
		func() (string, deque.Deque[T]) {
			return "Slice", deque.Slice[T](b.capacity)
		},
		func() (string, deque.Deque[T]) {
			return "LinkedList", deque.LinkedList[T]()
		},
	}
}
