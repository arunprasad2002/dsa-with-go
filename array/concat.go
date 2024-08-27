package array

func GetConcatenation(nums []int) []int {
	result := make([]int, 2*len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = nums[i]
		result[len(nums)+i] = nums[i]
	}
	return result
}
