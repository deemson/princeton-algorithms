package sequence

import (
	"github.com/deemson/princeton-algorithms/lib/collection/iterator"
	"github.com/deemson/princeton-algorithms/lib/collection/sized"
)

type Sequence[T any] interface {
	Iterator() iterator.Iterator[T]
}

type SizedSequence[T any] interface {
	sized.Sized
	Sequence[T]
}
