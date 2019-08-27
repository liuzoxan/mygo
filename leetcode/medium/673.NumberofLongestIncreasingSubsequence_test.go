package medium

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func findNumberOfLIS(nums []int) int {
	size := len(nums)

	length := make([]int, size)
	count := make([]int, size)
	for i := 0; i < size; i++ {
		count[i] = 1
	}

	for j, n := range nums {
		for i := 0; i < j; i++ {
			if nums[i] < n {
				if length[i] >= length[j] {
					length[j] = 1 + length[i]
					count[j] = count[i]
				} else if length[i]+1 == length[j] {
					count[j] = count[j] + count[i]
				}
			}
		}
	}

	var maxLen int
	for _, l := range length {
		if l > maxLen {
			maxLen = l
		}
	}

	var sum int
	for i, c := range count {
		if length[i] == maxLen {
			sum += c
		}
	}

	return sum
}

func TestFindNumberOfLIS(t *testing.T) {
	Convey("test detect capital", t, func() {
		So(findNumberOfLIS([]int{1, 3, 5, 4, 7}), ShouldEqual, 2)
		So(findNumberOfLIS([]int{2, 2, 2, 2, 2}), ShouldEqual, 5)
		So(findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2}), ShouldEqual, 3)
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
