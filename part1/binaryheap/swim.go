package binaryheap

import (
	"github.com/deemson/princeton-algorithms/lib/collection/indexed"
	"github.com/deemson/princeton-algorithms/lib/compare"
)

func Swim[T any](array indexed.SizedMutable[T], less compare.Func[T], fromIndex, toIndex int) {
	childIndex := fromIndex
	parentIndex := ParentIndex(childIndex)
	for parentIndex >= toIndex && less(array.Get(childIndex), array.Get(parentIndex)) {
		indexed.Swap[T](array, childIndex, parentIndex)
		childIndex = parentIndex
		parentIndex = ParentIndex(childIndex)
	}
}

func SwimToTop[T any](array indexed.SizedMutable[T], less compare.Func[T], fromIndex int) {
	Swim(array, less, fromIndex, 0)
}

func SwimFromBottom[T any](array indexed.SizedMutable[T], less compare.Func[T], toIndex int) {
	Swim(array, less, array.Size()-1, toIndex)
}

func SwimFromBottomToTop[T any](array indexed.SizedMutable[T], less compare.Func[T]) {
	SwimFromBottom(array, less, 0)
}