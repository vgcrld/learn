
// Simple setup

package main


import (
  "fmt"
  "strings"
)


func main() {

  // Some Text
  text := `
  Needles and pins
  Needles and pins
  Sew me a sail
  To catch me the wind
  rich rich rich rich rich rich rich
  `

  fmt.Println(text)

  words  := strings.Fields(text)
  counts := map[string]int{}   //empty map

  fmt.Println(len(words))

  for _,word := range words {
    counts[strings.ToLower(word)]++
  }

  fmt.Println(counts)

}
