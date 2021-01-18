package main

import (
	"fmt"
	"math"
)

type Maths interface {
	Multiply() float64
	Raise() float64
}

type Something struct {
	a, i float64
}

type Otherthing struct {
	a, i float64
}

func (o Something) Multiply() float64 {
	return o.a * o.i
}

func (o Something) Raise() float64 {
	return math.Pow(o.a, o.i)
}

func (o Otherthing) Multiply() float64 {
	return o.a * o.i

}

func (o Otherthing) Raise() float64 {
	return math.Pow(o.a, o.i)
}

func main() {

	vals := []Maths{
		Otherthing{2, 4},
		Otherthing{2, 4},
		Otherthing{2, 4},
		Otherthing{2, 4},
	}

	for _, v := range vals {
		fmt.Println(v.Raise())
		fmt.Println(v.Multiply())
	}
}
