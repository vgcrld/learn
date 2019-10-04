
// Simple setup

package main


import "fmt"
import "log"

// Remember the Cap the fields. 
type Square struct {
  Center Point
  Length int
}

type Point struct {
  X int
  Y int
}

func (p *Point) Move(x int, y int) {
  p.X += x
  p.Y += y
}

func NewSquare(x int, y int, len int) (*Square, error) {
  if len <=0 {
    return nil, fmt.Errorf("length must be > 0")
  }

  s := &Square{
    Center: Point{x,y},
    Length: len,
  }

  return s, nil
}

func (s *Square) Move(dx int, dy int) {
  s.Center.Move(dx, dy)
}


func main() {


  fmt.Println("Here we go with Square and Point.")
  sqr, err := NewSquare(1,2,50)
  if err != nil {
    log.Fatalf("Error")
  }
  fmt.Printf("%+v\n", sqr)

  for i := 1; i <= 10; i++ {
    fmt.Println(i)
  }



}
