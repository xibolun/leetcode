package link

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	var source = []struct {
		a   []int
		rse []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, v := range source {
		a := &ListNode{}
		rse := &ListNode{}
		buildListNode(v.a, 0, a)
		buildListNode(v.rse, 0, rse)

		if !reflect.DeepEqual(a, rse) {
			t.Errorf("test fail, result: %v, should be :%v", a, rse)
		}
	}
}
func buildListNode(v []int, index int, a *ListNode) {
	if index > len(v)-1 {
		return
	}
	if a == nil {
		a = &ListNode{Val: v[index]}
	}
	buildListNode(v, index+1, a.Next)
}

// 206. 反转链表
// https://leetcode.cn/problems/reverse-linked-list/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode

	for cur.Next != nil {
		// 将next值先记录
		nxt := cur.Next
		// next的指向要向前，而不是向后
		cur.Next = pre
		// 修改pre的内存指针地址为cur
		pre = cur
		// 将下一节点放到cur来遍历，获取Next
		cur = nxt
	}
	return pre
}
