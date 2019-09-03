package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func numEquivDominoPairs(dominoes [][]int) int {
	var dom [100]int

	for _, d := range dominoes {
		if d[0] <= d[1] {
			dom[d[0]*10+d[1]] ++
		} else {
			dom[d[1]*10+d[0]] ++
		}
	}

	sum := 0
	for _, c := range dom {
		sum += (c - 1) * c / 2
	}

	return sum
}

func TestNumEquivDominoPairs(t *testing.T) {
	Convey("test", t, func() {
		So(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}), ShouldEqual, 1)
	})
}
