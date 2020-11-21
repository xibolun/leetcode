package tree

/**
二叉树是否对称
首先需要反转一下二叉树，然后和原二叉树再进行比对
*/
func isSymmetry(first *TreeNode) bool {
	invertTree := invertTree(first)
	return isSameTree(first, invertTree)
}
