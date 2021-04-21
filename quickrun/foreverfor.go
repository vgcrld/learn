// Simple setup

package main

import (
	"fmt"
	"time"
)

func main() {

	a := 0

	//Here we can see how to run forever. Just for{}
	for {

		fmt.Printf("%8d ", a)

		//use this code to break out
		a++
		if a > 10000 {
			break
		}

		time.Sleep(time.Millisecond * 10)

	}

}
