package switchcondition

import "fmt"

func Describe() {
	fmt.Println("\n\nSwitch")
	fmt.Println()

	letter := "a"

	switch letter {
	case "a":
		fmt.Println("letter === a")
	case "b":
		fmt.Println("letter === b")
	default:
		fmt.Println("Unrecognized letter")
	}
}
