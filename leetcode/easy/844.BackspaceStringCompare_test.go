package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func backspaceCompare1(S string, T string) bool {
	var s, t []byte

	for _, v := range []byte(S) {
		if v != '#' {
			s = append(s, v)
		} else {
			if len(s) > 0 {
				s = s[:len(s)-1]
			}
		}
	}

	for _, v := range []byte(T) {
		if v != '#' {
			t = append(t, v)
		} else {
			if len(t) > 0 {
				t = t[:len(t)-1]
			}
		}
	}

	if len(s) != len(t) {
		return false
	}

	for i := range s {
		if s[i] != t[i] {
			return false
		}
	}

	return true
}

func backspaceBuild1(S string) []byte {
	var r []byte

	s := []byte(S)
	skip := 0
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == '#' {
			skip++
		} else {
			if skip > 0 {
				skip--
			} else {
				r = append(r, s[i])
			}
		}
	}

	return r
}

func backspaceBuild2(s string) []byte {
	var r []byte
	skip := 0
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == '#' {
			skip++
		} else {
			if skip > 0 {
				skip--
			} else {
				r = append(r, s[i])
			}
		}
	}

	return r
}

func backspaceBuild(s string) string {
	for i := 0; i < len(s); {
		if s[i] == '#' {
			if i == 0 {
				s = s[i+1:]
			} else if i == 1 {
				s = s[i+1:]
				i = 0
			} else {
				s = s[:i-1] + s[i+1:]
				i = i - 1
			}
		} else {
			i++
		}
	}

	return s
}

func backspaceCompare(S string, T string) bool {
	return backspaceBuild(S) == backspaceBuild(T)
}

func TestBackspaceCompare(t *testing.T) {
	Convey("test", t, func() {
		So(backspaceCompare(string("ab#c"), string("ac")), ShouldBeTrue)
		So(backspaceCompare(string("a##c"), string("c")), ShouldBeTrue)
		So(backspaceCompare(string("#c"), string("c")), ShouldBeTrue)
		So(backspaceCompare(string("abc#c"), string("abc###")), ShouldBeFalse)
	})
}
