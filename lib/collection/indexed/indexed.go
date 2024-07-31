package indexed

import "github.com/deemson/princeton-algorithms/lib/collection/sized"

type Indexed[T any] interface {
	Get(index int) T
}

type SizedIndexed[T any] interface {
	sized.Sized
	Indexed[T]
}
