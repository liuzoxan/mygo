package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func largeGroupPositions(S string) [][]int {
	var res [][]int
	if len(S) < 3 {
		return res
	}

	preStart := 0
	preCount := 1
	for i := 1; i < len(S); i++ {
		if S[i] != S[i-1] {
			if preCount >= 3 {
				res = append(res, []int{preStart, preStart + preCount - 1})
			}
			preCount = 1
			preStart = i
		} else {
			preCount++
		}
	}

	if preCount >= 3 {
		res = append(res, []int{preStart, preStart + preCount - 1})
	}

	return res
}

func TestLargeGroupPositions(t *testing.T) {
	Convey("test", t, func() {
		t.Log(largeGroupPositions("abbxxxxzzy"))
		t.Log(largeGroupPositions("abc"))
		t.Log(largeGroupPositions("abcdddeeeeaabbbcd"))
		t.Log(largeGroupPositions("aaa"))
	})
}
