package tree

// [513. 找树左下角的值](https://leetcode.cn/problems/find-bottom-left-tree-value/description/)
// *** 需要重点思考一下
// 用队列的方式
func findBottomLeftValue(root *TreeNode) int {
	node := root
	q := []*TreeNode{node}
	for len(q) > 0 {
		// node为第一个节点
		// q[1:]为其他节点
		node, q = q[0], q[1:]

		// 一定要先push右节点
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}
	return q[0].Val
}
func findBottomLeftValue2(root *TreeNode) int {
	row := []*TreeNode{root} // 每一行的节点
	isLastRow := false
	for len(row) > 0 {
		if isLastRow {
			return row[0].Val
		}
		nxt := []*TreeNode{}
		for _, node := range row {
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
			isLastRow = (isLastRow && node.Left == nil && node.Right == nil)
		}
		row = nxt
	}
	return 0
}
