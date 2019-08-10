package main

import (
	"fmt"
)

func main() {

	a := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("%T\n", a)

}

func translateChar(c int) (int, int, int) {
	var a int = 0
	var s int = 0
	return a, c, s
}
