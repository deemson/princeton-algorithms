package binaryheap

import (
	"github.com/deemson/princeton-algorithms/lib/collection"
	"strings"
)

func findIndentationLevel(index int) int {
	level := -1
	powerOf2 := 1
	for powerOf2 <= index+1 {
		powerOf2 *= 2
		level++
	}
	return level
}

func visualize[T any](array collection.SizedIndexedMutable[T], toString func(item T) string, atIndex int) string {
	indentationLevel := findIndentationLevel(atIndex)
	lines := make([]string, 1, 3)
	lines[0] = strings.Repeat("  ", indentationLevel) + toString(array.Get(atIndex))
	childIndex := ChildIndex(atIndex)
	if childIndex < array.Size() {
		lines = append(lines, visualize(array, toString, childIndex))
		childIndex++
		if childIndex < array.Size() {
			lines = append(lines, visualize(array, toString, childIndex))
		}
	}
	return strings.Join(lines, "\n")
}

func Visualize[T any](array collection.SizedIndexedMutable[T], toString func(item T) string) string {
	return visualize(array, toString, 0)
}
