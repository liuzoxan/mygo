package chaolesson

import "testing"

func TestArrayInit(t *testing.T)  {
	var arr1 [3]int
	arr2 := [4]int{1,2,3,4}
	arr3 := [...]int{1,2,3,4,5}
	arr1[1] = 5
	t.Log(arr1, arr2, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1,2,3,4}
	for _, v := range arr3 {
		t.Log(v)
	}
	arr5 := [2][2]int{{1,2}, {3,4}}
	t.Log(arr5)
}

func TestArraySlice(t *testing.T) {
	arr3 := [...]int{1,2,3,4,5}
	t.Log(arr3[1:3], arr3[1:1])
}

