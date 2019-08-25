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

func levelOrder1(root *TreeNode) [][]int {
	var levels [][]int
	order(root, 0, &levels)
	return levels
}

func levelOrder(root *TreeNode) [][]int {
	var levels [][]int

	var queue []*TreeNode
	if root == nil {
		return levels
	}
	queue = append(queue, root)

	for len(queue) > 0 {
		qLen := len(queue)
		var li []int
		for i := 0; i < qLen; i++ {
			node := queue[0]
			li = append(li, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levels = append(levels, li)
	}

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
