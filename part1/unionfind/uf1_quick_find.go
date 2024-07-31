package unionfind

func NewQuickFind(size int) *QuickFind {
	values := make([]int, size)
	for i := 0; i < size; i++ {
		values[i] = i
	}
	return &QuickFind{values: values}
}

type QuickFind struct {
	values []int
}

// Union is as slow as O(N)
func (u *QuickFind) Union(i1 int, i2 int) {
	i1id := u.values[i1]
	i2id := u.values[i2]
	for index, value := range u.values {
		if value == i1id {
			u.values[index] = i2id
		}
	}
}

func (u *QuickFind) Connected(i1 int, i2 int) bool {
	return u.values[i1] == u.values[i2]
}
