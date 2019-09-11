package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func repeatedStringMatch2(A string, B string) int {
	la := len(A)
	lb := len(B)
	max := lb/la + 2
	str := ""
	n := 1
	for n <= max {
		str += A
		if strings.Contains(str, B) {
			return n
		}
		n++
	}

	return -1
}

func repeatedStringMatch(A string, B string) int {
	la := len(A)
	lb := len(B)
	max := lb/la + 2
	str := ""
	n := 1
	for n <= max {
		str += A
		if strings.Contains(str, B) {
			return n
		}
		n++
	}

	return -1
}

func TestRepeatedStringMatch(t *testing.T) {
	Convey("test", t, func() {
		So(repeatedStringMatch("abcd", "cdabcdab"), ShouldEqual, 3)
	})
}
