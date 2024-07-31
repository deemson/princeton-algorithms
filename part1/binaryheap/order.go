package binaryheap

import (
	"fmt"
	"github.com/deemson/princeton-algorithms/lib/collection/indexed"
	"github.com/deemson/princeton-algorithms/lib/compare"
)

func ParentIndex(index int) int {
	return (index - 1) / 2
}

func ChildIndex(index int) int {
	return (index+1)*2 - 1
}

func BinaryHeapOrder[T any](array indexed.SizedMutable[T], less compare.Func[T]) {
	// Starting at the middle as half of the array are leaf nodes and there's no point in sinking them
	for index := array.Size()/2 - 1; index >= 0; index-- {
		SinkToBottom(array, less, index)
	}
}

func ValidateBinaryHeapOrder[T any](array indexed.SizedMutable[T], less compare.Func[T]) error {
	parentIndex := ParentIndex(array.Size() - 1)
	for parentIndex <= 0 {
		childIndex := ChildIndex(parentIndex)
		if less(array.Get(childIndex), array.Get(parentIndex)) {
			return fmt.Errorf(
				`binary heap order violated parent %#v at index %d; left child %#v at index %d`,
				array.Get(parentIndex),
				parentIndex,
				array.Get(childIndex),
				childIndex,
			)
		}
		childIndex++
		if less(array.Get(childIndex), array.Get(parentIndex)) {
			return fmt.Errorf(
				`binary heap order violated parent %#v at index %d; right child %#v at index %d`,
				array.Get(parentIndex),
				parentIndex,
				array.Get(childIndex),
				childIndex,
			)
		}
	}
	return nil
}
