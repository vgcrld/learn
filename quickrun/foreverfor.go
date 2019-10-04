// Simple setup

package main

import (
	"fmt"
)

func main() {

	a := 0

	//Here we can see how to run forever. Just for{}
	for {

		fmt.Printf("%4d ", a)
		a++
		if a > 10000 {
			break
		}

	}

}
