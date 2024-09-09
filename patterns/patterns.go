package patterns

import (
	"fmt"
	"strings"
)

// PrintTriangle prints a right-angled triangle of '*' characters.
func PrintTriangle(number int) {
	for i := 1; i <= number; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}

// PrintTriangleReverse prints an inverted right-angled triangle of '*' characters.
func PrintTriangleReverse(number int) {
	for i := number; i >= 1; i-- {
		fmt.Println(strings.Repeat("*", i))
	}
}
