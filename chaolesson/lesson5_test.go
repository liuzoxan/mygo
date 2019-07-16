package chaolesson

import "testing"

func TestFib(t *testing.T) {
	t.Log("log test")

	a := 1
	b := 1
	for i := 0; i < 10; i++ {
		t.Log(a)
		temp := a
		a = b
		b = temp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a

	t.Log(a, b)
}

func TestConst(t *testing.T) {
	const (
		Monday = iota +1
		Tuesday
		Wednesday
	)
	t.Log(Monday, Tuesday, Wednesday)

	const (
		Readable = 1 << iota
		Writable
		Executable
	)
	t.Log(Readable, Writable, Executable)
}