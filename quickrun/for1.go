// Simple setup

package main


import (
  "fmt"
)

func main() {

  for i := 0; i < 3; i++ {
    fmt.Println(i)
  }

  pl()

  for i := 0; i < 13; i++ {
    if i > 5 && i < 10 { continue }  // or break
    fmt.Println(i)
  }

  pl()

  a := 10
  for a >=  0 {
    fmt.Println(a)
    a--
  }

}

func pl() {
  fmt.Println("-------")
}
