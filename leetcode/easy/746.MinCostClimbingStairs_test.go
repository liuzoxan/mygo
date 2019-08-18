package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func minCostClimbingStairs(cost []int) int {
	var f1, f2 int

	for i := len(cost) - 1; i >= 0; i-- {
		f0 := cost[i] + int(math.Min(float64(f1), float64(f2)))
		f2 = f1
		f1 = f0
	}

	return int(math.Min(float64(f1), float64(f2)))
}

func TestMinCostClimbingStairs(t *testing.T) {
	Convey("test", t, func() {
		So(minCostClimbingStairs([]int{1, 2, 3}), ShouldEqual, 2)
		So(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}), ShouldEqual, 6)
	})
}
