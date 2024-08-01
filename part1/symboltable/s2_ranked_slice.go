package symboltable

import (
	"github.com/deemson/princeton-algorithms/lib/collection"
	"github.com/deemson/princeton-algorithms/lib/compare"
	"github.com/deemson/princeton-algorithms/part1/deque"
)

func RankedSlice[K, V any](less, equal compare.Func[K], capacity int) SymbolTable[K, V] {
	return SymbolTable[K, V]{
		algorithm: rankedSliceAlgorithm[K, V]{
			deque: deque.Slice[Pair[K, V]](capacity),
			less:  less,
			equal: equal,
		},
	}
}

type rankedSliceAlgorithm[K, V any] struct {
	deque deque.Deque[Pair[K, V]]
	less  compare.Func[K]
	equal compare.Func[K]
}

func (a rankedSliceAlgorithm[K, V]) Size() int {
	return a.deque.Size()
}

func (a rankedSliceAlgorithm[K, V]) Get(key K) (V, bool) {
	var emptyValue V
	rank := a.rank(key)
	if rank == a.Size() {
		return emptyValue, false
	}
	pair := a.deque.Get(rank)
	if a.equal(pair.Key, key) {
		return pair.Value, true
	}
	return emptyValue, false
}

func (a rankedSliceAlgorithm[K, V]) Set(key K, value V) {
	rank := a.rank(key)
	pair := Pair[K, V]{
		Key:   key,
		Value: value,
	}
	switch {
	case rank == a.Size():
		a.deque.AddLast(pair)
	case a.equal(a.deque.Get(rank).Key, key):
		a.deque.Set(rank, pair)
	default:
		a.deque.AddAtIndex(rank, pair)
	}
}

func (a rankedSliceAlgorithm[K, V]) Delete(key K) bool {
	rank := a.rank(key)
	if rank == a.Size() {
		return false
	}
	pair := a.deque.Get(rank)
	if a.equal(pair.Key, key) {
		a.deque.RemoveAtIndex(rank)
		return true
	}
	return false
}

func (a rankedSliceAlgorithm[K, V]) Iterator() collection.Iterator[Pair[K, V]] {
	return a.deque.Iterator()
}

func (a rankedSliceAlgorithm[K, V]) keys() Set[K] {
	d := deque.Slice[Pair[K, struct{}]](a.Size() * 2)
	a.deque.Each(func(pair Pair[K, V]) bool {
		d.AddLast(Pair[K, struct{}]{
			Key:   pair.Key,
			Value: struct{}{},
		})
		return true
	})
	return Set[K]{
		algorithm: rankedSliceAlgorithm[K, struct{}]{
			deque: d,
			less:  a.less,
			equal: a.equal,
		},
	}
}

// rank returns the number of keys in this symbol table strictly less than key
func (a rankedSliceAlgorithm[K, V]) rank(key K) int {
	lo := 0
	hi := a.deque.Size() - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case a.less(key, a.deque.Get(mid).Key):
			hi = mid - 1
		case a.less(a.deque.Get(mid).Key, key):
			lo = mid + 1
		default:
			return mid
		}
	}
	return lo
}
