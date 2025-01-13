package tree

import "math"

// [98. 验证二叉搜索树](https://leetcode.cn/problems/validate-binary-search-tree/description/)

// 前序遍历
func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func dfs(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}
	x := root.Val
	return (x > left && x < right) && dfs(root.Left, left, x) && dfs(root.Right, x, right)
}

// 这个算法的错误在于未将根节点传入，只判断每个root的左右是否符合二叉搜索树的特点
func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil {
		return root.Val < root.Right.Val
	}

	if root.Right == nil {
		return root.Val > root.Left.Val
	}
	return root.Left.Val < root.Val && root.Val < root.Right.Val && isValidBST(root.Left) && isValidBST(root.Right)
}
