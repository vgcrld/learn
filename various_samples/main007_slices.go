
// Simple setup

package main


import (
  "fmt"
)


func main() {

  names := []int{1,222,3}
  fmt.Println(names[1])    // 222

  // read as: array of strings with { these items in them }
  loons := []string{"bugs","daffy","taz"}


  fmt.Println(len(loons))  // 3
  fmt.Println(loons[0])    // bugs
  fmt.Println(loons[1:])   // daffy taz

  for i := range loons {
    fmt.Println(i)
  }

  for _, name := range loons {
    fmt.Println(name)
  }

  loons = append(loons, "elmer")
  fmt.Println(len(loons))

  for _, name := range loons {
    fmt.Println(name)
  }

}
