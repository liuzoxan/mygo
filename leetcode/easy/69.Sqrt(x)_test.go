package easy

import "testing"

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	res := 0
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
			res = right
		} else {
			left = mid + 1
		}
	}

	return res
}

func TestSqrt(t *testing.T) {
	t.Log(mySqrt(1))
	t.Log(mySqrt(2))
	t.Log(mySqrt(4))
	t.Log(mySqrt(8))
	t.Log(mySqrt(9))
}
