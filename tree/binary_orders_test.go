package tree

import (
	"fmt"
	"testing"
)

func TestThreeOrders(t *testing.T) {
	var root = &TreeNode{
		Left: &TreeNode{
			Left:  nil,
			Right: nil,
			Level: 0,
			Val:   2,
		},
		Right: &TreeNode{
			Left:  nil,
			Right: nil,
			Level: 0,
			Val:   3,
		},
		Level: 0,
		Val:   1,
	}

	threeOrders(root)

	fmt.Println(pres)
	fmt.Println(ins)
	fmt.Println(posts)
}

/*
*
对{1,2,3}的二叉树进行排序
先序、中序、后序
*/
var pres, ins, posts []int

func threeOrders(root *TreeNode) [][]int {
	pres, ins, posts = []int{}, []int{}, []int{}
	visit(root)
	return [][]int{pres, ins, posts}
}

func visit(node *TreeNode) {
	if node != nil {
		pres = append(pres, node.Val)
		visit(node.Left)
		ins = append(ins, node.Val)
		visit(node.Right)
		posts = append(posts, node.Val)
	}
}
