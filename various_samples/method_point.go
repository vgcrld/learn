package main

import "fmt"

type Point struct {
  X int
  Y int
}

// *Point is the receiver
func (p *Point) Move(dx int, dy int) {
  p.X += dx
  p.Y += dy
}

func main() {

  // make p a pointer to a Point so we can change the contents
  p := &Point{1,2}

  fmt.Printf("%+v\n",p)

  p.Move(2,3)
  fmt.Printf("%+v\n",p)

  p.Move(12,3003)
  fmt.Printf("%+v\n",p)
}
