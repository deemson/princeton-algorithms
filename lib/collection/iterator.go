package collection

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

func Each[T any](iterator Iterator[T], f func(item T) bool) {
	for iterator.HasNext() && f(iterator.Next()) {
	}
}
