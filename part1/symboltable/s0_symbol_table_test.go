package symboltable

import (
	"github.com/deemson/princeton-algorithms/lib/compare"
	"github.com/stretchr/testify/assert"
	"testing"
)

type allAlgorithmsParams[K, V any] struct {
}

func (p allAlgorithmsParams[K, V]) do(f func(t *testing.T, table SymbolTable[K, V])) {

}

func TestSymbolTable(t *testing.T) {
	allAlgorithmsParams[string, int]{}.do(func(t *testing.T, symbolTable SymbolTable[string, int]) {
		symbolTable.Set("one", 1)
		symbolTable.Set("two", 2)
		symbolTable.Set("three", 3)
		assert.Equal(t, []string{"one", "three", "two"}, symbolTable.Keys().ToSortedSlice(compare.OrderedLess[string]))
		assert.Equal(t, []int{1, 3, 2}, symbolTable.MustGetMany(symbolTable.Keys().ToSortedSlice(compare.OrderedLess[string])...))
		symbolTable.Set("two", 4)
		assert.Equal(t, []int{1, 3, 4}, symbolTable.MustGetMany(symbolTable.Keys().ToSortedSlice(compare.OrderedLess[string])...))
		symbolTable.Delete("two")
		assert.Equal(t, []int{1, 3}, symbolTable.MustGetMany(symbolTable.Keys().ToSortedSlice(compare.OrderedLess[string])...))
		symbolTable.MustDeleteMany("one", "three")
		assert.True(t, symbolTable.IsEmpty())
	})
}
