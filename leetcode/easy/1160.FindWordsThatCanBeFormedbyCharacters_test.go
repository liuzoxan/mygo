package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func countCharacters(words []string, chars string) int {
	tableS := make(map[rune]int, 26)
	for _, r := range chars {
		tableS[r] ++
	}

	// dynamic table for loop
	table := make(map[rune]int, 26)

	var res string
	for _, str := range words {
		// fresh table
		for r, c := range tableS {
			table[r] = c
		}

		match := true
		for _, r := range str {
			table[r]--
			if table[r] < 0 {
				match = false
				break
			}
		}
		if match {
			res += str
		}
	}

	return len(res)
}

func TestCountCharacters(t *testing.T) {
	Convey("test", t, func() {
		So(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"), ShouldEqual, 6)
		So(countCharacters([]string{"hello","world","leetcode"}, "welldonehoneyr"), ShouldEqual, 10)
	})
}
