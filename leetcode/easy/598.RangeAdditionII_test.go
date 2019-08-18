package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func maxCount(m int, n int, ops [][]int) int {
	minR, minC := m, n
	for _, o := range ops {
		minR = int(math.Min(float64(o[0]), float64(minR)))
		minC = int(math.Min(float64(o[1]), float64(minC)))
	}

	return minR * minC
}

func TestMaxCount(t *testing.T) {
	Convey("test", t, func() {
		So(maxCount(3, 3, [][]int{{2, 2}, {3, 3}}), ShouldEqual, 4)
		So(maxCount(1, 3, [][]int{{2, 2}, {3, 3}}), ShouldEqual, 2)
		So(maxCount(4, 3, [][]int{{5, 2}, {1, 3}}), ShouldEqual, 2)
	})
}
