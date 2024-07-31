package deque

import "fmt"

func ensureIndexNotNegative(index int) {
	if index < 0 {
		panic(fmt.Sprintf("Deque index must not be negative (%d)", index))
	}
}

func ensureIndexLessThanSize(index int, size int) {
	ensureIndexNotNegative(index)
	if index >= size {
		panic(fmt.Sprintf("Deque index (%d) must be less than size (%d)", index, size))
	}
}

func ensureIndexLessOrEqualSize(index int, size int) {
	ensureIndexNotNegative(index)
	if index > size {
		panic(fmt.Sprintf("Deque index (%d) must be less or equal size (%d)", index, size))
	}
}

func ensureNotEmpty(size int) {
	if size == 0 {
		panic("Deque must not be empty")
	}
}
