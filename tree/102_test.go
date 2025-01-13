package tree

// [102. 二叉树的层序遍历](https://leetcode.cn/problems/binary-tree-level-order-traversal/description/)
func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	row := []*TreeNode{root} // 每一行的节点
	for len(row) > 0 {
		val := []int{}
		nxt := []*TreeNode{}
		for _, node := range row {
			val = append(val, node.Val)
			// 将当前行所有的左右非空节点加入到下一行当中
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}
		ans = append(ans, val)
		row = nxt
	}
	return ans
}
