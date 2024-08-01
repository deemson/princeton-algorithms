package compare

func ComparableEqual[T comparable](item1, item2 T) bool {
	return item1 == item2
}
