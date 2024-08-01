package symboltable_test

import (
	"github.com/deemson/princeton-algorithms/lib/compare"
	"github.com/deemson/princeton-algorithms/part1/symboltable"
	"github.com/stretchr/testify/assert"
	"testing"
)

type allAlgorithmsParams[K, V any] struct {
	equal compare.Func[K]
}

func (p allAlgorithmsParams[K, V]) do(t *testing.T, f func(t *testing.T, symbolTable symboltable.SymbolTable[K, V])) {
	namedSymbolTables := []struct {
		name        string
		symbolTable symboltable.SymbolTable[K, V]
	}{
		{"UnorderedLinkedList", symboltable.UnorderedLinkedList[K, V](p.equal)},
	}
	for _, namedSymbolTable := range namedSymbolTables {
		t.Run(namedSymbolTable.name, func(t *testing.T) {
			f(t, namedSymbolTable.symbolTable)
		})
	}
}

func TestSymbolTable(t *testing.T) {
	allAlgorithmsParams[string, int]{
		equal: compare.ComparableEqual[string],
	}.do(t, func(t *testing.T, symbolTable symboltable.SymbolTable[string, int]) {
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
