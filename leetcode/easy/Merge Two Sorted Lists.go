package easy

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	arr := make([]int, 0, 50)

	s1 := list1
	s2 := list2

	for s1 != nil {
		arr = append(arr, s1.Val)
		if s1.Next != nil {
			s1 = s1.Next
		} else {
			s1 = nil
		}
	}

	for s2 != nil {
		arr = append(arr, s2.Val)
		if s2.Next != nil {
			s2 = s2.Next
		} else {
			s2 = nil
		}
	}

	sort.Ints(arr)

	dummy := &ListNode{}
	cur := dummy

	for _, v := range arr {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	return dummy.Next
}
