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

func PrintRightAlignTriangle(number int) {
	for i := 1; i <= number; i++ {
		star := i
		spaces := number - star

		// Print spaces
		for j := 1; j <= spaces; j++ {
			fmt.Print(" ")
		}

		// Print stars
		for k := 1; k <= star; k++ {
			fmt.Print("*")
		}

		fmt.Println("")
	}
}
