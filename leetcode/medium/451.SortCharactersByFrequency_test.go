package medium

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func frequencySort(s string) string {
	dict := make(map[rune]int)
	count := make(map[int][]string)

	// bucket count
	for _, r := range s {
		dict[r]++
	}

	// count index
	var max int
	for r, c := range dict {
		count[c] = append(count[c], strings.Repeat(string(r), c))
		if c > max {
			max = c
		}
	}

	// index to response
	var res string
	for i := max; i > 0; i-- {
		for _, str := range count[i] {
			res += str
		}
	}

	return res
}

func TestFrequencySort(t *testing.T) {
	Convey("test", t, func() {
		So(frequencySort("tree"), ShouldEqual, "eetr")
		So(frequencySort("cccaaa"), ShouldEqual, "cccaaa")
		So(frequencySort("Aabb"), ShouldEqual, "bbAa")
	})
}
