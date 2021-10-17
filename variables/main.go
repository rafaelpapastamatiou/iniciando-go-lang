package variables

import "fmt"

func Describe() {
	declarations()
	types()
	consts()
}

func declarations() {
	fmt.Printf("\n\nDeclaring variables: \n\n")

	a := 10
	fmt.Printf("%v \n", a)

	var b int
	b = 22
	fmt.Printf("%v \n", b)

	var c string = "Hello"
	fmt.Printf("%s \n", c)

	var d, e string = "D", "E"
	fmt.Printf("%s, %s \n", d, e)
}

func types() {
	fmt.Printf("\n\nVar types \n\n")

	a := 10
	b := 15.5
	c := "Hello"
	d := 'H'
	e := true
	f := `Hello with backticks/backquotes`

	fmt.Printf("%T => %v \n", a, a)
	fmt.Printf("%T => %v \n", b, b)
	fmt.Printf("%T => %v \n", c, c)
	fmt.Printf("%T => %v \n", d, d)
	fmt.Printf("%T => %v \n", e, e)
	fmt.Printf("%T => %v \n", f, f)
}

func consts() {
	fmt.Println("\n\nConstants")
	const c1 = 10
	const c2 int = 20

	const (
		c3 = 30
		c4 = 40
		c5 = 50
	)

	fmt.Printf("\nc1 => %v", c1)
	fmt.Printf("\nc2 => %v", c2)
	fmt.Printf("\nc3 => %v", c3)
	fmt.Printf("\nc4 => %v", c4)
	fmt.Printf("\nc5 => %v", c5)

	fmt.Println()
}
