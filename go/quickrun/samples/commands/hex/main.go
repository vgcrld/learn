package main

import "fmt"

func main() {

	for x := 0; x <= 255; x++ {

		fmt.Printf("%v is %X in hex and %0b in binary.\n", x, x, x)
	}

}
