package deque

import "github.com/gogolibs/iterator"

func LinkedList[T any]() Deque[T] {
	return Deque[T]{
		algorithm: &linedListAlgorithm[T]{
			size: 0,
			head: nil,
			tail: nil,
		},
	}
}

type linedListAlgorithm[T any] struct {
	size int
	head *linkedListNode[T]
	tail *linkedListNode[T]
}

func (a *linedListAlgorithm[T]) Size() int {
	return a.size
}

func (a *linedListAlgorithm[T]) Get(index int) T {
	return a.nodeAtIndex(index).item
}

func (a *linedListAlgorithm[T]) Iterator() iterator.Iterator[T] {
	return &linkedListIterator[T]{
		node: a.head,
	}
}

func (a *linedListAlgorithm[T]) Set(index int, item T) {
	a.nodeAtIndex(index).item = item
}

func (a *linedListAlgorithm[T]) AddAtIndex(index int, item T) {
	if index == a.size {
		prevTail := a.tail
		a.tail = &linkedListNode[T]{
			item: item,
			prev: prevTail,
			next: nil,
		}
		a.size++
		if a.size > 1 {
			prevTail.next = a.tail
		} else {
			a.head = a.tail
		}
	} else {
		currentNode := a.nodeAtIndex(index)
		addedNode := &linkedListNode[T]{
			item: item,
			prev: currentNode.prev,
			next: currentNode,
		}
		a.size++
		if currentNode.prev != nil {
			currentNode.prev.next = addedNode
		} else {
			a.head = addedNode
		}
		currentNode.prev = addedNode
	}
}

func (a *linedListAlgorithm[T]) RemoveAtIndex(index int) T {
	var item T
	switch index {
	case 0:
		item = a.head.item
		a.head = a.head.next
		if a.size == 1 {
			a.tail = nil
		} else {
			a.head.prev = nil
		}
	case a.size - 1:
		item = a.tail.item
		a.tail = a.tail.prev
		a.tail.next = nil
	default:
		removedNode := a.nodeAtIndex(index)
		item = removedNode.item
		removedNode.prev.next = removedNode.next
		removedNode.next.prev = removedNode.prev
	}
	a.size--
	return item
}

func (a *linedListAlgorithm[T]) nodeAtIndex(index int) *linkedListNode[T] {
	var node *linkedListNode[T]
	// loop through the nodes either from head or tail -- whichever requires fewer iterations
	if 2*index < a.size-1 {
		// looping from head is faster
		node = a.head
		for loopIndex := 0; loopIndex < index; loopIndex++ {
			node = node.next
		}
	} else {
		// looping from tail is faster
		node = a.tail
		for loopIndex := a.size - 1; loopIndex > index; loopIndex-- {
			node = node.prev
		}
	}
	return node
}

type linkedListNode[T any] struct {
	item T
	prev *linkedListNode[T]
	next *linkedListNode[T]
}

type linkedListIterator[T any] struct {
	node *linkedListNode[T]
}

func (i *linkedListIterator[T]) HasNext() bool {
	return i.node != nil
}

func (i *linkedListIterator[T]) Next() T {
	item := i.node.item
	i.node = i.node.next
	return item
}
