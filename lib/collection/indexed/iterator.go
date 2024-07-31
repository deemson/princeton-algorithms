package indexed

func ToIterator[T any](indexed SizedIndexed[T]) *Iterator[T] {
	return &Iterator[T]{
		indexed: indexed,
		index:   0,
	}
}

type Iterator[T any] struct {
	indexed SizedIndexed[T]
	index   int
}

func (i *Iterator[T]) HasNext() bool {
	return i.index < i.indexed.Size()
}

func (i *Iterator[T]) Next() T {
	item := i.indexed.Get(i.index)
	i.index++
	return item
}
