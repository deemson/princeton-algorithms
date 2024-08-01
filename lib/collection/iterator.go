package collection

func IteratorFromItems[T any](items ...T) *SliceIterator[T] {
	return IteratorFromSlice(items)
}

func IteratorFromSlice[T any](s []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		slice: s,
		index: 0,
	}
}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

func Each[T any](iterator Iterator[T], f func(item T) bool) {
	for iterator.HasNext() && f(iterator.Next()) {
	}
}

type SliceIterator[T any] struct {
	slice []T
	index int
}

func (i *SliceIterator[T]) HasNext() bool {
	return i.index < len(i.slice)
}

func (i *SliceIterator[T]) Next() T {
	item := i.slice[i.index]
	i.index++
	return item
}

func IteratorToSlice[T any](iterator Iterator[T], capacity int) []T {
	slice := make([]T, 0, capacity)
	for iterator.HasNext() {
		slice = append(slice, iterator.Next())
	}
	return slice
}
