package deque_test

import (
	"fmt"
	"testing"
)

func getBenchmarkSizes() []int {
	return []int{10, 100, 1000}
}

func BenchmarkDeque_Get(b *testing.B) {
	dequeMakers := testTableBuilder[int]{
		capacity: 2,
	}.build()

	sizes := getBenchmarkSizes()

	for _, makeDeque := range dequeMakers {
		for _, size := range sizes {
			name, d := makeDeque()
			for number := 0; number < size; number++ {
				d.AddLast(number)
			}
			b.Run(fmt.Sprintf("%s/%d", name, size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					d.Get(size / 2)
				}
			})
		}
	}
}
