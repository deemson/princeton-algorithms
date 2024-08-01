package priorityqueue

type PriorityQueue[T any] struct {
	algorithm algorithm[T]
}

func (q PriorityQueue[T]) Push(item T) {
	q.algorithm.Push(item)
}

func (q PriorityQueue[T]) PushMany(items ...T) {
	for _, item := range items {
		q.Push(item)
	}
}

func (q PriorityQueue[T]) Pop() T {
	return q.algorithm.Pop()
}

func (q PriorityQueue[T]) PopMany(number int) []T {
	slice := make([]T, number)
	for index := 0; index < number; index++ {
		slice[index] = q.Pop()
	}
	return slice
}
