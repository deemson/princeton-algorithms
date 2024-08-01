package priorityqueue

import (
	"github.com/deemson/princeton-algorithms/part1/deque"
	"github.com/gogolibs/collection"
	"github.com/gogolibs/compare"
)

func OrderedArray[T any](less compare.Func[T], capacity int) PriorityQueue[T] {
	return PriorityQueue[T]{
		algorithm: orderedArrayAlgorithm[T]{
			deque: deque.Slice[T](capacity),
			less:  less,
		},
	}
}

type orderedArrayAlgorithm[T any] struct {
	deque deque.Deque[T]
	less  compare.Func[T]
}

func (a orderedArrayAlgorithm[T]) Push(item T) {
	a.deque.AddLast(item)
	otherItems := a.deque.ToSlice()
	for index := a.deque.Size() - 2; index >= 0; index-- {
		if a.less(otherItems[index], item) {
			break
		}
		collection.SwapIndexes[T](a.deque, index, a.deque.Size()-1)
	}
}

func (a orderedArrayAlgorithm[T]) Pop() T {
	return a.deque.RemoveFirst()
}
