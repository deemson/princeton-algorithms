package indexed

import "github.com/deemson/princeton-algorithms/lib/collection/sized"

type Mutable[T any] interface {
	Indexed[T]
	Set(index int, item T)
}

func Swap[T any](indexed Mutable[T], index1, index2 int) {
	tmp := indexed.Get(index1)
	indexed.Set(index1, indexed.Get(index2))
	indexed.Set(index2, tmp)
}

type SizedMutable[T any] interface {
	sized.Sized
	Mutable[T]
}
