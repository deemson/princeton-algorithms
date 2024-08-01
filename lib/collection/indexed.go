package collection

type Indexed[T any] interface {
	Get(index int) T
}

type SizedIndexed[T any] interface {
	Sized
	Indexed[T]
}

func ToIterator[T any](indexed SizedIndexed[T]) *IteratorByIndex[T] {
	return &IteratorByIndex[T]{
		indexed: indexed,
		index:   0,
	}
}

type IteratorByIndex[T any] struct {
	indexed SizedIndexed[T]
	index   int
}

func (i *IteratorByIndex[T]) HasNext() bool {
	return i.index < i.indexed.Size()
}

func (i *IteratorByIndex[T]) Next() T {
	item := i.indexed.Get(i.index)
	i.index++
	return item
}

type IndexedMutable[T any] interface {
	Indexed[T]
	Set(index int, item T)
}

func Swap[T any](indexed IndexedMutable[T], index1, index2 int) {
	tmp := indexed.Get(index1)
	indexed.Set(index1, indexed.Get(index2))
	indexed.Set(index2, tmp)
}

type SizedIndexedMutable[T any] interface {
	Sized
	IndexedMutable[T]
}
