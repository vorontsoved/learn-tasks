package easy

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	var data []int
	for cur := head; cur != nil; cur = cur.Next {
		data = append(data, cur.Val)
	}

	for i, j := 0, len(data)-1; i <= j; {
		if data[i] != data[j] {
			return false
		}
		i++
		j--
	}
	return true
}
