package sorting

import (
	"github.com/gogolibs/collection"
	"github.com/gogolibs/compare"
)

// MergeSort is a divide and conquer (split-sort-merge) algorithm.
// The execution time is as fast as O(N*log2(N)).
// Traditional implementation via recursion.
func MergeSort[T any](array collection.SizedIndexedMutable[T], less compare.Func[T]) {
	aux := collection.NewSliceAdapter(make([]T, array.Size()))
	mergeSort(array, aux, less, 0, array.Size()-1)
}

// BottomUpMergeSort implements the same merge sort only without recursion
func BottomUpMergeSort[T any](array collection.SizedIndexedMutable[T], less compare.Func[T]) {
	aux := collection.NewSliceAdapter(make([]T, array.Size()))
	// partitionSize size grows as 1 2 4 8 ...
	partitionSize := 1
	for partitionSize < array.Size() {
		for lo := 0; lo < array.Size()-partitionSize; lo += 2 * partitionSize {
			mid := lo + partitionSize - 1
			hi := min(lo+2*partitionSize-1, array.Size()-1)
			merge(array, aux, less, lo, mid, hi)
		}
		partitionSize += partitionSize
	}
}

func mergeSort[T any](array, aux collection.SizedIndexedMutable[T], less compare.Func[T], lo, hi int) {
	// If the thresholds overlap the sorting is done.
	if lo >= hi {
		return
	}
	// Divide and conquer in action: the array is sliced in two parts...
	mid := lo + (hi-lo)/2
	// ...and each part is sorted independently...
	mergeSort(array, aux, less, lo, mid)
	mergeSort(array, aux, less, mid+1, hi)
	// ... and the two parts are merged back together.
	merge(array, aux, less, lo, mid, hi)
}

// merge is a core function in a divide and conquer algorithms of MergeSort and BottomUpMergeSort.
// merge merges two parts of the indexed array together via temporary aux storage array.
// It assumes that both halves of the array are sorted individually, and they need to be merged
// together so that the result is sorted as well.
func merge[T any](array, aux collection.SizedIndexedMutable[T], less compare.Func[T], lo, mid, hi int) {
	for index := lo; index <= hi; index++ {
		aux.Set(index, array.Get(index))
	}
	// leftMarker moves from lo to mid
	leftMarker := lo
	// rightMarker moves from mid+1 to hi
	rightMarker := mid + 1
	for index := lo; index <= hi; index++ {
		switch {
		// When the left marker has crossed the middle point, it means there's nothing left
		// in the left part and the remainder of the right part can just be copied.
		case leftMarker > mid:
			array.Set(index, aux.Get(rightMarker))
			rightMarker++
		// The same goes for the right part when the right marker crosses hi threshold.
		case rightMarker > hi:
			array.Set(index, aux.Get(leftMarker))
			leftMarker++
		// At this point both parts are not exhausted,
		// so the smaller item from either left or right side
		// is copied to 'array' and respective marker is increased.
		case less(aux.Get(rightMarker), aux.Get(leftMarker)):
			array.Set(index, aux.Get(rightMarker))
			rightMarker++
		default: // lessOrEqual(aux.Get(leftMarker), aux.Get(rightMarker))
			array.Set(index, aux.Get(leftMarker))
			leftMarker++
		}
	}
}
