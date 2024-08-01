package binaryheap_test

import (
	"github.com/deemson/princeton-algorithms/part1/binaryheap"
	"github.com/gogolibs/collection"
	"github.com/gogolibs/compare"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBinaryHeapOrder(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input: "CBAGFED",
			expected: []string{
				`A`,
				`  B`,
				`    G`,
				`    F`,
				`  C`,
				`    E`,
				`    D`,
			},
		},
		{
			input: "CBAHGFED",
			expected: []string{
				`A`,
				`  B`,
				`    D`,
				`      H`,
				`    G`,
				`  C`,
				`    F`,
				`    E`,
			},
		},
	}
	for _, testCase := range testCases {
		array := collection.NewSliceAdapter(sliceStringIntoCharacters(testCase.input))
		binaryheap.Order(array, compare.Less[string])
		actual := binaryheap.Visualize(array, stringAsIs)
		assert.Equal(t, strings.Join(testCase.expected, "\n"), actual)
		assert.NoError(t, binaryheap.ValidateOrder(array, compare.Less[string]))
	}
}
