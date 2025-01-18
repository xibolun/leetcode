package link

import (
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	var source = []struct {
		a   []int
		rse []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
	}

	for _, v := range source {
		a := &ListNode{}
		rse := &ListNode{}
		buildListNode(v.a, 0, a)
		buildListNode(v.rse, 0, rse)
		r := middleNode(a)
		if !reflect.DeepEqual(r, rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, rse)
		}
	}
}

// - [876. 链表的中间结点](https://leetcode.cn/problems/middle-of-the-linked-list/description/)

// tag: link, 双指针, 快慢指针, 找链表中间节点
// 核心思想：
// 1. 快指针走2步，慢指针走1步
// 2. 快指针走完的时候，慢指针就是中间节点；
// 3. 链表长度为奇数时，快指针最终走到last节点
// 4. 链表长度为偶数时，快指针最终走到空节点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
}
