
// Simple setup

package main


import "fmt"


// Main start here
func main() {

  values := []int{1,2,3,4,5}
  ret := doubleAt(values,2)
  if ret {
    fmt.Println("It's true")
  }

  fmt.Println(values)


}

// A fucntion to go
func doubleAt(values []int, i int) bool {
  values[i] *= 2
  return false
}

