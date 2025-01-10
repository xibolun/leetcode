package link

// func TestHasCycle(t *testing.T) {
// 	var source = []struct {
// 		a   []int
// 		rse bool
// 	}{
// 		{[]int{3, 2, 0, -4}, true},
// 		{[]int{1, 2}, true},
// 		// {[]int{1}, false},
// 		// {[]int{1, 2}, false},
// 	}

// 	for _, v := range source {
// 		a := &ListNode{}
// 		buildListNode(v.a, 0, a)
// 		r := hasCycle(a)
// 		if !reflect.DeepEqual(r, v.rse) {
// 			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
// 		}
// 	}
// }

// [141. 环形链表](https://leetcode.cn/problems/linked-list-cycle/description/)

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
		// 用指针地址来判断
		if slow == fast {
			return true
		}
	}
	return false
}
