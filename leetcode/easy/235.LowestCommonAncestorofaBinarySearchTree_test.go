package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func find(root, p, q *TreeNode) *TreeNode {
	if root.Val < p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}

	return find(root, p, q)
}

func TestLowestCommonAncestor(t *testing.T) {
	Convey("test", t, func() {
		t1 := &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
			},
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
			},
		}

		t.Log(lowestCommonAncestor(t1, &TreeNode{
			Val: 2,
		}, &TreeNode{
			Val: 8,
		}).Val)

		t.Log(lowestCommonAncestor(t1, &TreeNode{
			Val: 4,
		}, &TreeNode{
			Val: 5,
		}).Val)

		t.Log(lowestCommonAncestor(t1, &TreeNode{
			Val: 2,
		}, &TreeNode{
			Val: 4,
		}).Val)

		t.Log(lowestCommonAncestor(t1, &TreeNode{
			Val: 3,
		}, &TreeNode{
			Val: 7,
		}).Val)
	})
}
