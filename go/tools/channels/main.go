/*

 */
package main

import (
	"fmt"
	"math/rand"
)

var (
	firstNames = []string{
		"bob",
		"sue",
		"rich",
		"mark",
	}
	lastNames = []string{
		"McNickle",
		"Davis",
		"Buddy",
		"Soup",
		"Beible",
	}
)

type Employee struct {
	Name  string
	Grade int
}

func main() {

	for i := 1; i <= 10; i++ {
		// r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r := rand.New(rand.NewSource(99))
		fmt.Println(r.Int())
	}
}
