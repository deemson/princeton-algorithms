package symboltable

import (
	"github.com/deemson/princeton-algorithms/lib/collection"
)

type algorithm[K, V any] interface {
	collection.Sized
	collection.Sequence[Pair[K, V]]
	Get(key K) (V, bool)
	Set(key K, value V)
	Delete(key K) bool
	keys() Set[K]
}
