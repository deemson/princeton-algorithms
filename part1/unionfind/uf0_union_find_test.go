package unionfind_test

import (
	"github.com/deemson/princeton-algorithms/part1/unionfind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const size = 10

func forEachAlgorithm(t *testing.T, f func(t *testing.T, unionFind unionfind.UnionFind)) {
	namedUnionFinds := []struct {
		name      string
		unionFind unionfind.UnionFind
	}{
		{"QuickFind", unionfind.NewQuickFind(size)},
		{"QuickUnion", unionfind.NewQuickUnion(size)},
		{"WeightedQuickUnion", unionfind.NewWeightedQuickUnion(size)},
		{"WeightedCompressedQuickUnion", unionfind.NewWeightedCompressedQuickUnion(size)},
	}
	for _, namedUnionFind := range namedUnionFinds {
		t.Run(namedUnionFind.name, func(t *testing.T) {
			f(t, namedUnionFind.unionFind)
		})
	}
}

func TestUnionFind(t *testing.T) {
	forEachAlgorithm(t, func(t *testing.T, unionFind unionfind.UnionFind) {
		for i := 1; i < size; i++ {
			require.False(t, unionFind.Connected(i-1, i))
		}
		unionFind.Union(3, 5)
		assert.True(t, unionFind.Connected(3, 5))
		unionFind.Union(6, 3)
		assert.True(t, unionFind.Connected(3, 6))
		assert.True(t, unionFind.Connected(5, 6))
		unionFind.Union(7, 8)
		assert.True(t, unionFind.Connected(7, 8))
		assert.False(t, unionFind.Connected(3, 8))
	})
}
