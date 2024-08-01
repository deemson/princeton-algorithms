package binaryheap

import (
	"github.com/gogolibs/collection"
	"github.com/gogolibs/compare"
)

func Swim[T any](array collection.SizedIndexedMutable[T], less compare.Func[T], fromIndex, toIndex int) {
	childIndex := fromIndex
	parentIndex := ParentIndex(childIndex)
	for parentIndex >= toIndex && less(array.Get(childIndex), array.Get(parentIndex)) {
		collection.SwapIndexes[T](array, childIndex, parentIndex)
		childIndex = parentIndex
		parentIndex = ParentIndex(childIndex)
	}
}

func SwimToTop[T any](array collection.SizedIndexedMutable[T], less compare.Func[T], fromIndex int) {
	Swim(array, less, fromIndex, 0)
}

func SwimFromBottom[T any](array collection.SizedIndexedMutable[T], less compare.Func[T], toIndex int) {
	Swim(array, less, array.Size()-1, toIndex)
}

func SwimFromBottomToTop[T any](array collection.SizedIndexedMutable[T], less compare.Func[T]) {
	SwimFromBottom(array, less, 0)
}
