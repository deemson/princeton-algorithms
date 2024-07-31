package unionfind

func NewWeightedQuickUnion(size int) *WeightedQuickUnion {
	parentOf := make([]int, size)
	treeSize := make([]int, size)
	for i := 0; i < size; i++ {
		parentOf[i] = i
		treeSize[i] = 1
	}
	return &WeightedQuickUnion{parentOf: parentOf, treeSize: treeSize}
}

// WeightedQuickUnion addresses main drawback of QuickUnion: trees growing too tall.
// Avoid tree growth by storing tree sizes and checking it before attaching trees to one another.
// Union and find will take at most O(log(N)) operations to complete this way.
type WeightedQuickUnion struct {
	parentOf []int
	treeSize []int
}

func (u *WeightedQuickUnion) treeRoot(i int) int {
	root := i
	for root != u.parentOf[root] {
		root = u.parentOf[root]
	}
	return root
}

func (u *WeightedQuickUnion) Union(i1 int, i2 int) {
	i1TreeRoot := u.treeRoot(i1)
	i2TreeRoot := u.treeRoot(i2)
	if u.treeSize[i1TreeRoot] < u.treeSize[i2TreeRoot] {
		u.parentOf[i1TreeRoot] = i2TreeRoot
		u.treeSize[i2TreeRoot] += u.treeSize[i1TreeRoot]
	} else {
		u.parentOf[i2TreeRoot] = i1TreeRoot
		u.treeSize[i1TreeRoot] += u.treeSize[i2TreeRoot]
	}
}

func (u *WeightedQuickUnion) Connected(i1 int, i2 int) bool {
	return u.treeRoot(i1) == u.treeRoot(i2)
}
