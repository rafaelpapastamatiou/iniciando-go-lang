package functions

import "fmt"

func Describe() {
	fmt.Println("\n\nFunctions")

	fmt.Println("\nSum function")

	fmt.Println(`
func sum(a int, b int) int {
  return a + b
}

a := 10
b := 5

c := sum(a, b)

c == 15
	`)

	fmt.Println("\n\nNamed return")

	fmt.Println(`
func namedReturn(a string) (x string) {
	x = a
	return
}

y := namedReturn("hello")

y == hello
	`)

	fmt.Println("\n\nMore returns")

	fmt.Println(`
func moreReturns(a string, b string) (string, string) {
	return a, b
}	

x, y := moreReturns("hello", "world")

x == hello
y == world
	`)

	fmt.Println("\n\nVariadic Functions")

	fmt.Println(`
func variadicFunc(x ...int) int {
  var result int

  for _, value := range x {
    result += value
  }

  return result
}

sum := variadicFunc(5,5,5)

sum == 15
	`)

	fmt.Println("\n\nAnonymous Functions")

	fmt.Println(`
x := 0

add := func() {
  x += 2
  return x
}

add() == 2
	`)

	fmt.Println("\n\nFunction inside function")

	fmt.Println(`
func funcInsideFunc() func() int {
  x := 10

  return func() int {
    return x * x
  }
}

y := funcInsideFunc()

y == 100
	`)

}
