package main

import (
	"fmt"
)

func main() {

	aa := map[string]int{}

	aa["rich"] = 1
	aa["bob"] = 100

	fmt.Printf("%T %v\n", aa, aa)
}
