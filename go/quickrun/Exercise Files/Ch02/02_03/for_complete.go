// "for" loop examples
package main

import (
	"fmt"
)

// Iteration
func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

  // break will leave a for loop
	fmt.Println("----")
	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

  // continue with go to next iteration
	fmt.Println("----")
	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

  // Breaks up the for components
	fmt.Println("----")
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

  // for { is like a while true loop you must break it
	fmt.Println("----")
	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}
}
