package patterns

import "fmt"

func PrintTriangle(number int) {
	for i := 0; i <= number; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func PrintTriangleReverse(number int) {
	for i := number; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
