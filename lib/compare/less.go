package compare

import "cmp"

func OrderedLess[T cmp.Ordered](item1, item2 T) bool {
	return item1 < item2
}
