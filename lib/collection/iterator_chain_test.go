package collection_test

import (
	"github.com/deemson/princeton-algorithms/lib/collection"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChainIterators(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := make([]int, 0, 3)
	collection.Each(collection.ChainIterators[int](
		collection.IteratorFromItems[int](1, 2, 3),
	), func(item int) bool {
		actual = append(actual, item)
		return true
	})
	assert.Equal(t, expected, actual)
}
