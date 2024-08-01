package sorting_test

import (
	"github.com/deemson/princeton-algorithms/part1/sorting"
	"github.com/gogolibs/collection"
	"github.com/gogolibs/compare"
	"github.com/stretchr/testify/assert"
	"testing"
)

func forEachAlgorithm[T any](t *testing.T, less compare.Func[T], f func(t *testing.T, sort func(array collection.SizedIndexedMutable[T]))) {
	namedSorts := []struct {
		name string
		sort sorting.SortFunc[T]
	}{
		{"SelectionSort", sorting.SelectionSort[T]},
		{"InsertionSort", sorting.InsertionSort[T]},
		{"ShellSort", sorting.ShellSort[T]},
		{"MergeSort", sorting.MergeSort[T]},
		{"BottomUpMergeSort", sorting.BottomUpMergeSort[T]},
		{"QuickSort", sorting.QuickSort[T]},
		{"HeapSort", sorting.HeapSort[T]},
	}
	for _, namedSort := range namedSorts {
		t.Run(namedSort.name, func(t *testing.T) {
			f(t, func(array collection.SizedIndexedMutable[T]) {
				namedSort.sort(array, less)
			})
		})
	}
}

func TestSort_BunchOfInts(t *testing.T) {
	forEachAlgorithm(t, compare.Less[int], func(t *testing.T, sort func(array collection.SizedIndexedMutable[int])) {
		input := collection.NewSliceAdapter([]int{
			42,
			17,
			100500,
			123,
			3,
			13,
			256,
			4242,
			127,
		})
		expected := []int{
			3,
			13,
			17,
			42,
			123,
			127,
			256,
			4242,
			100500,
		}
		sort(input)
		actual := collection.ToSlice[int](input)
		assert.Equal(t, expected, actual)
	})
}

func TestSort_BunchOfStrings(t *testing.T) {
	forEachAlgorithm(t, compare.Less[string], func(t *testing.T, sort func(array collection.SizedIndexedMutable[string])) {
		input := collection.NewSliceAdapter([]string{"super", "algorithm", "main"})
		expected := []string{"algorithm", "main", "super"}
		sort(input)
		actual := collection.ToSlice[string](input)
		assert.Equal(t, expected, actual)
	})
}
