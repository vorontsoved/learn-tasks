package easy

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	current := dummy

	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return dummy.Next
}
