// Example of "switch" statement

package main

import (
	"fmt"
)

func main() {
	x := 2

  // Case x directly
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Printf("many")
	}

  // Case individual expressions. Only one is executed.
	switch {
	case x > 100:
		fmt.Println("x is very big")
	case x > 10:
		fmt.Println("x is big")
	default:
		fmt.Println("x is small")
	}

}
