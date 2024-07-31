package binaryheap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindIndentationLevel(t *testing.T) {
	indentations := []int{
		0,
		1,
		1,
		2,
		2,
		2,
		2,
		3,
		3,
		3,
		3,
		3,
		3,
		3,
		3,
	}
	for input, expected := range indentations {
		actual := findIndentationLevel(input)
		name := fmt.Sprintf("%d: %d == %d", input, expected, actual)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, expected, actual)
		})
	}
}
