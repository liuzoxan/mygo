package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func rob(nums []int) int {
	var max, prev float64

	for _, v := range nums {
		temp := max
		max = math.Max(prev+float64(v), max)
		prev = temp
	}

	return int(max)
}

func TestRob(t *testing.T) {
	Convey("test", t, func() {
		So(rob([]int{0, 1, 2, 1, 2, 1, 3}), ShouldEqual, 7)
		So(rob([]int{0, 1, 2, 9, 2, 1, 9}), ShouldEqual, 19)
		So(rob([]int{1, 0, 1, 9, 2, 1, 9}), ShouldEqual, 19)
		So(rob([]int{1, 2, 1, 1, 5, 1, 9}), ShouldEqual, 19)
	})
}
