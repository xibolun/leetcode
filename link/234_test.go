package link

// tag: 链表,回文链表,反转链表,链表中间节点
func isPalindrome(head *ListNode) bool {
	middleNode := middleNode(head)
	head2 := reverseList(middleNode) // 反转中间节点后面的链表

	// 对比两个链表的数值是否完全相等
	for {
		if head2 == nil || head2.Next == nil {
			break
		}

		if head.Val != head2.Val {
			return false
		}

		head = head.Next
		head2 = head2.Next
	}
	return true
}
