package symboltable

import (
	"github.com/deemson/princeton-algorithms/part1/deque"
	"github.com/gogolibs/compare"
	"github.com/gogolibs/iterator"
)

func UnorderedLinkedList[K, V any](equal compare.Func[K]) SymbolTable[K, V] {
	return SymbolTable[K, V]{
		algorithm: unorderedLinkedListAlgorithm[K, V]{
			deque: deque.LinkedList[Pair[K, V]](),
			equal: equal,
		},
	}
}

type unorderedLinkedListAlgorithm[K, V any] struct {
	deque deque.Deque[Pair[K, V]]
	equal compare.Func[K]
}

func (a unorderedLinkedListAlgorithm[K, V]) Iterator() iterator.Iterator[Pair[K, V]] {
	return a.deque.Iterator()
}

func (a unorderedLinkedListAlgorithm[K, V]) Size() int {
	return a.deque.Size()
}

func (a unorderedLinkedListAlgorithm[K, V]) Get(key K) (V, bool) {
	var value V
	isFound := !a.deque.Each(func(pair Pair[K, V]) bool {
		if a.equal(pair.Key, key) {
			value = pair.Value
			return false
		}
		return true
	})
	return value, isFound
}

func (a unorderedLinkedListAlgorithm[K, V]) Set(key K, value V) {
	indexToReplace := -1
	index := 0
	isFound := !a.deque.Each(func(pair Pair[K, V]) bool {
		if a.equal(pair.Key, key) {
			indexToReplace = index
			return false
		}
		index++
		return true
	})
	pair := Pair[K, V]{
		Key:   key,
		Value: value,
	}
	if isFound {
		a.deque.Set(indexToReplace, pair)
	} else {
		a.deque.AddLast(pair)
	}
}

func (a unorderedLinkedListAlgorithm[K, V]) Delete(key K) bool {
	indexToDelete := -1
	index := 0
	isFound := !a.deque.Each(func(pair Pair[K, V]) bool {
		if a.equal(pair.Key, key) {
			indexToDelete = index
			return false
		}
		index++
		return true
	})
	if isFound {
		a.deque.RemoveAtIndex(indexToDelete)
	}
	return isFound
}
