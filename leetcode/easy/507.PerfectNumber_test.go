package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func checkPerfectNumber1(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	m := make(map[int]int)
	start := int(math.Sqrt(float64(num)))
	for start > 1 {
		if m[start] == 1 {
			start--
			continue
		}

		if num%start == 0 {
			m[start] = 1
			m[num/start] = 1
			sum += start + num/start
		}
		start--
	}

	return sum == num
}

func checkPerfectNumber(num int) bool {
	var result int

	if num <= 1 {
		return false
	}
	i := 2
	for i < num/i && i != num {
		if num%i == 0 {
			result += i
			result += num / i
		}
		i++
	}

	return (result + 1) == num
}

func TestCheckPerfectNumber(t *testing.T) {
	Convey("test detect capital", t, func() {
		So(checkPerfectNumber(28), ShouldBeTrue)
	})
}
