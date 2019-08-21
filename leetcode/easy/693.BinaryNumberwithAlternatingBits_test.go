package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

func hasAlternatingBits1(n int) bool {
	var sum, add int
	if n%4 == 2 || n%4 == 1 {
		add = n % 4
		for sum < n {
			sum += add
			add *= 4
		}
	} else {
		return false
	}

	return sum == n
}

func hasAlternatingBits(n int) bool {
	pre := -1

	for n > 0 {
		if pre == n&1 {
			return false
		}
		pre = n & 1
		n >>= 1
	}

	return true
}

func hasAlternatingBits2(n int) bool {
	s := strconv.FormatInt(int64(n), 2)
	bits := []byte(s)
	if len(bits) == 0 {
		return false
	}

	if len(bits) == 1 {
		return true
	}

	for i := 1; i < len(bits); i++ {
		if bits[i] == bits[i-1] {
			return false
		}
	}

	return true
}

func TestHasAlternatingBits(t *testing.T) {
	Convey("test", t, func() {
		So(hasAlternatingBits(1), ShouldBeTrue)
		So(hasAlternatingBits(2), ShouldBeTrue)
		So(hasAlternatingBits(5), ShouldBeTrue)
		So(hasAlternatingBits(10), ShouldBeTrue)
		So(hasAlternatingBits(7), ShouldBeFalse)
		So(hasAlternatingBits(11), ShouldBeFalse)
		So(hasAlternatingBits(21), ShouldBeTrue)
	})
}
