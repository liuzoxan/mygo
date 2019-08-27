package medium

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func findNumberOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	} else if len(nums) == 0 {
		return 0
	}

	count := 1
	max := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[i-1] {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}

	if count > max {
		max = count
	}

	return max
}

func TestFindNumberOfLIS(t *testing.T) {
	Convey("test detect capital", t, func() {
		So(findNumberOfLIS([]int{1, 3, 5, 4, 7}), ShouldEqual, 2)
		So(findNumberOfLIS([]int{2, 2, 2, 2, 2}), ShouldEqual, 5)
	})
}

func modifySlice(s []int) {
	s[0] = 0
	s = append(s, 4)
}

func modifyMap(m map[int]int) {
	m[1] = 2
	m[0] = 0
}

func TestSliceMap(t *testing.T) {
	s := []int{1, 2, 3}
	t.Log("slice:", s)
	modifySlice(s)
	t.Log("slice modified:", s)

	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	t.Log("map:", m)
	modifyMap(m)
	t.Log("map modified:", m)
}
