package binaryheap

import (
	"github.com/deemson/princeton-algorithms/lib/collection/indexed"
	"github.com/deemson/princeton-algorithms/lib/compare"
)

func Sink[T any](array indexed.SizedMutable[T], less compare.Func[T], fromIndex, toIndex int) {
	parentIndex := fromIndex
	childIndex := ChildIndex(parentIndex)
	for childIndex <= toIndex {
		if childIndex < toIndex && less(array.Get(childIndex+1), array.Get(childIndex)) {
			childIndex++
		}
		if less(array.Get(parentIndex), array.Get(childIndex)) {
			break
		}
		indexed.Swap[T](array, parentIndex, childIndex)
		parentIndex = childIndex
		childIndex = ChildIndex(parentIndex)
	}
}

func SinkFromTop[T any](array indexed.SizedMutable[T], less compare.Func[T], to int) {
	Sink(array, less, 0, to)
}

func SinkToBottom[T any](array indexed.SizedMutable[T], less compare.Func[T], fromIndex int) {
	Sink(array, less, fromIndex, array.Size()-1)
}

func SinkFromTopToBottom[T any](array indexed.SizedMutable[T], less compare.Func[T]) {
	SinkToBottom(array, less, 0)
}
