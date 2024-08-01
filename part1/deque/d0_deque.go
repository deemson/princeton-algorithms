package deque

import (
	"github.com/deemson/princeton-algorithms/lib/collection"
)

type Deque[T any] struct {
	algorithm algorithm[T]
}

func (d Deque[T]) IsEmpty() bool {
	return collection.IsEmpty(d)
}

func (d Deque[T]) Size() int {
	return d.algorithm.Size()
}

func (d Deque[T]) Iterator() collection.Iterator[T] {
	return d.algorithm.Iterator()
}

func (d Deque[T]) ToSlice() []T {
	return collection.ToSlice[T](d)
}

func (d Deque[T]) AddAtIndex(index int, item T) {
	ensureIndexLessOrEqualSize(index, d.Size())
	d.algorithm.AddAtIndex(index, item)
}

func (d Deque[T]) AddFirst(item T) {
	d.AddAtIndex(0, item)
}

func (d Deque[T]) AddLast(item T) {
	d.AddAtIndex(d.Size(), item)
}

func (d Deque[T]) AddAllLast(items ...T) {
	for _, item := range items {
		d.AddLast(item)
	}
}

func (d Deque[T]) RemoveAtIndex(index int) T {
	ensureNotEmpty(d.Size())
	ensureIndexLessThanSize(index, d.Size())
	return d.algorithm.RemoveAtIndex(index)
}

func (d Deque[T]) RemoveFirst() T {
	return d.RemoveAtIndex(0)
}

func (d Deque[T]) RemoveLast() T {
	return d.RemoveAtIndex(d.Size() - 1)
}

func (d Deque[T]) Get(index int) T {
	ensureIndexLessThanSize(index, d.Size())
	return d.algorithm.Get(index)
}

func (d Deque[T]) Set(index int, item T) {
	ensureIndexLessThanSize(index, d.Size())
	d.algorithm.Set(index, item)
}

func (d Deque[T]) Each(f func(item T) bool) bool {
	return collection.Each(d.Iterator(), f)
}
