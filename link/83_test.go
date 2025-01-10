package link

// [83. 删除排序链表中的重复元素](https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur != nil && cur.Next != nil {
		nxt := cur.Next
		// 若cur的值和下一节点的值相同，则删除下一节点
		if cur.Val == nxt.Val {
			cur.Next = nxt.Next
		} else {
			// 否则向下移动一个节点
			cur = nxt
		}
	}
	return head
}
