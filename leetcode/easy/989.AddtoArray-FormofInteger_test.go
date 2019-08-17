package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func addToArrayForm1(A []int, K int) []int {
	var res, resR []int

	// convert K to slice, but reversed...
	var k []int
	for K > 0 {
		k = append(k, K%10)
		K /= 10
	}

	post := 0
	sum := 0
	i := 0
	for ; i < len(A); i++ {
		if i > len(k)-1 {
			sum = A[len(A)-1-i] + post
		} else {
			sum = A[len(A)-1-i] + k[i] + post
		}
		res = append(res, sum%10)
		post = sum / 10
	}

	for ; i < len(k); i++ {
		sum = k[i] + post
		res = append(res, sum%10)
		post = sum / 10
	}

	if post != 0 {
		res = append(res, post)
	}

	for i := len(res) - 1; i >= 0; i-- {
		resR = append(resR, res[i])
	}

	return resR
}

func addToArrayForm(A []int, K int) []int {
	var res, resR []int

	aLen := len(A)
	post := 0
	sum := 0
	for K > 0 {
		k := K % 10
		K /= 10

		if aLen > 0 {
			sum = k + A[aLen-1] + post
			aLen--
		} else {
			sum = k + post
		}

		res = append(res, sum%10)
		post = sum / 10
	}

	for aLen > 0 {
		sum = A[aLen-1] + post
		aLen--
		res = append(res, sum%10)
		post = sum / 10
	}

	if post != 0 {
		res = append(res, post)
	}

	for i := len(res) - 1; i >= 0; i-- {
		resR = append(resR, res[i])
	}

	return resR
}

func TestAddToArrayForm(t *testing.T) {
	Convey("test", t, func() {
		t.Log(addToArrayForm([]int{9, 8, 5}, 15))
		t.Log(addToArrayForm([]int{8, 5}, 15))
		t.Log(addToArrayForm([]int{5}, 15))
		t.Log(addToArrayForm([]int{5, 2}, 150))
		t.Log(addToArrayForm([]int{}, 150))
		t.Log(addToArrayForm([]int{1, 5, 0}, 0))
		t.Log(addToArrayForm([]int{0}, 0))
	})
}
