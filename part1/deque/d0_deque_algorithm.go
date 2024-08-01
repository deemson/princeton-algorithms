package deque

import (
	"github.com/deemson/princeton-algorithms/lib/collection"
)

type algorithm[T any] interface {
	Size() int
	Get(index int) T
	Set(index int, item T)
	AddAtIndex(index int, item T)
	RemoveAtIndex(index int) T
	Iterator() collection.Iterator[T]
}
