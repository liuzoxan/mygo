package chaolesson

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1,2,3,4,5}
	// t.Log(a == c)
	t.Log(a == b)
}

func TestZeroBit(t *testing.T) {
	t.Log(2&^1, 2&^0, 2&^2)

	const (
		Readable = 1 << iota
		Writable
		Executable
	)
	a := 7
	t.Log(Readable, Writable, Executable)
	t.Log(a&^Readable, a&^Writable, a&^Executable)

}
