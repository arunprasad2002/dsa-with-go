package array

func DigitFrequency(number int, target int) int {
	count := 0
	for number != 0 {
		digit := number % 10
		number = number / 10
		if digit == target {
			count++
		}

	}
	return count
}
