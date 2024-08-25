package main

import (
	"fmt"

	"example.com/dsa/search"
)

func main() {
	numbers := []int{1, 23, 44, 5, 56, 6, 6}
	target := 1
	ans := search.BinarySearch(numbers, target)
	fmt.Println(ans)
}
