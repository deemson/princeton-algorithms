package iterator

func Each[T any](iterator Iterator[T], f func(item T) bool) {
	for iterator.HasNext() && f(iterator.Next()) {
	}
}
