package sorting

import (
	"github.com/deemson/princeton-algorithms/lib/collection/indexed"
	"github.com/deemson/princeton-algorithms/lib/compare"
)

// ShellSort is similar to previous algorithm (InsertionSort).
// InsertionSort moves a single item at a time, thus being slow
// when sorting an array that requires a lot of swaps.
// ShellSort tries to compensate this by gradually sorting the array
// in a similar fashion, but with bigger steps (>1). Step is equal to 1
// at the very end of the algorithm, efficiently becoming
// an InsertionSort when the array is almost sorted.
//
// The worst-case number of compares used by shell sort with 3x+1 sequence is O(N^1.5).
func ShellSort[T any](array indexed.SizedMutable[T], less compare.Func[T]) {
	step := 1
	// Choosing the biggest step < array.Size() from Knuth's shell sort step sequence (3*x + 1).
	for step < array.Size()/3 {
		step = 3*step + 1
	}
	for step >= 1 {
		// Make an insertion sort for all the steps in the sequence.
		// The outer loop is still incremented by 1 each iteration.
		for outerLoopIndex := step; outerLoopIndex < array.Size(); outerLoopIndex++ {
			// However, the inner loop is decremented by 'step'.
			for innerLoopIndex := outerLoopIndex; innerLoopIndex > step-1; innerLoopIndex -= step {
				// We swap (i.e. move items closer to the start of the array)
				// until the item finds it's place (not less than the item to the left - step).
				// When step == 1 and all items have found their place, the array is sorted.
				if !less(array.Get(innerLoopIndex), array.Get(innerLoopIndex-step)) {
					break
				}
				indexed.Swap[T](array, innerLoopIndex, innerLoopIndex-step)
			}
		}
		step /= 3
	}
}
