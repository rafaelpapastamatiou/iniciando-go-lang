package conditionals

import "fmt"

func Describe() {
	fmt.Println("\n\nConditionals")
	fmt.Println()

	a := 10

	if a > 10 {
		fmt.Println("a > 10")
	} else if a > 5 {
		fmt.Println("a > 5")
	} else {
		fmt.Println("a < 5 < 10")
	}

	b := true

	if x := "Cool"; b {
		fmt.Println(x)
	}
}
