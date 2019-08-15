package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return string("")
	}

	minLen := -1
	for _, v := range strs {
		if minLen == -1 {
			minLen = len(v)
		} else if minLen > len(v) {
			minLen = len(v)
		}
	}

	same := true
	i := 0
	for ; i < minLen; i++ {
		for j, v := range strs {
			if j == 0 {
				continue
			}
			if v[i] != strs[0][i] {
				same = false
				break
			}
		}
		if !same {
			break
		}
	}
	if same {
		return strs[0][:minLen]
	} else {
		return strs[0][:i]
	}
}

func divideConquer(strs []string) string {
	if len(strs) == 0 {
		return string("")
	}

	if len(strs) == 1 {
		return strs[0]
	}

	if len(strs) == 2 {
		i := 0
		for ; i < len(strs[0]) && i < len(strs[1]); i++ {
			if strs[0][i] == strs[1][i] {
				continue
			} else {
				break
			}
		}
		return strs[0][:i]
	}

	left := divideConquer(strs[:len(strs)/2])
	right := divideConquer(strs[len(strs)/2:])

	return divideConquer([]string{left, right})
}

func longestCommonPrefix(strs []string) string {
	return divideConquer(strs)
}

func TestLongestCommonPrefix(t *testing.T) {
	Convey("test", t, func() {
		So(longestCommonPrefix([]string{"abc", "abcd", "abcde"}), ShouldEqual, string("abc"))
		So(longestCommonPrefix([]string{"abc", "abcd", "abcde"}), ShouldEqual, string("abc"))
		So(longestCommonPrefix([]string{"abc", "d", "abcde"}), ShouldEqual, string(""))
		So(longestCommonPrefix([]string{"abc", "", "abcde"}), ShouldEqual, string(""))
		So(longestCommonPrefix([]string{}), ShouldEqual, string(""))
		// t.Log(string("abc")[:0])
	})
}
