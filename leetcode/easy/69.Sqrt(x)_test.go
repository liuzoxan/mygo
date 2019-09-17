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

func mySqrtReal(x int, precise float64) float64 {
	if x == 0 || x == 1 {
		return float64(x)
	}

	res := 0.0
	left, right := 0.0, float64(x)
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == float64(x) {
			return mid
		} else if mid*mid > float64(x) {
			right = mid - precise
			res = right
		} else {
			left = mid + precise
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

	t.Log(mySqrtReal(1, 0.001))
	t.Log(mySqrtReal(2, 0.001))
	t.Log(mySqrtReal(4, 0.001))
	t.Log(mySqrtReal(8, 0.001))
	t.Log(mySqrtReal(9, 0.001))
}
