package sorting

import (
	"github.com/deemson/princeton-algorithms/lib/collection"
	"github.com/deemson/princeton-algorithms/lib/compare"
	"math/rand"
)

func QuickSort[T any](array collection.SizedIndexedMutable[T], less compare.Func[T]) {
	shuffle(array)
	quickSort(array, less, 0, array.Size()-1)
}

func shuffle[T any](array collection.SizedIndexedMutable[T]) {
	for index := 0; index < array.Size(); index++ {
		shuffledIndex := rand.Intn(index + 1)
		collection.Swap[T](array, index, shuffledIndex)
	}
}

func quickSort[T any](array collection.SizedIndexedMutable[T], less compare.Func[T], left, right int) {
	if left >= right {
		return
	}
	itemInPlaceIndex := partition(array, less, left, right)
	quickSort(array, less, left, itemInPlaceIndex-1)
	quickSort(array, less, itemInPlaceIndex+1, right)
}

// partition is the core of QuickSort. It ensures that randomly picked item from the array
// is put in such a position, so that all the items to the left are smaller and all items to the right
// are bigger. Function partition returns the index of the element after this is done.
func partition[T any](array collection.SizedIndexedMutable[T], less compare.Func[T], left, right int) int {
	partitioningItem := array.Get(left)
	leftMarker := left + 1
	rightMarker := right
	for {
		// move left marker as long as items are less than partitioning item
		// stop at first item that is not
		for less(array.Get(leftMarker), partitioningItem) {
			if leftMarker == right {
				break
			}
			leftMarker++
		}
		// move right marker as long as items are not less than partitioning item
		// stop at first item that is
		for !less(array.Get(rightMarker), partitioningItem) {
			if rightMarker == left {
				break
			}
			rightMarker--
		}
		// check if markers overlap which means the pair to swap was not found
		if leftMarker >= rightMarker {
			break
		}
		// swap the found pair, repeat until condition above is met
		collection.Swap[T](array, leftMarker, rightMarker)
	}
	// put partitioned element in place
	collection.Swap[T](array, left, rightMarker)
	// return it's current index
	return rightMarker
}
