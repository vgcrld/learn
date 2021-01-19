
// Simple setup

package main


import (
  "fmt"
)


func main() {

  // be expicit
  var i int
  i = 1
  pp(i)

  // Let it figure it out, must be a new var
  x := 2
  pp(x)

  // Of course a literal is ok
  pp(3)

  x = 5
  double(&x) // Pass by pointer
  fmt.Printf("The value is %v\n", x)

  // Return multiple values. 
  val, err := bigger10(2,2)
  if err != nil {
    fmt.Printf("Bad, %v\n",err)
  } else {
    fmt.Printf("OK %v\n", val)
  }

  // Don't "panic" use error and check this way
  val, err = errorReturn(2)
  if err != nil {
    fmt.Println("ERROR")
  } else {
    fmt.Println("OK")
  }

}

func errorReturn(val int) (int, error) {
  if val > 5 {
    return val, fmt.Errorf("This is an error %v", val)
  } else {
    return val, nil
  }
}

// Pass in values are strict
func pp( val int ) {
  fmt.Printf("The value is %v\n", val)
}

// Change
func double( val *int ) {
  // *val is dereferenced - it's not really the val here but an address
  *val *= 2
}

// Function can return more than one value
// If a function can error than the 2nd val
// Will be the error - and error is a type
func bigger10(x int, y int) (int, error) {
  if x * y > 10 {
    return (x*y), nil
  } else {
    return 0, fmt.Errorf("error, (%f) not bigger than 10", (x*y))
  }
}

