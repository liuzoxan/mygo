package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func getMin(root *TreeNode, min *int) {
	if root.Left == nil || root.Right == nil {
		return
	}

	if root.Left.Val != root.Val {
		if *min == -1 {
			*min = root.Left.Val
		} else if root.Left.Val < *min {
			*min = root.Left.Val
		}
	} else {
		getMin(root.Left, min)
	}

	if root.Right.Val != root.Val {
		if *min == -1 {
			*min = root.Right.Val
		} else if root.Right.Val < *min {
			*min = root.Right.Val
		}
	} else {
		getMin(root.Right, min)
	}
}

func findSecondMinimumValue(root *TreeNode) int {
	if root.Left == nil {
		return -1
	}

	min := -1
	getMin(root, &min)

	return min
}

func TestFindSecondMinimumValue(t *testing.T) {
	Convey("test", t, func() {
		t1 := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		}

		So(findSecondMinimumValue(t1), ShouldEqual, 5)

		t2 := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		}

		So(findSecondMinimumValue(t2), ShouldEqual, -1)

	})
}
