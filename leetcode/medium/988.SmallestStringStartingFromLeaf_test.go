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

func smallestFromLeaf(root *TreeNode) string {
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

	})
}
