package symboltable

import (
	"github.com/deemson/princeton-algorithms/lib/collection"
	"github.com/deemson/princeton-algorithms/lib/compare"
	"sort"
)

type Set[T any] struct {
	algorithm algorithm[T, struct{}]
}

func (s Set[T]) Size() int {
	return s.algorithm.Size()
}

func (s Set[T]) Iterator() collection.Iterator[T] {
	return collection.TransformIterator(s.algorithm.Iterator(), func(pair Pair[T, struct{}]) T {
		return pair.Key
	})
}

func (s Set[T]) ToSlice() []T {
	return collection.ToSlice[T](s)
}

func (s Set[T]) ToSortedSlice(less compare.Func[T]) []T {
	slice := s.ToSlice()
	sort.Slice(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
	return slice
}
