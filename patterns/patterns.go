package patterns

import (
	"fmt"
	"strings"
)

// PrintTriangle prints a right-angled triangle of '*' characters.
func Pattern1(number int) {
	for i := 1; i <= number; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}

// PrintTriangleReverse prints an inverted right-angled triangle of '*' characters.
func Pattern2(number int) {
	for i := number; i >= 1; i-- {
		fmt.Println(strings.Repeat("*", i))
	}
}

func Pattern3(number int) {
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

func Pattern4(number int) {
	for i := number; i >= 1; i-- {
		spaces := number - i
		stars := i

		for space := 1; space <= spaces; space++ {
			fmt.Print(" ")
		}
		for star := 1; star <= stars; star++ {
			fmt.Print("*")
		}

		fmt.Println("")
	}
}

func Pattern5(number int) {
	spaces := number / 2
	stars := 1

	for i := 1; i <= number; i++ {

		for j := 1; j <= spaces; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= stars; k++ {
			fmt.Print("*")
		}
		fmt.Println("")

		if i <= number/2 {
			spaces--
			stars += 2
		} else {
			spaces++
			stars -= 2
		}
	}

}

func Pattern6(number int) {
	stars := number/2 + 1
	spaces := 1
	for i := 1; i <= number; i++ {
		for j := 1; j <= stars; j++ {
			fmt.Print("*")
		}
		for k := 1; k <= spaces; k++ {
			fmt.Print(" ")
		}

		for l := 1; l <= stars; l++ {
			fmt.Print("*")
		}

		fmt.Println("")

		if i > number/2 {
			stars++
			spaces -= 2
		} else {
			stars--
			spaces += 2
		}

	}
}
