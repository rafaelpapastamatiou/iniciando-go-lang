package visibility

import "fmt"

// visible to another packages because the initial character V is capitalized
func VisibleFunc() {
	fmt.Println("\nI'm visible")
}

// invisible to another packages because the initial character v is not capitalized
func invisibleFunc() {
	fmt.Println("\nI'm invisible")
}
