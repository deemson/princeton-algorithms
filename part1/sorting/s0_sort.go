package sorting

import (
	"github.com/deemson/princeton-algorithms/lib/collection"
	"github.com/deemson/princeton-algorithms/lib/compare"
)

type SortFunc[T any] func(array collection.SizedIndexedMutable[T], less compare.Func[T])
