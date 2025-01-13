package tree

//  [103. 二叉树的锯齿形层序遍历](https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/)

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	row := []*TreeNode{root} // 每一行的节点
	isEven := false          //是否偶数行，若为偶数行，则需要反转数组
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
		if isEven {
			reverse(val)
		}
		ans = append(ans, val)
		row = nxt
		isEven = !isEven
	}
	return ans
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
