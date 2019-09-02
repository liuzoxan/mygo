package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"sort"
	"testing"
)

type Cost struct {
	A int
	B int
}

type Costs []Cost

func (c Costs) Len() int {
	return len(c)
}

func (c Costs) Less(i, j int) bool {
	return math.Abs(float64(c[i].A)-float64(c[i].B)) > math.Abs(float64(c[j].A)-float64(c[j].B))
}

func (c Costs) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func twoCitySchedCost1(costs [][]int) int {
	// sort by diff abs(A - B)
	var nCosts Costs
	for _, c := range costs {
		nCosts = append(nCosts, Cost{
			A: c[0],
			B: c[1],
		})
	}
	sort.Sort(nCosts)

	// select correct city from large diff
	countA := len(costs) / 2
	countB := countA
	minCost := 0
	for _, c := range nCosts {
		if countA == 0 {
			minCost += c.B
			continue
		}

		if countB == 0 {
			minCost += c.A
			continue
		}

		if c.A < c.B {
			countA--
			minCost += c.A
		} else {
			countB--
			minCost += c.B
		}
	}

	return minCost
}

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][1]-costs[i][0] > costs[j][1]-costs[j][0]
	})

	l := len(costs) / 2
	minCost := 0
	for i := 0; i < l; i++ {
		minCost = minCost + costs[i][0] + costs[i+l][1]
	}

	return minCost
}

func TestTwoCitySchedCost(t *testing.T) {
	Convey("test", t, func() {
		So(twoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}), ShouldEqual, 110)
		So(twoCitySchedCost([][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}), ShouldEqual, 1859)
	})
}
