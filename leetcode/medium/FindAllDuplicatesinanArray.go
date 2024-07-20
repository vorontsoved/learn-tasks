package medium

func findDuplicates(nums []int) []int {
	a := make(map[int]bool, len(nums))

	uniqle := make([]int, 0, len(nums))

	for _, val := range nums {
		_, ok := a[val]
		if ok {
			a[val] = true
			continue
		}
		a[val] = false
	}

	for key, value := range a {
		if value {
			uniqle = append(uniqle, key)
		}
	}

	return uniqle

}
