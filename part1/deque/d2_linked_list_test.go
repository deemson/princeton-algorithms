package deque_test

import (
	"github.com/deemson/princeton-algorithms/part1/deque"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_AddingLast(t *testing.T) {
	d := deque.LinkedList[string]()
	d.AddLast("one")
	assert.Equal(t, []string{"one"}, d.ToSlice())
	d.AddLast("two")
	assert.Equal(t, []string{"one", "two"}, d.ToSlice())
}

func TestLinkedList_GettingAndSettingFromHead(t *testing.T) {
	d := deque.LinkedList[int]()
	d.AddAllLast(1, 2, 3, 4, 5)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, d.ToSlice())
	assert.Equal(t, 2, d.Get(1))
	d.Set(1, 42)
	assert.Equal(t, []int{1, 42, 3, 4, 5}, d.ToSlice())
}

func TestLinkedList_GettingAndSettingFromTail(t *testing.T) {
	d := deque.LinkedList[int]()
	d.AddAllLast(1, 2, 3, 4, 5)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, d.ToSlice())
	assert.Equal(t, 4, d.Get(3))
	d.Set(3, 42)
	assert.Equal(t, []int{1, 2, 3, 42, 5}, d.ToSlice())
}

func TestLinkedList_AddingHead(t *testing.T) {
	d := deque.LinkedList[int]()
	d.AddLast(2)
	d.AddFirst(1)
	assert.Equal(t, []int{1, 2}, d.ToSlice())
}

func TestLinkedList_AddingMiddle(t *testing.T) {
	d := deque.LinkedList[int]()
	d.AddLast(3)
	d.AddFirst(1)
	d.AddAtIndex(1, 2)
	assert.Equal(t, []int{1, 2, 3}, d.ToSlice())
}

func TestLinkedList_RemovingHead(t *testing.T) {
	d := deque.LinkedList[int]()
	d.AddAllLast(1, 2)
	assert.Equal(t, 1, d.RemoveFirst())
	assert.Equal(t, []int{2}, d.ToSlice())
	assert.Equal(t, 2, d.RemoveFirst())
	assert.True(t, d.IsEmpty())
}

func TestLinkedList_RemovingTail(t *testing.T) {
	d := deque.LinkedList[int]()
	d.AddAllLast(1, 2)
	assert.Equal(t, 2, d.RemoveLast())
	assert.Equal(t, []int{1}, d.ToSlice())
	assert.Equal(t, 1, d.RemoveLast())
	assert.True(t, d.IsEmpty())
}

func TestLinkedList_RemovingMiddle(t *testing.T) {
	d := deque.LinkedList[int]()
	d.AddAllLast(1, 2, 3)
	assert.Equal(t, 2, d.RemoveAtIndex(1))
	assert.Equal(t, []int{1, 3}, d.ToSlice())
}
