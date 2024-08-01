package binaryheap_test

import (
	"github.com/deemson/princeton-algorithms/part1/binaryheap"
	"github.com/gogolibs/collection"
	"github.com/gogolibs/compare"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwim_Simple_BottomToTop(t *testing.T) {
	slice := []string{"C", "B", "A"}
	binaryheap.SwimFromBottomToTop(collection.NewSliceAdapter(slice), compare.Less[string])
	assert.Equal(t, []string{"A", "B", "C"}, slice)
}
