package template

import "testing"

func dfs(root *TreeNode, visited *[]int) {
	*visited = append(*visited, root.value)
	// process current node here

	if root.left != nil {
		// if root.left.value not in visited
		dfs(root.left, visited)
	}

	if root.right != nil {
		// if root.right.value not in visited
		dfs(root.right, visited)
	}
}

func TestDFS(t *testing.T) {
	visited := &[]int{}
	root := &TreeNode{}
	dfs(root, visited)
}