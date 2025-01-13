package tree

// [100. 相同的树](https://leetcode.cn/problems/same-tree/description/)
/**
两棵二叉树是否相等
*/
func isSameTree(first *TreeNode, second *TreeNode) bool {
	// if first == nil && second == nil {
	// 	return true
	// }
	// if first == nil || second == nil {
	// 	return false
	// }
	// if first.Val != second.Val {
	// 	return false
	// }
	if first == nil || second == nil {
		return first == second
	}

	return isSameTree(first.Left, second.Left) && isSameTree(first.Right, second.Right)
}
