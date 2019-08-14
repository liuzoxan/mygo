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

func backspaceCompare(S string, T string) bool {
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

func TestBackspaceCompare(t *testing.T) {
	Convey("test", t, func() {
		So(backspaceCompare(string("ab#c"), string("ac")), ShouldBeTrue)
		So(backspaceCompare(string("a##c"), string("c")), ShouldBeTrue)
		So(backspaceCompare(string("#c"), string("c")), ShouldBeTrue)
		So(backspaceCompare(string("abc#c"), string("abc###")), ShouldBeFalse)
	})
}
