package sorting

import (
	"github.com/deemson/princeton-algorithms/lib/collection/indexed"
	"github.com/deemson/princeton-algorithms/lib/compare"
)

// SelectionSort does ~N^2/2 compares and N swaps
func SelectionSort[T any](array indexed.SizedMutable[T], less compare.Func[T]) {
	// Outer loop that iterates through all the items with the intention to swap at the end.
	for outerLoopIndex := 0; outerLoopIndex < array.Size(); outerLoopIndex++ {
		// Each iteration of the outer loop it tries to find the minimum item index to the right
		// of the current index and swap this item with the current item.
		mininumItemIndex := outerLoopIndex
		// Inner loop tries to find the minimum item to the right of the current one.
		// This is why it starts at the current index of the outer loop.
		for innerLoopIndex := outerLoopIndex; innerLoopIndex < array.Size(); innerLoopIndex++ {
			if less(array.Get(innerLoopIndex), array.Get(mininumItemIndex)) {
				mininumItemIndex = innerLoopIndex
			}
		}
		indexed.Swap[T](array, outerLoopIndex, mininumItemIndex)
	}
}
