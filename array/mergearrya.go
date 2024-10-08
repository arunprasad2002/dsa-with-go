package array

func Merge(first []int, second []int) []int {
	var result []int
	for _, value := range first {
		result = append(result, value)
	}

	for _, value := range second {
		result = append(result, value)
	}

	return result
}
