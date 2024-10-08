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

func Pattern7(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			if i == j {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func Pattern8(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			if i+j == number+1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func Pattern9(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			if i == j || i+j == number+1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func Pattern10(number int) {
	outerSpaces := number / 2
	innerSpaces := -1
	for i := 1; i <= number; i++ {
		// Printing outer spaces
		for j := 1; j <= outerSpaces; j++ {
			fmt.Print(" ")
		}
		// Printing the first '*'
		fmt.Print("*")

		// Printing inner spaces (if any)
		if innerSpaces > 0 {
			for k := 1; k <= innerSpaces; k++ {
				fmt.Print(" ")
			}
			// Printing the second '*' if it's not the top/bottom point
			fmt.Print("*")
		}

		fmt.Println("")

		// Adjust outer and inner spaces based on whether you are in the top or bottom half
		if i <= number/2 {
			outerSpaces--
			innerSpaces += 2
		} else {
			outerSpaces++
			innerSpaces -= 2
		}
	}
}

func Pattern11(number int) {
	counter := 1
	for i := 1; i <= number; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(counter, "\t")
			counter++
		}
		fmt.Println()
	}
}

func Pattern12(number int) {
	a := 0
	b := 1
	for i := 1; i <= number; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(a, "\t")
			c := a + b
			a = b
			b = c
		}
		fmt.Println("")

	}
}

func Pattern13(number int) {
	for i := 0; i < number; i++ {
		icj := 1
		for j := 0; j <= i; j++ {
			fmt.Print(icj, "\t")
			icj1 := icj * (i - j) / (j + 1)
			icj = icj1
		}
		fmt.Println("")
	}
}

func Pattern14(number int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d X %d  = %d \n", number, i, number*i)
	}
}

func Pattern15(number int) {
	stars := 1
	spaces := number / 2
	answer := 1
	for i := 1; i <= number; i++ {
		for j := 1; j <= spaces; j++ {
			fmt.Print(" ")
		}
		innerAnswer := answer
		for k := 1; k <= stars; k++ {
			fmt.Print(innerAnswer)
			if k <= stars/2 {
				innerAnswer++
			} else {
				innerAnswer--
			}
		}

		fmt.Println()
		if i < number/2+1 {
			stars += 2
			spaces -= 1
			answer++
		} else {
			stars -= 2
			spaces += 1
			answer--
		}
	}
}

func Pattern16(number int) {
	spaces := 2*number - 3
	stars := 1

	for i := 1; i <= number; i++ {
		answer := 1
		for j := 1; j <= stars; j++ {
			fmt.Print(answer)
			answer++
		}
		for k := 1; k <= spaces; k++ {
			fmt.Print(" ")
		}

		if i == number {
			stars--
			answer--
		}

		for l := 1; l <= stars; l++ {
			answer--
			fmt.Print(answer)
		}
		stars++
		spaces -= 2
		fmt.Println()
	}
}

func Pattern17(number int) {
	stars := 1
	spaces := number / 2

	for i := 1; i <= number; i++ {

		for k := 1; k <= spaces; k++ {
			if i == number/2+1 {
				fmt.Print("*\t")
			} else {
				fmt.Print("\t")
			}

		}

		for j := 1; j <= stars; j++ {
			fmt.Print("*\t")
		}

		if i <= number/2 {
			stars++
		} else {
			stars--
		}

		fmt.Println()

	}
}

func Pattern18(number int) {
	stars := number
	spaces := 0
	for i := 1; i <= number; i++ {
		for space := 1; space <= spaces; space++ {
			fmt.Print("\t")
		}
		for star := 1; star <= stars; star++ {
			if i > 1 && i <= number/2 && star > 1 && star < stars {
				fmt.Print("\t")
			} else {
				fmt.Print("*\t")
			}

		}
		for space := 1; space <= spaces; space++ {
			fmt.Print("\t")
		}
		if i <= number/2 {
			stars -= 2
			spaces++
		} else {
			stars += 2
			spaces--
		}
		fmt.Println()
	}
}

func Pattern19(number int) {
	for i := 1; i <= number; i++ {
		for star := 1; star <= number; star++ {
			if star > number/2+1 && star < number && i <= number/2 || i > 1 && i <= number/2 && star < number/2+1 || i > number/2+1 && star > 1 && star <= number/2 || i > number/2+1 && star > number/2+1 && i < number {
				fmt.Print("\t")
			} else {
				fmt.Print("*\t")
			}

		}
		fmt.Println()
	}
}

func Pattern20(number int) {

	for i := 0; i <= number; i++ {
		for star := 1; star <= number; star++ {
			if star == 1 || star == number || i == star && i > number/2 || i+star == number+1 && i > number/2 {
				fmt.Print("*\t")
			} else {
				fmt.Print("\t")
			}

		}

		fmt.Println()
	}

}
