package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func detectCapitalUse(word string) bool {
	bytes := []byte(word)
	firstLower := false
	secondLower := false
	for i, b := range bytes {
		if i == 0 {
			if b >= 'a' && b <= 'z' {
				firstLower = true
			}
			continue
		}

		if i == 1 {
			if b >= 'a' && b <= 'z' {
				secondLower = true
				continue
			} else {
				secondLower = false
				if firstLower {
					return false
				}
			}
		}

		if firstLower && secondLower {
			if b >= 'A' && b <= 'Z' {
				return false
			}
		} else if !firstLower && !secondLower {
			if b >= 'a' && b <= 'z' {
				return false
			}
		} else if !firstLower && secondLower {
			if b >= 'A' && b <= 'Z' {
				return false
			}
		}
	}

	return true
}

func TestDetectCapital(t *testing.T) {
	Convey("test detect capital", t, func() {
		So(detectCapitalUse("Abc"), ShouldBeTrue)
		So(detectCapitalUse("ABC"), ShouldBeTrue)
		So(detectCapitalUse("abc"), ShouldBeTrue)
		So(detectCapitalUse("aBC"), ShouldBeFalse)
		So(detectCapitalUse("AbC"), ShouldBeFalse)
	})
}
