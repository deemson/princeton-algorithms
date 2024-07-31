package sorting

import (
	"github.com/deemson/princeton-algorithms/lib/collection/indexed"
	"github.com/deemson/princeton-algorithms/lib/compare"
)

type SortFunc[T any] func(indexed indexed.SizedMutable[T], less compare.Func[T])
