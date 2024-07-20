package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow := head
	newHead := &ListNode{
		Val: slow.Val,
	}

	for slow.Next != nil {
		slow = slow.Next
		x := &ListNode{
			Val:  slow.Val,
			Next: newHead,
		}

		newHead = x
	}

	return newHead
}
