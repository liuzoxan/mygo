package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func checkPossibility(nums []int) bool {
	breakpoint := false
	for i, v := range nums {
		if i == 0 || nums[i-1] <= v {
			continue
		}

		if breakpoint {
			return false
		}

		breakpoint = true
		if i == 1 || i == len(nums)-1 {
			continue
		} else if nums[i-2] <= v || nums[i-1] <= nums[i+1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func TestCheckPossibility(t *testing.T) {
	Convey("test", t, func() {
		So(checkPossibility([]int{1, 2, 3}), ShouldBeTrue)
		So(checkPossibility([]int{1, 2, 3, 2, 3}), ShouldBeTrue)
		So(checkPossibility([]int{1, 2, 2, 2, 3}), ShouldBeTrue)
		So(checkPossibility([]int{1, 2, 5, 2, 3}), ShouldBeTrue)
		So(checkPossibility([]int{1, 2, 5, 2, 7}), ShouldBeTrue)
		So(checkPossibility([]int{1, 2, 3, 4, 2}), ShouldBeTrue)
		So(checkPossibility([]int{1}), ShouldBeTrue)
		So(checkPossibility([]int{1, 2, 3, 2, 1}), ShouldBeFalse)
	})
}
