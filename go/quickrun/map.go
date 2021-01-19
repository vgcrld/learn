
// Simple setup

package main


import (
  "fmt"
)


func main() {

  stock := "AMZN"

  stocks := map[string]float64{
    "AMZN": 1234.3,
    "GOOG": 1223.19,
    "MSFT": 98.61,
  }

  fmt.Println(len(stocks))      // 3
  fmt.Println(stocks["GOOG"])   // 1223.19
  fmt.Println(stocks["DSFDF"])  // 0 - not found
  value,ok := stocks[stock]
  fmt.Println(ok)
  if !ok {
    fmt.Printf("STOCK not found, value %v\n", value)
  } else {
    fmt.Printf("%v found, value %v\n", stock, value)
  }

}
