// This is a test go program.

package main

import (
	"fmt"
)

func main() {

  x, y := 1, 2

  switch {
  case y == 2:
    fmt.Println("y is 2")
  case x == 1:
      fmt.Println("x is 1")
  default:
    fmt.Println("default")
  }

}
