package main

import "fmt"

func main() {

	a := 5
	b := &a

	fmt.Println("a=", a, "b=", b)
	fmt.Println("*b=", *b)
	*b = 10 // we can update a by dereferencing b
	fmt.Println(*b)
	fmt.Println(a)

}
