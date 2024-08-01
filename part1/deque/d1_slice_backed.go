package deque

import (
	"github.com/deemson/princeton-algorithms/lib/collection"
)

func makeSlice(capacity int) []any {
	slice := make([]any, capacity)
	return slice
}

func Slice[T any](capacity int) Deque[T] {
	return Deque[T]{
		algorithm: &sliceAlgorithm[T]{
			slice:          makeSlice(capacity),
			firstItemIndex: 0,
			lastItemIndex:  0,
			size:           0,
		},
	}
}

type sliceAlgorithm[T any] struct {
	slice          []any
	firstItemIndex int
	lastItemIndex  int
	size           int
}

func (a *sliceAlgorithm[T]) Size() int {
	return a.size
}

func (a *sliceAlgorithm[T]) Get(index int) T {
	return a.slice[a.normalizeIndex(index)].(T)
}

func (a *sliceAlgorithm[T]) Set(index int, item T) {
	a.slice[a.normalizeIndex(index)] = item
}

func (a *sliceAlgorithm[T]) AddAtIndex(index int, item T) {
	a.growIfRequired()
	// move items out of the way either at the end or at the beginning -- whichever requires fewer items to move
	if 2*index < a.size {
		// fewer items to move at the start
		a.firstItemIndex--
		if a.firstItemIndex < 0 {
			a.firstItemIndex = a.capacity() - 1
		}
		for shiftIndex := 0; shiftIndex < index; shiftIndex++ {
			a.Set(shiftIndex, a.Get(shiftIndex+1))
		}
	} else {
		// fewer items to move at the end
		a.lastItemIndex++
		if a.lastItemIndex == a.capacity() {
			a.lastItemIndex = 0
		}
		for shiftIndex := a.size; shiftIndex > index; shiftIndex-- {
			a.Set(shiftIndex, a.Get(shiftIndex-1))
		}
	}
	a.Set(index, item)
	a.size++
}

func (a *sliceAlgorithm[T]) RemoveAtIndex(index int) T {
	var item T
	// move item at index to first or last position -- whichever requires fewer swaps and remove it
	if 2*index < a.size-1 {
		// fewer items to move at the start
		for swapIndex := index; swapIndex > 0; swapIndex-- {
			a.swap(swapIndex, swapIndex-1)
		}
		item = a.Get(0)
		a.slice[a.normalizeIndex(0)] = nil
		a.firstItemIndex++
		if a.firstItemIndex == a.capacity() {
			a.firstItemIndex = 0
		}
	} else {
		// fewer items to move at the end
		for swapIndex := index; swapIndex < a.size-1; swapIndex++ {
			a.swap(swapIndex, swapIndex+1)
		}
		item = a.Get(a.size - 1)
		a.slice[a.normalizeIndex(a.size-1)] = nil
		a.lastItemIndex--
		if a.lastItemIndex < 0 {
			a.lastItemIndex = a.capacity() - 1
		}
	}
	a.size--
	a.shrinkIfRequired()
	return item
}

func (a *sliceAlgorithm[T]) Iterator() collection.Iterator[T] {
	return collection.ToIterator[T](a)
}

// normalizeIndex makes sure that index that comes from the outer code which
// is going to be in range [0..size-1] to be in range [firstItemIndex..lastItemIndex-1]
// that sliceAlgorithm understands
func (a *sliceAlgorithm[T]) normalizeIndex(index int) int {
	return (a.firstItemIndex + index) % a.capacity()
}

func (a *sliceAlgorithm[T]) capacity() int {
	return len(a.slice)
}

// growIfRequired resizes the slice to twice capacity when it's full.
func (a *sliceAlgorithm[T]) growIfRequired() {
	if a.size == a.capacity() {
		a.resize(a.capacity() * 2)
	}
}

// shrinkIfRequired resizes the slice to half the size when it's quarter full.
// The resize is done at quarter capacity to avoid "thrashing" (constant resizing)
// when working with half-full slice and doing add-remove operations.
func (a *sliceAlgorithm[T]) shrinkIfRequired() {
	if a.size > 0 && a.size == a.capacity()/4 {
		a.resize(a.capacity() / 2)
	}
}

// resize resizes the slice to have len == capacity.
func (a *sliceAlgorithm[T]) resize(capacity int) {
	slice := makeSlice(capacity)
	for index := 0; index < a.size; index++ {
		slice[index] = a.Get(index)
	}
	a.slice = slice
	a.firstItemIndex = 0
	a.lastItemIndex = a.size
}

func (a *sliceAlgorithm[T]) swap(index1, index2 int) {
	collection.Swap[T](a, index1, index2)
}
