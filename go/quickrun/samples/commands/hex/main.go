package main

import "fmt"

func main() {

	for x := 0; x < 100; x++ {

		fmt.Printf("%X is %0b in decimal.\n", x, x)
	}

}
