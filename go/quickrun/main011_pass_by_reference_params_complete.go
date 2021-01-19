package main

import (
	"fmt"
)

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

// notice "*" on params and on *n. 
// dereference this
func doublePtr(n *int) {
	*n *= 2
}

func main() {
	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values)

	val := 10
	double(val)
	fmt.Println(val)

  // Here we send val in with &val to send in the pointer
  // &val allows us to change val in the routine
	doublePtr(&val)
	fmt.Println(val)
}
