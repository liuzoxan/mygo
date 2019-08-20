package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func gardenNoAdj(N int, paths [][]int) []int {
	plant := make(map[int]int)
	nextPaths := make(map[int][][]int)
	for _, p := range paths {
		if _, OK := nextPaths[p[0]]; !OK {
			nextPaths[p[0]] = [][]int{}
		}
		nextPaths[p[0]] = append(nextPaths[p[0]], p)

		if _, OK := nextPaths[p[1]]; !OK {
			nextPaths[p[1]] = [][]int{}
		}
		nextPaths[p[1]] = append(nextPaths[p[1]], p)
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= 4; j++ {
			same := false
			for _, p := range nextPaths[i] {
				if p[0] == i {
					if plant[p[1]] == j {
						same = true
						break
					}
				} else if p[1] == i {
					if plant[p[0]] == j {
						same = true
						break
					}
				}
			}
			if !same {
				plant[i] = j
				break
			}
		}
	}

	var res []int
	for i := 1; i <= N; i++ {
		res = append(res, plant[i])
	}
	return res
}

func TestGardenNoAdj(t *testing.T) {
	Convey("test", t, func() {
		t.Log(gardenNoAdj(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
		t.Log(gardenNoAdj(5, [][]int{{1, 2}, {2, 3}, {3, 1}, {3, 4}, {5, 4}}))
		t.Log(gardenNoAdj(4, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}}))
	})
}
