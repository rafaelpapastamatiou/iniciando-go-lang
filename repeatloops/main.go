package repeatloops

import "fmt"

func Describe() {
	fmt.Println("\n\nRepeat loops")

	fmt.Printf("\nCommon for\n\n")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("\nWhile")
	x := 0
	for x < 4 {
		fmt.Printf("\nx => %v", x)
		x++
	}

	fmt.Println("\n\nUntil specified condition(s)")
	y := 0
	for {
		if y == 6 {
			break
		}

		fmt.Printf("\ny => %v", y)

		y++
	}
}
