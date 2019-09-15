package medium

import (
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

func maxProfitAssignment1(difficulty []int, profit []int, worker []int) int {
	profitM := make(map[int]int)
	for i, d := range difficulty {
		if profit[i] > profitM[d] {
			profitM[d] = profit[i]
		}
	}

	sort.Slice(difficulty, func(i, j int) bool {
		return difficulty[i] < difficulty[j]
	})

	maxProfitM := make(map[int]int)
	maxProfit := 0
	for _, d := range difficulty {
		if maxProfit < profitM[d] {
			maxProfit = profitM[d]
		}
		maxProfitM[d] = maxProfit
	}


	sum := 0
	l := len(difficulty)
	for _, a := range worker {
		index := sort.SearchInts(difficulty, a)
		if index == l {
			sum += maxProfitM[difficulty[l-1]]
			continue
		}

		if difficulty[index] == a {
			sum += maxProfitM[a]
		} else {
			if index != 0 {
				sum += maxProfitM[difficulty[index-1]]
			}
		}
	}

	return sum
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	profitM := make(map[int]int, len(difficulty))
	for i, d := range difficulty {
		if profit[i] > profitM[d] {
			profitM[d] = profit[i]
		}
	}

	sort.Slice(difficulty, func(i, j int) bool {
		return difficulty[i] < difficulty[j]
	})

	maxProfitM := make(map[int]int, len(difficulty))
	maxProfit := 0
	for _, d := range difficulty {
		if maxProfit < profitM[d] {
			maxProfit = profitM[d]
		}
		maxProfitM[d] = maxProfit
	}


	sum := 0
	l := len(difficulty)
	for _, a := range worker {
		index := sort.SearchInts(difficulty, a)
		if index == l {
			sum += maxProfitM[difficulty[l-1]]
			continue
		}

		if difficulty[index] == a {
			sum += maxProfitM[a]
		} else {
			if index != 0 {
				sum += maxProfitM[difficulty[index-1]]
			}
		}
	}

	return sum
}

func TestMaxProfitAssignment(t *testing.T) {
	Convey("test detect capital", t, func() {
		t.Log(sort.SearchInts([]int{1, 3, 4, 6, 7, 9}, 2))
		t.Log(sort.SearchInts([]int{1, 3, 4, 6, 7, 9}, 0))
		t.Log(maxProfitAssignment([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}))
		t.Log(maxProfitAssignment(
			[]int{66, 1,  28, 73, 53, 35, 45, 60, 100, 44, 59, 94, 27, 88, 7,  18, 83, 18, 72,  63},
			[]int{66, 20, 84, 81, 56, 40, 37, 82, 53,  45, 43, 96, 67, 27, 12, 54, 98, 19, 47,  77},
			[]int{61, 33, 68, 38, 63, 45, 1,  10, 53,  23, 66, 70, 14, 51, 94, 18, 28, 78, 100, 16}))
	})
}
