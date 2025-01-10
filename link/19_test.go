package link

// [19. 删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head} // 这个dummpyNode相当于是一个虚拟头节点，若不创建dummyNode，head本身变更后就无法正常返回
	// 先让right走N
	// 然后l\r同时走，r走到最后的时候，l即为倒数第n个节点
	// 因为他们的差值为n
	right := dummyNode
	left := dummyNode

	for i := 0; i < n; i++ {
		right = right.Next
	}
	for right != nil && right.Next != nil {
		right = right.Next
		left = left.Next
	}
	// 此时的left就是倒数N个节点，删除倒数第N个节点
	left.Next = left.Next.Next
	return dummyNode.Next
}
