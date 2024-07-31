package sorting

import (
	"github.com/deemson/princeton-algorithms/lib/collection/sliceadapter"
	"github.com/deemson/princeton-algorithms/lib/compare"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	actual := []int{1, 3, 5, 2, 4, 6}
	expected := []int{1, 2, 3, 4, 5, 6}
	merge(
		sliceadapter.New(actual),
		sliceadapter.New(make([]int, len(actual))),
		compare.OrderedLess[int],
		0, 2, 5,
	)
	assert.Equal(t, expected, actual)
}
