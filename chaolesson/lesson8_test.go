package chaolesson

import "testing"

func TestLoop(t *testing.T) {
	a := 10
	for a > 0 {
		t.Log(a)
		a --
	}

	// for {
	// 	t.Log(1)
	// }

}

func TestSwitch(t *testing.T) {
	num := 5
	switch {
	case 0 < num && num < 3:
		t.Log("small")
	case 4 < num && num < 6:
		t.Log("medium")
	case num < 10:
		t.Log("large")

	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5;i++{
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0 - 3")
		}
	}
}