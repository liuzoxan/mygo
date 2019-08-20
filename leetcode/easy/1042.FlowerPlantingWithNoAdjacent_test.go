package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func gardenNoAdj(N int, paths [][]int) []int {
	nextPaths := make([][]int, N)
	for _, p := range paths {
		i, j := p[0]-1, p[1]-1
		nextPaths[i] = append(nextPaths[i], j)
		nextPaths[j] = append(nextPaths[j], i)
	}

	plant := make([]int, N)
	for i := 0; i < N; i++ {
		isUsed := [5]bool{}
		for _, p := range nextPaths[i] {
			isUsed[plant[p]] = true
		}

		for j := 1; j <= 4; j++ {
			if !isUsed[j] {
				plant[i] = j
				break
			}
		}
	}

	return plant
}

func TestGardenNoAdj(t *testing.T) {
	Convey("test", t, func() {
		t.Log(gardenNoAdj(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
		t.Log(gardenNoAdj(5, [][]int{{1, 2}, {2, 3}, {3, 1}, {3, 4}, {5, 4}}))
		t.Log(gardenNoAdj(4, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}}))
	})
}
