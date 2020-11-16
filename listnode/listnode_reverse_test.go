package listnode

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestReverseList(t *testing.T) {

}

func ReverseList(pHead *ListNode) *ListNode {
	var prev *ListNode
	for pHead != nil {
		prev, pHead, pHead.Next = pHead, pHead.Next, prev
	}
	return prev
}
