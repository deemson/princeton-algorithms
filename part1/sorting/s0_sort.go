package sorting

import (
	"github.com/gogolibs/collection"
	"github.com/gogolibs/compare"
)

type SortFunc[T any] func(array collection.SizedIndexedMutable[T], less compare.Func[T])
