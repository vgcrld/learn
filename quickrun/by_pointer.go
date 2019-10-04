
// Simple setup

package main


import (
  "fmt"
)


func main() {

  val := 2
  doublePointer(&val)
  fmt.Println(val)

}

// Here we use val by pointer  *int 
func doublePointer(val *int){
  *val *= 2 // de-ref the val
}

