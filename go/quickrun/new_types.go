// Simple setup

package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

var (
	name  Name
	names Names
	age   int
)

type Name string
type Names []Name
type Others []string

func main() {

	names = append(names, "richard")
	names = append(names, "sue")

	spew.Dump(names)

	var xx Names
	for _, y := range []Name{"rich", "sue"} {
		p(y)
		xx = append(xx, y)
	}

	spew.Dump(xx)

}

func p(v interface{}) {
	fmt.Printf("%v\n", v)
}
