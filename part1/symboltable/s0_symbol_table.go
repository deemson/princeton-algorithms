package symboltable

import (
	"fmt"
	"github.com/deemson/princeton-algorithms/lib/collection"
)

type SymbolTable[K, V any] struct {
	algorithm algorithm[K, V]
}

func (t SymbolTable[K, V]) IsEmpty() bool {
	return collection.IsEmpty(t)
}

func (t SymbolTable[K, V]) Size() int {
	return t.algorithm.Size()
}

func (t SymbolTable[K, V]) Iterator() collection.Iterator[Pair[K, V]] {
	return t.algorithm.Iterator()
}

func (t SymbolTable[K, V]) Get(key K) (V, bool) {
	return t.algorithm.Get(key)
}

func (t SymbolTable[K, V]) MustGet(key K) V {
	item, ok := t.Get(key)
	if !ok {
		panic(fmt.Sprintf("symbol table does not contain key '%#v'", key))
	}
	return item
}

func (t SymbolTable[K, V]) MustGetMany(keys ...K) []V {
	values := make([]V, len(keys))
	for index, key := range keys {
		values[index] = t.MustGet(key)
	}
	return values
}

func (t SymbolTable[K, V]) Set(key K, value V) {
	t.algorithm.Set(key, value)
}

func (t SymbolTable[K, V]) Delete(key K) bool {
	return t.algorithm.Delete(key)
}

func (t SymbolTable[K, V]) MustDelete(key K) {
	if !t.Delete(key) {
		panic(fmt.Sprintf("symbol table does not contain key '%#v'", key))
	}
}

func (t SymbolTable[K, V]) MustDeleteMany(keys ...K) {
	for _, key := range keys {
		t.MustDelete(key)
	}
}

func (t SymbolTable[K, V]) Keys() Set[K] {
	return t.algorithm.keys()
}
