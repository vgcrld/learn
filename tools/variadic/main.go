package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum([]int{1, 2, 3}...))
}

func sum(i ...int) int {
	x := 0
	for _, i := range i {
		x += i
	}
	return x
}
