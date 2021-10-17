package pointers

import "fmt"

func Describe() {
	fmt.Println("\n\nPointers")

	x := 10
	fmt.Printf("\nx (addr) => %v", &x)
	fmt.Printf("\nx => %v", x)

	y := &x
	*y = 20
	fmt.Printf("\n\ny (addr) => %v", y)
	fmt.Printf("\ny => %v\n", *y)

	fmt.Printf("\nx (addr) => %v", &x)
	fmt.Printf("\nx (after modifying y pointer) => %v\n", x)

	var z *int = &x
	fmt.Printf("\nz (addr) => %v", z)
	fmt.Printf("\nz => %v\n", *z)

	b := xpto(z)
	fmt.Printf("\nb (addr) => %v", b)
	fmt.Printf("\nb => %v\n", *b)
}

func xpto(a *int) *int {
	*a = 100

	return a
}
