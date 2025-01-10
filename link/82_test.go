package link

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		val := cur.Next.Val
		nxt := cur.Next
		nxxt := cur.Next.Next
		// 若存在两个连续节点相等
		if nxt.Val == nxxt.Val {
			// 递归删除相等节点
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			// 否则cur向下移动
			cur = nxt.Next
		}
	}
	return dummy.Next
}
