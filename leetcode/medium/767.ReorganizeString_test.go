package medium

import (
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
	sort.Sort(nums)

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

func TestReorganizeString(t *testing.T) {
	Convey("test detect capital", t, func() {
		So(reorganizeString("aab"), ShouldEqual, "aba")
		So(reorganizeString("aaab"), ShouldEqual, "")
		So(reorganizeString("aaaaba"), ShouldEqual, "")
		So(reorganizeString("baaba"), ShouldEqual, "ababa")
		So(reorganizeString("vvvlo"), ShouldEqual, "vlvov")
	})
}
