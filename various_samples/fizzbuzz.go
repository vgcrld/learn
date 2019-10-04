// If divisible by 3 then fizz
//              by 5 then buzz
//              by 3 && 5 then buzz

package main

import (
  "fmt"
)

func main() {

  for a := 1; a <= 20; a++ {

    switch {
    case a / 3 == 0:
      fmt.Println(a)
    case a%3 == 0 && a%5 == 0:
      fmt.Println("fizz buzz")
    case a%3 == 0:
      fmt.Println("fizz")
    case a%5 == 0:
      fmt.Println("buzz")
    default:
      fmt.Println(a)
    }

  }

}

