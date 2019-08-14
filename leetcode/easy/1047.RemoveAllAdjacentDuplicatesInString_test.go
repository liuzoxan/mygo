package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func removeDuplicates1(S string) string {
	for i := 0; i < len(S); {
		if i == 0 {
			i++
			continue
		}
		if S[i] == S[i-1] {
			S = S[:i-1] + S[i+1:]
			i = i - 1
		} else {
			i++
		}
	}

	return S
}

func removeDuplicates(S string) string {
	var r []byte
	for i := 0; i < len(S); i++ {
		if len(r) == 0 || S[i] != r[len(r)-1] {
			r = append(r, S[i])
		} else {
			r = r[:len(r)-1]
		}
	}

	return string(r)
}

func TestRemoveDuplicates(t *testing.T) {
	Convey("test", t, func() {
		So(removeDuplicates("abbaca"), ShouldEqual, "ca")
		So(removeDuplicates("abbaabbbca"), ShouldEqual, "abca")
	})
}
