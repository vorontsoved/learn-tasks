package easy

func singleNumber(nums []int) int {
	dict := make(map[int]bool, 0)

	for _, val := range nums {
		_, ok := dict[val]
		if !ok {
			dict[val] = false
		} else {
			dict[val] = true
		}
	}

	for _, val := range nums {
		value, _ := dict[val]
		if !value {
			return val
		}
	}
	return 0
}
