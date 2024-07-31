package unionfind

func NewWeightedCompressedQuickUnion(size int) *WeightedCompressedQuickUnion {
	parentOf := make([]int, size)
	treeSize := make([]int, size)
	for i := 0; i < size; i++ {
		parentOf[i] = i
		treeSize[i] = 1
	}
	return &WeightedCompressedQuickUnion{parentOf: parentOf, treeSize: treeSize}
}

// WeightedCompressedQuickUnion implements another improvement for QuickUnion:
// compression of trees during treeRoot look-ups.
type WeightedCompressedQuickUnion struct {
	parentOf []int
	treeSize []int
}

func (u *WeightedCompressedQuickUnion) treeRoot(i int) int {
	root := i
	for root != u.parentOf[root] {
		root = u.parentOf[root]
	}
	iWithNonRootParent := i
	for iWithNonRootParent != root {
		iWithNonRootParentNext := u.parentOf[iWithNonRootParent]
		u.parentOf[iWithNonRootParent] = root
		iWithNonRootParent = iWithNonRootParentNext
	}
	return root
}

func (u *WeightedCompressedQuickUnion) Union(i1 int, i2 int) {
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

func (u *WeightedCompressedQuickUnion) Connected(i1 int, i2 int) bool {
	return u.treeRoot(i1) == u.treeRoot(i2)
}
