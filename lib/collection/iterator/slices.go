package iterator

func Items[T any](items ...T) *Slice[T] {
	return FromSlice(items)
}

func FromSlice[T any](s []T) *Slice[T] {
	return &Slice[T]{
		slice: s,
		index: 0,
	}
}

type Slice[T any] struct {
	slice []T
	index int
}

func (i *Slice[T]) HasNext() bool {
	return i.index < len(i.slice)
}

func (i *Slice[T]) Next() T {
	item := i.slice[i.index]
	i.index++
	return item
}

func ToSlice[T any](iterator Iterator[T], capacity int) []T {
	slice := make([]T, 0, capacity)
	for iterator.HasNext() {
		slice = append(slice, iterator.Next())
	}
	return slice
}
