package main

import (
	"fmt"

	"example.com/dsa/array"
)

func main() {
	first := []int{1, 23, 44, 5, 56, 6, 6}
	second := []int{10, 20, 33, 4}
	result := array.Merge(first, second)
	fmt.Println(result)
}
