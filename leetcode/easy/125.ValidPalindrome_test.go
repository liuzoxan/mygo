package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func validChar(r byte) bool {
	return r >= 'a' && r <= 'z' || r >= '0' && r <= '9'
}

func isPalindrome(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}

	s = strings.ToLower(s)

	var i, j int
	for i, j = 0, l-1; i < j; {
		if !validChar(s[i]) {
			i++
			continue
		}

		if !validChar(s[j]) {
			j--
			continue
		}

		if s[i] == s[j] {
			i++
			j--
			continue
		} else {
			return false
		}
	}

	return true
}

func TestIsPalindrome(t *testing.T) {
	Convey("test", t, func() {
		So(isPalindrome("A man, a plan, a canal: Panama"), ShouldBeTrue)
		So(isPalindrome("race a car"), ShouldBeFalse)
		So(isPalindrome("0P"), ShouldBeFalse)
		So(isPalindrome("`l;`` 1o1 ??;l`"), ShouldBeTrue)
	})
}
