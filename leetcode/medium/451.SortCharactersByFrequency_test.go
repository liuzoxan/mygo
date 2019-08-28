package medium

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func frequencySort(s string) string {
	bytes := []byte(s)
	dict := make(map[byte]int)
	count := make(map[int][]byte)

	// bucket count
	for _, b := range bytes {
		if _, OK := dict[b]; !OK {
			dict[b] = 0
		}
		dict[b]++
	}

	// count index
	var max int
	for b, c := range dict {
		if _, OK := count[c]; !OK {
			count[c] = []byte{}
		}
		count[c] = append(count[c], b)
		if c > max {
			max = c
		}
	}

	// index to response
	var res string
	for i := max; i > 0; i-- {
		for _, b := range count[i] {
			for j := 0; j < i; j++ {
				res += string(b)
			}
		}
	}

	return res
}

func TestFrequencySort(t *testing.T) {
	Convey("test", t, func() {
		So(frequencySort("tree"), ShouldEqual, "eert")
		So(frequencySort("cccaaa"), ShouldEqual, "cccaaa")
		So(frequencySort("Aabb"), ShouldEqual, "bbAa")
	})
}
