package chaolesson

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

func TestSliceInt(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 3, 4, 5}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"1", "2", "3", "1", "2", "3", "1", "2", "3", "1", "2", "3"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
}

func TestSliceCompare(t *testing.T) {
	a := []int{1, 3}
	b := []int{1, 2, 3}
	// t.Log(a == b)
	t.Log(a, b)

}

func TestJson(t *testing.T) {
	var ids []string
	jstr := []byte(`{"1":{"a":1, "b":2}, "2":{"a":3, "b":4}}`)
	var data map[string]interface{} // map[1] = interface{} map[2]=interface{}
	json.Unmarshal(jstr, &data)
	for k, v := range data {
		ids = append(ids, k)
		t.Log(v.(map[string]interface{})["a"].(float64))
	}

	jstr2 := []byte(`{"first":{"a":1, "b":2}, "second":{"a":3, "b":4}}`)

	type Detail struct {
		A float64 `json:"a"`
		B float64 `json:"b"`
	}
	type Number struct {
		First  Detail `json:"first"`
		Second Detail `json:"second"`
	}

	var n Number
	json.Unmarshal(jstr2, &n)
	t.Log(n.Second.A)
}

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 3
	t.Log(m1[3])
	if v, ok := m1[3]; ok {
		t.Log("key exists", v)
	} else {
		t.Log("key no exist")
	}

	m2 := map[int]int{1: 1, 2: 2, 3: 3}
	for k, v := range m2 {
		t.Log(k, v)
	}
}

func TestMapFunc(t *testing.T) {
	m := map[int]func(int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](1), m[2](2), m[3](3))
}

func TestMapForSet(t *testing.T) {
	set := map[int]bool{}
	set[1] = true
	if set[1] {
		t.Log("exists")
	} else {
		t.Log("not exist")
	}

	set[3] = true
	t.Log(set[3])
	delete(set, 3)
	if set[3] {
		t.Log("exist")
	} else {
		t.Log("not exist")
	}
}

func TestString1(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	s = "\xE4\xBB\xA5"
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(len(s))
	c := []rune(s)
	t.Log(len(c))
	t.Logf("unicode %x", c[0])
	t.Logf("UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"

	sr := []rune(s)
	t.Log(len(s), len(sr))
	for _, c := range sr {
		t.Logf("%[1]x %[1]d %s", c, string(c))
	}

	for _, sc := range s {
		t.Logf("%s", reflect.TypeOf(sc))
		t.Log(string(sc))
	}

	s1 := "中国good"
	for _, c:= range s1 {
		t.Logf("%c", c)
	}
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("string :", s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log("int :", i)
	}
}
