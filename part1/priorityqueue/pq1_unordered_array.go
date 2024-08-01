package priorityqueue

import (
	"github.com/deemson/princeton-algorithms/part1/deque"
	"github.com/gogolibs/collection"
	"github.com/gogolibs/compare"
)

func UnorderedArray[T any](less compare.Func[T], capacity int) PriorityQueue[T] {
	return PriorityQueue[T]{
		algorithm: unorderedArrayAlgorithm[T]{
			deque: deque.Slice[T](capacity),
			less:  less,
		},
	}
}

// unorderedArrayAlgorithm scales as O(1) for inserts and as O(N) for deletes
type unorderedArrayAlgorithm[T any] struct {
	deque deque.Deque[T]
	less  compare.Func[T]
}

func (a unorderedArrayAlgorithm[T]) Push(item T) {
	a.deque.AddLast(item)
}

func (a unorderedArrayAlgorithm[T]) Pop() T {
	minIndex := 0
	minElement := a.deque.Get(0)
	collection.ForIndexedEach(a.deque, func(index int, item T) bool {
		if index == 0 {
			return true
		}
		if a.less(item, minElement) {
			minIndex = index
			minElement = item
		}
		return true
	})
	collection.SwapIndexes[T](a.deque, minIndex, a.deque.Size()-1)
	return a.deque.RemoveLast()
}
