package easy

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func rootPath(root *TreeNode, path string, paths *[]string) {
	if path == "" {
		path = fmt.Sprintf("%d", root.Val)
	} else {
		path = fmt.Sprintf("%s->%d", path, root.Val)
	}

	if root.Right == nil && root.Left == nil {
		*paths = append(*paths, path)
	} else {
		if root.Right != nil {
			rootPath(root.Right, path, paths)
		}

		if root.Left != nil {
			rootPath(root.Left, path, paths)
		}
	}
}

func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	if root == nil {
		return paths
	}
	rootPath(root, "", &paths)
	return paths
}

func TestBinaryTreePaths(t *testing.T) {
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

		t.Log(binaryTreePaths(t1))
	})
}
