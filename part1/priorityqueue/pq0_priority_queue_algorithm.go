package priorityqueue

type algorithm[T any] interface {
	Push(item T)
	Pop() T
}
