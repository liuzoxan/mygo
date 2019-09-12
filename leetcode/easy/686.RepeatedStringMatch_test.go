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
	max := len(B)/len(A) + 2
	n := len(B) / len(A)
	for n <= max {
		str := strings.Repeat(A, n)
		if strings.Contains(str, B) {
			return n
		}
		n++
	}

	return -1
}

func repeatedStringMatch4(A string, B string) int {
	byteA := []byte(A)
	count := 1
	for len(byteA) < len(B) {
		byteA = append(byteA, byteA[:len(A)]...)
		count++
	}
	if strings.Index(string(byteA), B) != -1 {
		return count
	}
	byteA = append(byteA, byteA[:len(A)]...)
	count++
	if strings.Index(string(byteA), B) != -1 {
		return count
	}
	return -1
}

func TestRepeatedStringMatch(t *testing.T) {
	Convey("test", t, func() {
		So(repeatedStringMatch("abcd", "cdabcdab"), ShouldEqual, 3)
	})
}
