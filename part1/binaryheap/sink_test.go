package binaryheap_test

import (
	"github.com/deemson/princeton-algorithms/lib/collection"
	"github.com/deemson/princeton-algorithms/lib/compare"
	"github.com/deemson/princeton-algorithms/part1/binaryheap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSink_Simple_TopToBottom(t *testing.T) {
	slice := sliceStringIntoCharacters("CBA")
	binaryheap.SinkFromTopToBottom(collection.Slice(slice), compare.OrderedLess[string])
	assert.Equal(t, sliceStringIntoCharacters("ABC"), slice)
}

func TestSink_To(t *testing.T) {
	slice := sliceStringIntoCharacters("GFEDCBA")
	binaryheap.SinkFromTop(collection.Slice(slice), compare.OrderedLess[string], 4)
	assert.Equal(t, sliceStringIntoCharacters("EFGDCBA"), slice)
}
