package priorityqueue_test

import (
	"github.com/deemson/princeton-algorithms/part1/priorityqueue"
	"github.com/gogolibs/compare"
	"github.com/stretchr/testify/assert"
	"testing"
)

func forEachAlgorithm[T any](t *testing.T, less compare.Func[T], f func(t *testing.T, priorityQueue priorityqueue.PriorityQueue[T])) {
	namedPriorityQueues := []struct {
		name          string
		priorityQueue priorityqueue.PriorityQueue[T]
	}{
		{"UnorderedArray", priorityqueue.UnorderedArray(less, 2)},
		{"OrderedArray", priorityqueue.OrderedArray(less, 2)},
		{"BinaryHeap", priorityqueue.BinaryHeap(less, 2)},
	}
	for _, namedPriorityQueue := range namedPriorityQueues {
		t.Run(namedPriorityQueue.name, func(t *testing.T) {
			f(t, namedPriorityQueue.priorityQueue)
		})
	}
}

func TestPriorityQueue(t *testing.T) {
	forEachAlgorithm[int](t, compare.Less[int], func(t *testing.T, priorityQueue priorityqueue.PriorityQueue[int]) {
		priorityQueue.PushMany(5, 1, 13, 100500, 42)
		expected := []int{1, 5, 13, 42, 100500}
		actual := priorityQueue.PopMany(len(expected))
		assert.Equal(t, expected, actual)
	})
}

func TestPriorityQueue_Reversed(t *testing.T) {
	forEachAlgorithm[int](t, compare.Greater[int], func(t *testing.T, priorityQueue priorityqueue.PriorityQueue[int]) {
		priorityQueue.PushMany(42, 100500, 13, 1, 5)
		expected := []int{100500, 42, 13, 5, 1}
		actual := priorityQueue.PopMany(len(expected))
		assert.Equal(t, expected, actual)
	})
}
