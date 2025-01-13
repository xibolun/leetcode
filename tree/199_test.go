package tree

// https://leetcode.cn/problems/binary-tree-right-side-view/description/
func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	getVal(root, 0, ans)
	return ans
}

func getVal(root *TreeNode, depth int, ans []int) {
	if root == nil {
		return
	}

	// 若元素个数与深度相同，说明每层有且只有一个节点加入了ans，所以这个元素并未加入到ans当中，直接写入即可
	if len(ans) == depth {
		ans = append(ans, root.Val)
	}
	getVal(root.Right, depth+1, ans)
	getVal(root.Left, depth+1, ans)
}

// 这个函数的错误在于未考虑深度，若左侧的深度比右侧深，那么左侧的叶子节点也需要加到右视图当中
func getRightVal(root *TreeNode, ans []int) {
	if root == nil {
		return
	}
	ans = append(ans, root.Val)
	getRightVal(root.Right, ans)
}
