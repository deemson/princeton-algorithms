package symboltable

import (
	"github.com/gogolibs/collection"
)

type algorithm[K, V any] interface {
	collection.Sized
	collection.Sequence[Pair[K, V]]
	Get(key K) (V, bool)
	Set(key K, value V)
	Delete(key K) bool
}
