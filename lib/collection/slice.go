package collection

func Slice[T any](slice []T) *SliceAdapter[T] {
	return &SliceAdapter[T]{
		slice: slice,
	}
}

type SliceAdapter[T any] struct {
	slice []T
}

func (a *SliceAdapter[T]) Size() int {
	return len(a.slice)
}

func (a *SliceAdapter[T]) Get(index int) T {
	return a.slice[index]
}

func (a *SliceAdapter[T]) Set(index int, item T) {
	a.slice[index] = item
}

func (a *SliceAdapter[T]) Iterator() Iterator[T] {
	return IteratorFromSlice(a.slice)
}
