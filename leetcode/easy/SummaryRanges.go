package easy

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var elements []string

	if len(nums) == 1 {
		elements = append(elements, strconv.Itoa(nums[0]))
		return elements
	}

	start := nums[0]
	end := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			end = nums[i]
			if i == len(nums)-1 {
				as := fmt.Sprintf("%s->%s", strconv.Itoa(start), strconv.Itoa(end))
				elements = append(elements, as)
			}
		} else {
			if start == end {
				elements = append(elements, strconv.Itoa(start))
				if i == len(nums)-1 {
					elements = append(elements, strconv.Itoa(nums[i]))
				}
				start = nums[i]
				end = nums[i]
			} else {
				as := fmt.Sprintf("%s->%s", strconv.Itoa(start), strconv.Itoa(end))
				elements = append(elements, as)
				start = nums[i]
				end = nums[i]
				if i == len(nums)-1 {
					elements = append(elements, strconv.Itoa(start))
				}
			}
		}
	}

	return elements
}
