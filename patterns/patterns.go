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
