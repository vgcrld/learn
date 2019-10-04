// Receiver example
package main

import (
	"fmt"
)

// Point is a 2d point
type Point struct {
	X int
	Y int
}

// Reads: add "func" to ( "Point") called "Move" with parms ( int, int ) (no return)
func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
  // p is a reference to point
	p := &Point{1, 2}
	p.Move(2, 3)
	fmt.Printf("%+v\n", p)
}
