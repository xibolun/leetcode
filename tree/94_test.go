package tree

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	var source = []struct {
		a   []int
		rse []int
	}{
		{[]int{1, -1, 2, 3}, []int{1, 3, 2}},
		{[]int{}, []int{}},
	}

	for _, v := range source {
		a := buildTree(v.a)
		r := inorderTraversal(a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 二叉树，中序遍历
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)

	loop(root, &ans)
	return ans
}
func loop(t *TreeNode, ans *[]int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		loop(t.Left, ans)
	}
	*ans = append(*ans, t.Val)

	if t.Right != nil {
		loop(t.Right, ans)
	}
}
