package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func profit(day int, holdPrice int, hold bool, prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	if day == len(prices)-1 {
		if hold {
			return prices[day] - holdPrice
		} else {
			return 0
		}
	}

	var p1, p2 int
	if hold {
		p1 = profit(day+1, holdPrice, true, prices)
		p2 = profit(day+1, 0, false, prices) + prices[day] - holdPrice
	} else {
		p1 = profit(day+1, prices[day], true, prices)
		p2 = profit(day+1, 0, false, prices)
	}

	if p1 >= p2 {
		return p1
	} else {
		return p2
	}
}

// exceed time limit
func maxProfit1(prices []int) int {
	return profit(0, 0, false, prices)
}

func TestMaxProfit(t *testing.T) {
	Convey("test", t, func() {
		So(maxProfit([]int{7, 1, 5, 3, 6, 4}), ShouldEqual, 7)
		So(maxProfit([]int{1, 2, 3, 4, 5}), ShouldEqual, 4)
		So(maxProfit([]int{7, 6, 4, 3, 1}), ShouldEqual, 0)
	})
}
