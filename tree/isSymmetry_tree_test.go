package tree

/*
*
二叉树是否对称
首先需要反转一下二叉树，然后和原二叉树再进行比对
*/
func isSymmetry2(first *TreeNode) bool {
	invertTree := invertTree(first)
	return isSameTree(first, invertTree)
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetryNode(root.Left, root.Right)
}
func isSymmetryNode(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSymmetryNode(p.Left, q.Right) && isSymmetryNode(p.Right, q.Left)
}
