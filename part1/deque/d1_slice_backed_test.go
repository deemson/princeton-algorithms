package deque_test

import (
	"github.com/deemson/princeton-algorithms/part1/deque"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceBacked_GrowByAddingLast(t *testing.T) {
	d := deque.SliceBacked[int](2)
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)
	assert.Equal(t, 3, d.Size())
	assert.Equal(t, []int{1, 2, 3}, d.ToSlice())
}

func TestSliceBacked_GrowByAddingFirst(t *testing.T) {
	d := deque.SliceBacked[int](2)
	d.AddFirst(1)
	d.AddFirst(2)
	d.AddFirst(3)
	assert.Equal(t, 3, d.Size())
	assert.Equal(t, []int{3, 2, 1}, d.ToSlice())
}

func TestSliceBacked_ShrinkByAddingLastAndRemovingFirst(t *testing.T) {
	d := deque.SliceBacked[int](4)
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)
	assert.Equal(t, 3, d.Size())
	assert.Equal(t, 1, d.RemoveFirst())
	assert.Equal(t, 2, d.RemoveFirst())
	assert.Equal(t, 1, d.Size())
	assert.Equal(t, 3, d.RemoveFirst())
	assert.True(t, d.IsEmpty())
}

func TestSliceBacked_ShrinkByAddingFirstAndRemovingLast(t *testing.T) {
	d := deque.SliceBacked[int](4)
	d.AddFirst(1)
	d.AddFirst(2)
	d.AddFirst(3)
	assert.Equal(t, 3, d.Size())
	assert.Equal(t, 1, d.RemoveLast())
	assert.Equal(t, []int{3, 2}, d.ToSlice())
	assert.Equal(t, 2, d.RemoveLast())
	assert.Equal(t, 1, d.Size())
	assert.Equal(t, 3, d.RemoveLast())
	assert.True(t, d.IsEmpty())
}

func TestSliceBacked_AddAtIndex_MovingItemsToBothEnds(t *testing.T) {
	d := deque.SliceBacked[int](2)
	d.AddLast(1)
	d.AddLast(3)
	d.AddLast(5)
	d.AddAtIndex(1, 2)
	assert.Equal(t, []int{1, 2, 3, 5}, d.ToSlice())
	d.AddAtIndex(3, 4)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, d.ToSlice())
}

func TestSliceBacked_RemoveAtIndex_MovingItemsToBothEnds(t *testing.T) {
	d := deque.SliceBacked[int](2)
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)
	d.AddLast(4)
	d.AddLast(5)
	assert.Equal(t, 4, d.RemoveAtIndex(3))
	assert.Equal(t, []int{1, 2, 3, 5}, d.ToSlice())
	assert.Equal(t, 2, d.RemoveAtIndex(1))
	assert.Equal(t, []int{1, 3, 5}, d.ToSlice())
}

func TestSliceBacked_RemoveAtIndex_FirstItemIndexWrapAround(t *testing.T) {
	d := deque.SliceBacked[int](4)
	d.AddLast(2)
	d.AddFirst(1)
	assert.Equal(t, []int{1, 2}, d.ToSlice())
	assert.Equal(t, 1, d.RemoveFirst())
	assert.Equal(t, []int{2}, d.ToSlice())
}
