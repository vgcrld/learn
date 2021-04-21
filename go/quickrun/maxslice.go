
// Simple setup

package main


import (
  "fmt"
)


func main() {


  sl  := []int{1,33,54,2322,1,2,12,13}
  max := sl[0]

  for _,val := range sl[1:] {
    if val > max { max = val }
  }

  fmt.Println(max)



}
