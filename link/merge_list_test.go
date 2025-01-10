package link

/*
*
合并两个有序的数组，
1,2,3,5 merge 2,4,6 = 1,2,2,3,4,5,6
*/
func mergeTwoList(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return nil
	}

	if list1.Val < list2.Val {
		mergeTwoList(list1.Next, list2)
		return list1
	}

	list2.Next = mergeTwoList(list1.Next, list2)
	return list2
}
