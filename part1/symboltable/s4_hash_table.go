package symboltable

import (
	"github.com/deemson/princeton-algorithms/lib/hash"
	"github.com/deemson/princeton-algorithms/part1/deque"
	"github.com/gogolibs/compare"
	"github.com/gogolibs/iterator"
)

func makeSliceOfDeques[K, V any](size int) []deque.Deque[Pair[K, V]] {
	slice := make([]deque.Deque[Pair[K, V]], size)
	for index := range slice {
		slice[index] = deque.LinkedList[Pair[K, V]]()
	}
	return slice
}

func HashTable[K, V any](hash hash.Func[K], equal compare.Func[K], capacity int, loadFactor float32) SymbolTable[K, V] {
	return SymbolTable[K, V]{
		algorithm: &hashTableAlgorithm[K, V]{
			size:                 0,
			hash:                 hash,
			occupiedSliceIndexes: 0,
			slice:                makeSliceOfDeques[K, V](capacity),
			equal:                equal,
			loadFactor:           loadFactor,
		},
	}
}

type hashTableAlgorithm[K, V any] struct {
	size                 int
	hash                 hash.Func[K]
	occupiedSliceIndexes int
	slice                []deque.Deque[Pair[K, V]]
	equal                compare.Func[K]
	loadFactor           float32
}

func (a *hashTableAlgorithm[K, V]) Size() int {
	return a.size
}

func (a *hashTableAlgorithm[K, V]) Iterator() iterator.Iterator[Pair[K, V]] {
	iterators := make([]iterator.Iterator[Pair[K, V]], len(a.slice))
	for index, d := range a.slice {
		iterators[index] = d.Iterator()
	}
	return iterator.Chain(iterators...)
}

func (a *hashTableAlgorithm[K, V]) Get(key K) (V, bool) {
	var value V
	isFound := false
	a.slice[a.sliceIndexForKey(key)].Each(func(pair Pair[K, V]) bool {
		if a.equal(pair.Key, key) {
			isFound = true
			value = pair.Value
			return false
		}
		return true
	})
	return value, isFound
}

func (a *hashTableAlgorithm[K, V]) Set(key K, value V) {
	sizeBeforeSet := a.size
	a.set(Pair[K, V]{
		Key:   key,
		Value: value,
	})
	if a.size > sizeBeforeSet {
		a.growIfRequired()
	}
}

func (a *hashTableAlgorithm[K, V]) Delete(key K) bool {
	d := a.slice[a.sliceIndexForKey(key)]
	index := 0
	return !d.Each(func(pair Pair[K, V]) bool {
		if a.equal(pair.Key, key) {
			d.RemoveAtIndex(index)
			a.size--
			if d.IsEmpty() {
				a.occupiedSliceIndexes--
			}
			a.shrinkIfRequired()
			return false
		}
		index++
		return true
	})
}

func (a *hashTableAlgorithm[K, V]) set(pair Pair[K, V]) {
	d := a.slice[a.sliceIndexForKey(pair.Key)]
	if d.IsEmpty() {
		d.AddLast(pair)
		a.occupiedSliceIndexes++
		a.size++
		return
	}
	index := 0
	isNotFound := d.Each(func(existingPair Pair[K, V]) bool {
		if a.equal(existingPair.Key, pair.Key) {
			d.Set(index, pair)
			return false
		}
		index++
		return true
	})
	if isNotFound {
		d.AddLast(pair)
	}
	a.size++
}

func (a *hashTableAlgorithm[K, V]) load() float32 {
	return float32(a.occupiedSliceIndexes) / float32(len(a.slice))
}

func (a *hashTableAlgorithm[K, V]) sliceIndexForKey(key K) int {
	return int(a.hash(key) % uint32(len(a.slice)))
}

func (a *hashTableAlgorithm[K, V]) resize(capacity int) {
	oldSlice := a.slice
	a.slice = makeSliceOfDeques[K, V](capacity)
	a.occupiedSliceIndexes = 0
	a.size = 0
	for _, d := range oldSlice {
		d.Each(func(pair Pair[K, V]) bool {
			a.set(pair)
			return true
		})
	}
}

func (a *hashTableAlgorithm[K, V]) growIfRequired() {
	if a.load() > a.loadFactor {
		a.resize(len(a.slice) * 2)
	}
}

func (a *hashTableAlgorithm[K, V]) shrinkIfRequired() {
	if a.size > 0 && a.load() <= a.loadFactor/4 {
		a.resize(len(a.slice) / 2)
	}
}
