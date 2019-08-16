package easy

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

func tree2str(t *TreeNode) string {
	var str bytes.Buffer
	if t == nil {
		return str.String()
	}
	str.WriteString(strconv.Itoa(t.Val))
	if t.Left != nil {
		str.WriteString("(")
		str.WriteString(tree2str(t.Left))
		str.WriteString(")")
	} else if t.Right != nil {
		str.WriteString("()")
	}

	if t.Right != nil {
		str.WriteString("(")
		str.WriteString(tree2str(t.Right))
		str.WriteString(")")
	}

	return str.String()
}

func TestTree2str(t *testing.T) {
	Convey("test", t, func() {
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
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		}
		t2 := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  3,
				Left: nil,
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		}
		t.Log(tree2str(t1))
		t.Log(tree2str(t2))
	})
}
