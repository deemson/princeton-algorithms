package sorting

import (
	"github.com/gogolibs/collection"
	"github.com/gogolibs/compare"
)

// InsertionSort does on average ~1/4N^2 compares and ~1/4N^2 swaps.
// With partially sorted arrays, though, it's performance can be close to linear.
func InsertionSort[T any](array collection.SizedIndexedMutable[T], less compare.Func[T]) {
	for outerLoopIndex := 1; outerLoopIndex < array.Size(); outerLoopIndex++ {
		// Every added item tries to find its place in the left part (sorted part) of the entire array.
		for innerLoopIndex := outerLoopIndex; innerLoopIndex > 0; innerLoopIndex-- {
			if less(array.Get(innerLoopIndex), array.Get(innerLoopIndex-1)) {
				collection.SwapIndexes[T](array, innerLoopIndex, innerLoopIndex-1)
			}
		}
	}
}
