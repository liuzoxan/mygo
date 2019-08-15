package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for ; left <= right; {
		pivot := (left + right) / 2
		if nums[pivot] == target {
			return pivot
		}
		if nums[pivot] > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return -1
}

func TestBinarySearch(t *testing.T) {
	Convey("test", t, func() {
		So(search([]int{0, 1, 2, 3}, 2), ShouldEqual, 2)
		So(search([]int{0, 1, 2, 3}, 1), ShouldEqual, 1)
		So(search([]int{0, 1, 2, 3}, 0), ShouldEqual, 0)
		So(search([]int{0, 1, 2, 3}, 4), ShouldEqual, -1)
	})
}
