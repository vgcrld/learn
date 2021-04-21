// This is a test go program.

package main

import (
	"fmt"
)

var (
	ii int
	xx int
)

const (
	Richard = "Davis"
)

func main() {

	xx = 10
	ii = 20
	fmt.Println(xx, ii, Richard)

	x, y := 10.0, 11.2
	z := 10

	fmt.Printf("x is %-0.2f and y is %-0.2f\n", x, y)

	switch z {
	case 1:
		fmt.Printf("%v is one\n", z)
	case 2:
		fmt.Printf("%v is one\n", z)
	default:
		fmt.Printf("%v is bigger then the max allowed!\n", z)
	}

	a, b := 10.0, 15.0

	if frac := a / b; frac > 0.5 {
		fmt.Printf("%0.2f frac is > 0.5\n", frac)
	} else {
		fmt.Printf("%0.2f frac is not > 0.5\n", frac)
	}

}
