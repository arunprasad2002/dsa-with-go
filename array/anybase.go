package array

func Power(num int, power int) int {
	var result int = 1
	if power != 0 {

		// Recursive function call to itself
		result = (num * Power(num, power-1))
	}
	return result
}
func Anybase(number int, base int) int {
	ans := 0
	power := 0
	for number != 0 {
		digit := number % base
		number = number / base
		ans += digit * Power(10, power)
		power++
	}
	return ans

}
