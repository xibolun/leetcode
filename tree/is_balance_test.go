package tree

import "math"

func isBalanced(root *TreeNode) bool {
	return getDepth(root) != -1
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := getDepth(root.Left)
	// 若为-1，说明不平衡，正常应该返回值为0
	if leftDepth == -1 {
		return -1
	}
	rightDepth := getDepth(root.Right)
	// 若为-1，或者左右高度不一，说明不平衡
	if rightDepth == -1 || math.Abs(float64(rightDepth)-float64(leftDepth)) > 1 {
		return -1
	}
	// 每次高度需要+1
	return max(rightDepth, leftDepth) + 1
}
