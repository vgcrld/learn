
// Simple setup

package main


import (
  "fmt"
)

// Create a "type" Trade which is a struct
type Trade struct {
  Symbol string
  Volume int
  Price  float64
  Buy    bool
}


func main() {

  t1 := Trade{"MSFT", 10, 99.98, true}
  fmt.Println(t1)
  fmt.Println(t1.Symbol)
  fmt.Println(t1.Buy)
  fmt.Printf("%+v\n",t1)

}
