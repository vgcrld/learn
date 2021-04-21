// fmt.Sprint example

package main

import (
	"fmt"
)

// Sprintf used to create a formated string (returns it)
func main() {
	n := 42
	s := fmt.Sprintf("%d", n)

	fmt.Printf("s = %q (type %T)\n", s, s)
}