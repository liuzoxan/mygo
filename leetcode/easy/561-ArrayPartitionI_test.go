package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

func arrayPairSum_Slow(nums []int) int {
	sort.Ints(nums)

	sum := 0

	l := len(nums)
	for i := 0; i < l/2; i++ {
		sum += nums[2*i]
	}
	return sum
}

func arrayPairSum(nums []int) int {
	var buckets [20001]int
	for _, v := range nums {
		buckets[v+10000]++
	}

	sum := 0
	l := len(buckets)
	odd := true
	for i := 0; i < l; i++ {
		for buckets[i] > 0 {
			if odd {
				sum += i - 10000
			}
			odd = !odd
			buckets[i]--
		}
	}
	return sum
}

func TestArrayPairSum(t *testing.T) {
	Convey("test", t, func() {
		So(arrayPairSum([]int{1, 3, 4, 2}), ShouldEqual, 4)
		So(arrayPairSum([]int{-1, 3, 4, 2}), ShouldEqual, 2)
		So(arrayPairSum([]int{-10, 89, 4, -1}), ShouldEqual, -6)
		So(arrayPairSum([]int{1, 2, 3, 2}), ShouldEqual, 3)
	})
}
