package easy

func missingNumber(nums []int) int {
	missing := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		missing ^= i ^ nums[i]
	}

	missing ^= n

	return missing
}
