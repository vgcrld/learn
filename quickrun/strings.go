
// Simple setup

package main


import (
  "fmt"
)


func main() {

  book := "The color of magic"
  fmt.Println(book)
  fmt.Println(len(book))
  fmt.Printf("book[0] = %v (type %T)\n",book[0],book[0])

  // Strings are Immutable - this won't work
  //book[0] = 160 

  // Slice include 4 to 10
  fmt.Println(book[4:11])

  // Slice include 4 to end
  fmt.Println(book[4:])

  // Slice include start to 4 (includes) 'The '
  fmt.Printf("'%s'\n",book[:4])

  // + to concat
  booknew := book + "/" + book
  fmt.Println(booknew)

  // Using fmt.Sprintf
  n := 42
  s := fmt.Sprintf("%d", n)
  fmt.Printf("s = %q (type %T)\n", s, s)
}
