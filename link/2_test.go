package link

// tag: 链表
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//不需要反转，因为结果也是反的
	var tmp, ans *ListNode

	carry := 0 // 进位，满十进一，否则进0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0 // 两个元素的值
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {

		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		if tmp == nil {
			tmp = &ListNode{Val: sum}
			ans = tmp
		} else {
			tmp.Next = &ListNode{Val: sum}
			tmp = tmp.Next
		}
	}
	if carry > 0 {
		tmp.Next = &ListNode{Val: carry}
	}
	return ans
}
