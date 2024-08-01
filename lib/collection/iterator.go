package collection

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

// Each iterates through all items of the Iterator and accepts a function
// that receives the current item and returns a flag signalling whether
// to continue iteration or not. Each itself returns a flag signalling whether
// the iterator has been exhausted (true) or not (false)
func Each[T any](iterator Iterator[T], f func(item T) bool) bool {
	for iterator.HasNext() {
		if !f(iterator.Next()) {
			return false
		}
	}
	return true
}
