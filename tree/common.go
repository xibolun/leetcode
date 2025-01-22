package tree

// TreeNode
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Level int
	Val   int
}

func buildTree(values []int) *TreeNode {
	if len(values) <= 0 {
		return nil
	}
	rootVal := values[0]
	if rootVal == -1 {
		return nil
	}
	root := &TreeNode{Val: rootVal}

	values = values[1:]
	root.Left = buildTree(values)

	if len(values) < 1 {
		return root
	}

	values = values[1:]
	if len(values) > 0 {
		root.Right = buildTree(values)
	}
	return root
}
