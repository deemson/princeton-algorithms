package sorting

import (
	"github.com/deemson/princeton-algorithms/lib/collection/indexed"
	"github.com/deemson/princeton-algorithms/lib/compare"
	"github.com/deemson/princeton-algorithms/part1/binaryheap"
)

func HeapSort[T any](array indexed.SizedMutable[T], less compare.Func[T]) {
	reversedLess := compare.Reversed(less)
	// Binary heap order with reversed less so that max (instead of min) element is on top
	binaryheap.BinaryHeapOrder(array, reversedLess)
	for index := array.Size() - 1; index > 0; index-- {
		// move max element in place (first to the end, second to the end-1, etc)
		indexed.Swap[T](array, 0, index)
		// as binary heap order now might be violated when we moved arbitrary element
		// to the top, we sink it down to the last unsorted index
		binaryheap.SinkFromTop(array, reversedLess, index-1)
	}
}
