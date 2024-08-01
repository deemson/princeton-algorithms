package sorting

import (
	"github.com/deemson/princeton-algorithms/part1/binaryheap"
	"github.com/gogolibs/collection"
	"github.com/gogolibs/compare"
)

func HeapSort[T any](array collection.SizedIndexedMutable[T], less compare.Func[T]) {
	reversedLess := compare.Reversed(less)
	// Binary heap order with reversed less so that max (instead of min) element is on top
	binaryheap.Order(array, reversedLess)
	for index := array.Size() - 1; index > 0; index-- {
		// move max element in place (first to the end, second to the end-1, etc)
		collection.SwapIndexes[T](array, 0, index)
		// as binary heap order now might be violated when we moved arbitrary element
		// to the top, we sink it down to the last unsorted index
		binaryheap.SinkFromTop(array, reversedLess, index-1)
	}
}
