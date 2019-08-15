package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func reverseWord(s *[]byte, left int, right int) {
	for ; left < right; {
		(*s)[left], (*s)[right] = (*s)[right], (*s)[left]
		left++
		right--
	}
}

func reverseWords(s string) string {
	if len(s) == 0 {
		return string("")
	}

	bytes := []byte(s)
	var left, right int
	for i, v := range s {
		if v == ' ' {
			right = i - 1
			reverseWord(&bytes, left, right)
			left = i + 1
		}
	}
	reverseWord(&bytes, left, len(s)-1)

	return string(bytes)
}

func TestReverseWords(t *testing.T) {
	Convey("test", t, func() {
		So(reverseWords("hello world"), ShouldEqual, string("olleh dlrow"))
		So(reverseWords("hello "), ShouldEqual, string("olleh "))
		So(reverseWords(" world"), ShouldEqual, string(" dlrow"))
		So(reverseWords(" a"), ShouldEqual, string(" a"))
		So(reverseWords(" ab"), ShouldEqual, string(" ba"))
	})
}
