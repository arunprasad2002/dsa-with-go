package array

func minimumOperations(nums []int) int {
	var one, two int
	for i := 0; i < len(nums); i++ {
		if nums[i]%3 == 1 {
			one++
		} else if nums[i]%3 == 2 {
			two++
		}
	}

	return one + two
}
