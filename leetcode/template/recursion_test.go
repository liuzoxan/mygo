package template

import "log"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func recursion(level int, root *TreeNode) {
	// recursion terminator
	if level < 0 || root == nil {
		return
	}

	// process logic in current level
	log.Println(level, root.value)

	// drill down
	recursion(level-1, root.left)
	recursion(level-1, root.right)

	// reverse the current level status is needed
	// reverse_state(level)
	return
}
