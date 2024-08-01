package sorting

import (
	"github.com/gogolibs/collection"
	"github.com/gogolibs/compare"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	actual := []int{1, 3, 5, 2, 4, 6}
	expected := []int{1, 2, 3, 4, 5, 6}
	merge(
		collection.NewSliceAdapter(actual),
		collection.NewSliceAdapter(make([]int, len(actual))),
		compare.Less[int],
		0, 2, 5,
	)
	assert.Equal(t, expected, actual)
}
