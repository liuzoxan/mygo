package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	var res []int
	for _, n1 := range nums1 {
		found := false
		start := false
		for _, n2 := range nums2 {
			if n2 == n1 {
				start = true
			}
			if start && n2 > n1 {
				res = append(res, n2)
				found = true
				break
			}
		}

		if !found {
			res = append(res, -1)
		}
	}
	return res

}

func dp(nums1 []int, nums2 []int, mark map[int]int) {
	for i, n := range nums2 {
		for j := len(nums1) - 1; j >= 0; j-- {
			if n > nums1[j] {
				mark[nums1[j]] = n
				nums1 = nums1[:j]
			}
		}

		if _, OK := mark[n]; OK {
			if i == len(nums2)-1 {
				return
			}

			nums1 = append(nums1, n)
			dp(nums1, nums2[i+1:], mark)
			break
		}
	}
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mark := make(map[int]int, 1000)
	for _, n := range nums1 {
		mark[n] = -1
	}

	dp([]int{}, nums2, mark)

	var res []int
	for _, n := range nums1 {
		res = append(res, mark[n])
	}

	return res

}

func TestNextGreaterElement(t *testing.T) {
	Convey("test detect capital", t, func() {
		t.Log(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
		t.Log(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
		t.Log(nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}))
		t.Log(nextGreaterElement([]int{1, 2}, []int{2, 1, 7}))
	})
}
