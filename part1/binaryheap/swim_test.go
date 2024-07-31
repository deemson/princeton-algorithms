package binaryheap_test

import (
	"github.com/deemson/princeton-algorithms/lib/collection/sliceadapter"
	"github.com/deemson/princeton-algorithms/lib/compare"
	"github.com/deemson/princeton-algorithms/part1/binaryheap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwim_Simple_BottomToTop(t *testing.T) {
	slice := []string{"C", "B", "A"}
	binaryheap.SwimFromBottomToTop(sliceadapter.New(slice), compare.OrderedLess[string])
	assert.Equal(t, []string{"A", "B", "C"}, slice)
}