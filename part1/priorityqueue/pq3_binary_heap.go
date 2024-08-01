package priorityqueue

import (
	"github.com/deemson/princeton-algorithms/part1/binaryheap"
	"github.com/deemson/princeton-algorithms/part1/deque"
	"github.com/gogolibs/collection"
	"github.com/gogolibs/compare"
)

func BinaryHeap[T any](less compare.Func[T], capacity int) PriorityQueue[T] {
	return PriorityQueue[T]{
		algorithm: binaryHeapAlgorithm[T]{
			deque: deque.Slice[T](capacity),
			less:  less,
		},
	}
}

type binaryHeapAlgorithm[T any] struct {
	deque deque.Deque[T]
	less  compare.Func[T]
}

func (a binaryHeapAlgorithm[T]) Push(item T) {
	a.deque.AddLast(item)
	if a.deque.Size() > 1 {
		binaryheap.SwimFromBottomToTop(a.deque, a.less)
	}
}

func (a binaryHeapAlgorithm[T]) Pop() T {
	collection.SwapIndexes[T](a.deque, 0, a.deque.Size()-1)
	item := a.deque.RemoveLast()
	if a.deque.Size() > 1 {
		binaryheap.SinkFromTopToBottom(a.deque, a.less)
	}
	return item
}
