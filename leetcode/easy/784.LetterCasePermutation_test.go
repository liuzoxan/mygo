package easy

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func add1(l *[]string, s []byte, i int) {
	if i == len(s)-1 {
		*l = append(*l, string(s))

		if s[i] < 'A' {
			return
		}
		if s[i] < 'a' {
			s[i] = s[i] + 32
		} else {
			s[i] = s[i] - 32
		}

		*l = append(*l, string(s))
		return
	}
	add1(l, s, i+1)
	if s[i] < 'A' {
		return
	}
	if s[i] < 'a' {
		s[i] = s[i] + 32
	} else {
		s[i] = s[i] - 32
	}
	add1(l, s, i+1)
}

func add(l *[]string, s *[]byte, i int) {
	if i == len(*s) {
		*l = append(*l, string(*s))
	} else {
		if (*s)[i] >= 'A' && (*s)[i] <= 'Z' {
			add(l, s, i+1)
			(*s)[i] += 32
			add(l, s, i+1)
		} else if (*s)[i] >= 'a' && (*s)[i] <= 'z' {
			add(l, s, i+1)
			(*s)[i] -= 32
			add(l, s, i+1)
		} else {
			add(l, s, i+1)
		}
	}
}

func letterCasePermutation(S string) []string {
	s := []byte(S)
	var r []string
	add(&r, &s, 0)

	return r
}

func letterCasePermutation1(S string) []string {
	var r []string
	s := []byte(S)

	for i := 0; i < len(S); i++ {
		if len(r) == 0 {
			r = append(r, string(bytes.ToUpper(s[i:i+1])))
			if s[i] < 'A' {
				continue
			}
			r = append(r, string(bytes.ToLower(s[i:i+1])))
			continue
		}
		l := len(r)
		var rt []string
		for j := 0; j < l; j++ {
			rt = append(rt, r[j]+string(bytes.ToUpper(s[i:i+1])))
			if s[i] < 'A' {
				continue
			}
			rt = append(rt, r[j]+string(bytes.ToLower(s[i:i+1])))
		}
		r = rt
	}

	return r
}

func TestLetterCasePermutation(t *testing.T) {
	Convey("test", t, func() {
		t.Log(letterCasePermutation("a"))
		t.Log(letterCasePermutation("ab"))
		t.Log(letterCasePermutation("a1bc"))
		t.Log(letterCasePermutation("a1b2c"))
	})
}
