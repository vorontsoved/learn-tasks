package easy

func findDisappearedNumbers(nums []int) []int {
	var q []int
	var dict = make(map[int]bool, 1)
	for _, val := range nums {
		dict[val] = true
	}

	for i, end := 1, len(nums); i <= end; i++ {

		_, ok := dict[i]

		if !ok {

			q = append(q, i)
		}
	}

	return q
}
