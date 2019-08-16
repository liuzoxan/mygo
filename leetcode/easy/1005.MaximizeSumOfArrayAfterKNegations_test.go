package easy

import (
	"container/heap"
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

func largestSumAfterKNegations1(A []int, K int) int {
	sort.Ints(A)
	sum := 0
	for i, v := range A {
		if v < 0 {
			if K > 0 {
				K--
				sum += -v
			} else {
				sum += v
			}
		} else if K%2 == 0 {
			sum += v
		} else {
			if i > 0 && -A[i-1] < v {
				sum -= -A[i-1]
				sum += A[i-1]
				sum += v
			} else {
				sum += -v
			}
			K = 0
		}
	}
	return sum
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func largestSumAfterKNegations(A []int, K int) int {
	h := intHeap(A)
	heap.Init(&h)

	for K > 0 {
		h[0] = -h[0]
		K--
		heap.Fix(&h, 0)
	}

	sum := 0
	for _, v := range h {
		sum += v
	}

	return sum
}

func TestLargestSumAfterKNegations(t *testing.T) {
	Convey("test", t, func() {
		So(largestSumAfterKNegations([]int{1, 2, 3}, 1), ShouldEqual, 4)
		So(largestSumAfterKNegations([]int{4, 2, 3}, 1), ShouldEqual, 5)
		So(largestSumAfterKNegations([]int{-1, 2, 3}, 1), ShouldEqual, 6)
		So(largestSumAfterKNegations([]int{3, -1, 0, 2}, 2), ShouldEqual, 6)
		So(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2), ShouldEqual, 13)
		So(largestSumAfterKNegations([]int{-8, -5, -5, -3, -2, 3}, 6), ShouldEqual, 22)
	})
}
