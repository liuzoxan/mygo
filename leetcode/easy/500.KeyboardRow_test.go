package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func findWords(words []string) []string {
	dict := map[rune]int{
		'q': 1,
		'w': 1,
		'e': 1,
		'r': 1,
		't': 1,
		'y': 1,
		'u': 1,
		'i': 1,
		'o': 1,
		'p': 1,
		'a': 2,
		's': 2,
		'd': 2,
		'f': 2,
		'g': 2,
		'h': 2,
		'j': 2,
		'k': 2,
		'l': 2,
		'z': 3,
		'x': 3,
		'c': 3,
		'v': 3,
		'b': 3,
		'n': 3,
		'm': 3,
	}

	var res []string
	for _, str := range words {
		tempStr := strings.ToLower(str)

		prev := 0
		same := true
		for _, b := range tempStr {
			if prev != 0 && dict[b] != prev {
				same = false
				break
			}
			prev = dict[b]
		}
		if same {
			res = append(res, str)
		}
	}

	return res
}

func TestFindWords(t *testing.T) {
	Convey("test", t, func() {
		So(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}), ShouldResemble, []string{"Alaska", "Dad"})
	})
}
