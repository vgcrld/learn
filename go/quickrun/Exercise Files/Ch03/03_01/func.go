// Basic function definition
package main

import (
	"fmt"
)

// add adds a to b; returns int
func add(a int, b int) int {
	return a + b
}

// divmod returns quotient and reminder
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

// You can't change values of basic types without a reference
func main() {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Printf("div=%d, mod=%d\n", div, mod)
}
