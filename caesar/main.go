package main

import (
	"fmt"
)

func main() {

	a := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("%T\n", a)
	fmt.Printf("%s\n", hello("rich"))

}

func translateChar(c int) (int, int, int) {
	var a int = 0
	var s int = 0
	return a, c, s
}

func hello(name string) string {
	return "hi, " + name
}
