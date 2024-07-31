package unionfind

func NewQuickUnion(size int) *QuickUnion {
	parentOf := make([]int, size)
	for i := 0; i < size; i++ {
		parentOf[i] = i
	}
	return &QuickUnion{parentOf: parentOf}
}

// QuickUnion fixes the drawbacks of QuickFind.
// Union is implemented by setting parent-child nodes creating a tree for complex unions.
// Drawbacks: trees can get tall making finds expensive (could be up to N in the worst case).
type QuickUnion struct {
	parentOf []int
}

// treeRoot is the method for finding union tree root.
// The roots are the nodes, that have never been set as another node's child
// therefore having the parent it had during the initialization, i.e. i == parentOf[i]
// meaning it is a parent to itself.
func (u *QuickUnion) treeRoot(i int) int {
	root := i
	for root != u.parentOf[root] {
		root = u.parentOf[root]
	}
	return root
}

// Union links two nodes by setting root of one as a parent to another
// creating a bigger tree in the process. This implementation
// sets the root of the n1 as a parent to the root of n2.
func (u *QuickUnion) Union(i1 int, i2 int) {
	i1TreeRoot := u.treeRoot(i1)
	i2TreeRoot := u.treeRoot(i2)
	u.parentOf[i2TreeRoot] = i1TreeRoot
}

func (u *QuickUnion) Connected(i1 int, i2 int) bool {
	return u.treeRoot(i1) == u.treeRoot(i2)
}
