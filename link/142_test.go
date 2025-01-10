package link

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next

		// 他们相遇时，head和slow也会在相遇点的地方相遇
		// https://www.bilibili.com/video/BV1KG4y1G7cu?spm_id_from=333.788.videopod.sections
		if fast == slow {
			for {
				if head == slow {
					return slow
				}
				head = head.Next
				slow = slow.Next
			}
		}
	}
	return nil
}

// 用Hash的方法
func detectCycle2(head *ListNode) *ListNode {
	mp := make(map[*ListNode]bool)
	for head != nil {
		if mp[head] {
			return head
		}
		mp[head] = true
		head = head.Next
	}
	return nil
}
