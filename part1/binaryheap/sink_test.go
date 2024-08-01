package binaryheap_test

import (
	"github.com/deemson/princeton-algorithms/part1/binaryheap"
	"github.com/gogolibs/collection"
	"github.com/gogolibs/compare"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSink_Simple_TopToBottom(t *testing.T) {
	slice := sliceStringIntoCharacters("CBA")
	binaryheap.SinkFromTopToBottom(collection.NewSliceAdapter(slice), compare.Less[string])
	assert.Equal(t, sliceStringIntoCharacters("ABC"), slice)
}

func TestSink_To(t *testing.T) {
	slice := sliceStringIntoCharacters("GFEDCBA")
	binaryheap.SinkFromTop(collection.NewSliceAdapter(slice), compare.Less[string], 4)
	assert.Equal(t, sliceStringIntoCharacters("EFGDCBA"), slice)
}
