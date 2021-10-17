package scope

import "fmt"

var y int = 20

func Describe() {
	fmt.Println("\n\nScope")
	x := 10
	fmt.Printf("\nx => %v", x)

	printX()
	printY()
	printZ()

	printYFromOther()
	printZFromOther()
}

func printX() {
	// error	fmt.Println(x)
	fmt.Println("\nCan't print X")
}

func printY() {
	fmt.Printf("\ny => %v", y)
}

func printZ() {
	fmt.Printf("\nz => %v", z)
}

var z int = 30
