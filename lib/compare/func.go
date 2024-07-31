package compare

type Func[T any] func(item1, item2 T) bool

func Reversed[T any](f Func[T]) Func[T] {
	return func(item1, item2 T) bool {
		return f(item2, item1)
	}
}
