package main

import (
	"fmt"
)

// Make me a people
type people struct {
	Name string
	Age  int
}

// add this function to people
func (t *people) getFull() string {
	name := t.Name
	age := fmt.Sprint(t.Age)
	full := name + " " + age
	return full
}

// Start here
func main() {

	// Create a type people
	person := people{"rich", 44}

	full := person.getFull()

	full = person.getFull()

	person.Name = "bob"
	fmt.Println(person)

	fmt.Printf("%s", full)

}
