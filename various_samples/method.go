
package main


import (
  "fmt"
)

// Create the struct
type Trade struct {
  Symbol string
  Volume int
  Price  float64
  Buy    bool
}

// Add the function to Trade
// Receiver is *Trade, Function is Value, returns float64
func (trade *Trade) Value() float64 {
  value := float64(trade.Volume) * trade.Price
  if trade.Buy {
    value = -value
  }
  return value
}


// Main
func main() {

  // Make the Trade struct
  t := Trade{
    Symbol: "MSFT",
    Volume: 10,
    Price:  99.88,
    Buy:    true,
  }

  // Call value on this
  fmt.Printf( "Stock %s has a value of $%0.2f\n", t.Symbol, t.Value())

}
