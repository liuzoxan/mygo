package medium

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"sort"
	"testing"
)

func helper(root *TreeNode, strs *[][]byte, cur []int) {
	if root == nil {
		return
	}
	cur = append(cur, root.Val)

	if root.Left == nil && root.Right == nil {
		var rev []byte
		for i := len(cur) - 1; i >= 0; i-- {
			rev = append(rev, byte(cur[i]+'a'))
		}
		*strs = append(*strs, rev)
		return
	}

	helper(root.Left, strs, cur)
	helper(root.Right, strs, cur)
}

func smallestFromLeaf1(root *TreeNode) string {
	var strs [][]byte

	helper(root, &strs, []int{})

	sort.Slice(strs, func(i, j int) bool {
		mLen := int(math.Min(float64(len(strs[i])), float64(len(strs[j]))))

		for m := 0; m < mLen; m++ {
			if strs[i][m] == strs[j][m] {
				continue
			} else {
				return strs[i][m] < strs[j][m]
			}
		}

		return len(strs[i]) < len(strs[j])
	})

	if len(strs) == 0 {
		return ""
	}

	return string(strs[0])
}

func divideConquer(root *TreeNode, parent *TreeNode) string {
	if root == nil {
		return ""
	}

	leftMin := divideConquer(root.Left, root)
	rightMin := divideConquer(root.Right, root)

	if leftMin == "" {
		return rightMin + string(root.Val+'a')
	}

	if rightMin == "" {
		return leftMin + string(root.Val+'a')
	}

	leftMin += string(root.Val + 'a')
	rightMin += string(root.Val + 'a')

	if leftMin+string(parent.Val+'a') < rightMin+string(parent.Val+'a') {
		return leftMin
	} else {
		return rightMin
	}
}

func smallestFromLeaf(root *TreeNode) string {
	parent := &TreeNode{
		Val:   0,
		Left:  root,
		Right: nil,
	}

	return divideConquer(root, parent)
}

func TestSmallestFromLeaf(t *testing.T) {
	Convey("test detect capital", t, func() {
		t1 := &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		}

		So(smallestFromLeaf(t1), ShouldEqual, "dba")

		t2 := &TreeNode{
			Val: 25,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		}

		So(smallestFromLeaf(t2), ShouldEqual, "ababz")

	})
}
