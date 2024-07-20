package easy

type NumArray struct {
	array []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		array: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if right < left {
		return 0
	}
	sum := 0

	for i := left; i <= right; i++ {
		sum = sum + this.array[i]
	}
	return sum
}
