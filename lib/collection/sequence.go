package collection

type Sequence[T any] interface {
	Iterator() Iterator[T]
}

type SizedSequence[T any] interface {
	Sized
	Sequence[T]
}

func ToSlice[T any](sequence SizedSequence[T]) []T {
	slice := make([]T, 0, sequence.Size())
	Each(sequence.Iterator(), func(item T) bool {
		slice = append(slice, item)
		return true
	})
	return slice
}
