package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func order(root *TreeNode, level int, levels *[][]int) {
	if root == nil {
		return
	}
	if len(*levels)-1 < level {
		*levels = append(*levels, []int{})
	}
	log.Println("level", level)
	(*levels)[level] = append((*levels)[level], root.Val)

	order(root.Left, level+1, levels)
	order(root.Right, level+1, levels)
}

func levelOrder(root *TreeNode) [][]int {
	var levels [][]int
	order(root, 0, &levels)
	return levels
}

func TestNextGreaterElement(t *testing.T) {
	Convey("test detect capital", t, func() {
		t1 := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		}

		t.Log(levelOrder(t1))
	})
}
