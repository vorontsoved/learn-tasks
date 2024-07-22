package easy

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)

	if target >= letters[n-1] {
		return letters[0]
	}

	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)/2

		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return letters[left]
}
