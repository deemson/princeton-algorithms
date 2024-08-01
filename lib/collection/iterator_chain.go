package collection

func ChainIterators[T any](iterators ...Iterator[T]) *ChainedIterator[T] {
	nonEmptyCount := 0
	for _, iterator := range iterators {
		if iterator.HasNext() {
			nonEmptyCount++
		}
	}
	nonEmptyIterators := make([]Iterator[T], 0, nonEmptyCount)
	for _, iterator := range iterators {
		if iterator.HasNext() {
			nonEmptyIterators = append(nonEmptyIterators, iterator)
		}
	}
	return &ChainedIterator[T]{
		iterators: nonEmptyIterators,
		index:     0,
	}
}

type ChainedIterator[T any] struct {
	iterators []Iterator[T]
	index     int
}

func (i *ChainedIterator[T]) HasNext() bool {
	for !i.iterators[i.index].HasNext() {
		i.index++
		if i.index == len(i.iterators) {
			return false
		}
	}
	return true
}

func (i *ChainedIterator[T]) Next() T {
	return i.iterators[i.index].Next()
}
