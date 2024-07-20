package easy

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i, end := 0, len(nums)-1; i < end; i++ {
		if nums[i] == nums[i+1] {
			return true
		}

	}
	return false
}
