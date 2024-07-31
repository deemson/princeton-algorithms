package sequence

import (
	"github.com/deemson/princeton-algorithms/lib/collection/iterator"
)

func ToSlice[T any](sequence SizedSequence[T]) []T {
	slice := make([]T, 0, sequence.Size())
	iterator.Each(sequence.Iterator(), func(item T) bool {
		slice = append(slice, item)
		return true
	})
	return slice
}
