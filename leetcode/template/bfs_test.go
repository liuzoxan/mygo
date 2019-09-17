package template

import "testing"

func bfs(root *TreeNode, visited *[]int) {
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		*visited = append(*visited, node.value)

		// process the node

		// generate other nodes
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

}

func TestBFS(t *testing.T) {
	root := &TreeNode{
		value: 1,
		left:  nil,
		right: nil,
	}
	visited := &[]int{}

	bfs(root, visited)

	t.Log([]int{1,2,3}[:])
	t.Log([]int{1,2,3}[2:])
	t.Log([]int{1,2,3}[3:])
	t.Log([]int{1,2,3}[:3])
	t.Log([]int{1,2,3,4}[0:3])
}
