package link

// tag: link, medium
// 题解思路：按灵神的思路
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB

	for {
		if p == q {
			break
		}
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}
		if q != nil {
			q = q.Next
		} else {
			q = headA
		}

	}
	return p
}
