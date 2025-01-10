package link

func reorderList(head *ListNode) {
	midNode := middleNode(head) // 利用快慢指针找到中间节点
	// 分隔链表为两个
	l1, l2 := splitLink(head, midNode)
	// 反转后面一个链表
	reverseList(l2)
	// 合并两个链表，交替连接节点
	head = mergetLink(l1, l2)
}

func splitLink(head, splitNode *ListNode) (*ListNode, *ListNode) {
	l1 := head
	curNode := head
	for curNode != splitNode {
		next := curNode.Next
		l1.Next = next
		curNode = next
	}
	return l1, splitNode
}

func mergetLink(l1, l2 *ListNode) *ListNode {
	for l1 != nil && l2 != nil {
		l1Next := l1.Next
		l2Next := l2.Next

		l1.Next = l2
		l2.Next = l1

		l1 = l1Next
		l2 = l2Next
	}
	return l1
}
