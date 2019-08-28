package medium

import (
	"container/heap"
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"strings"
	"testing"
)

// rune count item
type rCount struct {
	val   rune
	count int
}

// rune count list
type rList []rCount

func (s rList) Len() int {
	return len(s)
}

func (s rList) Less(i, j int) bool {
	return s[i].count > s[j].count
}

func (s rList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *rList) Push(x interface{}) {
	*s = append(*s, x.(rCount))
}

func (s *rList) Pop() interface{} {
	top := (*s)[0]
	*s = (*s)[1:]

	return top
}

func reorganizeString1(S string) string {
	if len(S) <= 1 {
		return S
	}

	// count rune
	count := make(map[rune]int)
	for _, r := range S {
		count[r] ++

		// larger than ...
		if count[r] > (len(S)+1)/2 {
			return ""
		}
	}

	// map to slice
	var nums rList
	for k, v := range count {
		nums = append(nums, rCount{
			val:   k,
			count: v,
		})
	}

	// sort by rune count
	sort.Sort(&nums)

	// get whole string
	var temp string
	for _, n := range nums {
		temp += strings.Repeat(string(n.val), n.count)
	}

	// rotate the string
	res := make([]rune, len(S))
	mid := (len(S) + 1) / 2
	for i, r := range temp {
		if i < mid {
			res[2*i] = r
		} else {
			res[2*(i-mid)+1] = r
		}
	}

	return string(res)
}

func reorganizeString(S string) string {
	if len(S) <= 1 {
		return S
	}

	// count rune
	count := make(map[rune]int)
	for _, r := range S {
		count[r] ++

		// larger than ...
		if count[r] > (len(S)+1)/2 {
			return ""
		}
	}

	// map to slice
	var nums rList
	for k, v := range count {
		nums = append(nums, rCount{
			val:   k,
			count: v,
		})
	}

	// sort by rune count
	heap.Init(&nums)

	// get whole string
	var res []rune
	for len(nums) > 0 {
		n1 := nums.Pop().(rCount)
		res = append(res, n1.val)
		n1.count--

		if len(nums) == 0 {
			break
		} else {
			n2 := nums.Pop().(rCount)
			res = append(res, n2.val)
			n2.count--

			if n2.count > 0 {
				heap.Push(&nums, n2)
			}
		}

		if n1.count > 0 {
			heap.Push(&nums, n1)
		}
	}

	return string(res)
}

func TestReorganizeString(t *testing.T) {
	Convey("test detect capital", t, func() {
		So(reorganizeString("aab"), ShouldEqual, "aba")
		So(reorganizeString("aaab"), ShouldEqual, "")
		So(reorganizeString("aaaaba"), ShouldEqual, "")
		So(reorganizeString("baaba"), ShouldEqual, "ababa")
		// So(reorganizeString("vvvlo"), ShouldEqual, "vlvov")
	})
}
