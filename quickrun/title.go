// Simple setup

package main

import (
	"fmt"
	"strings"
)

func main() {

	// Make the title case
	s := "rich davis is going to the store"
	fmt.Println(strings.Title(s))

	// Function Practice
	fmt.Println(add(1, 2))
	fmt.Println(times(100.3, 2.3423423))

	fmt.Println(addx(10, 11))
}

func addx(a int, b int) (int, int) {
	return (a + b), (a * b)
}

func add(a int, b int) int {
	return a + b
}

func times(a float64, b float64) float64 {
	return a * b
}
