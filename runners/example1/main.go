package main

import (
	"fmt"
)

// Make me a people
type people struct {
	First string
	Last  string
	Age   string
}

// add this function to people
func (t *people) getFullName() string {
	full := t.Last + ", " + t.First + ", Age: " + t.Age
	return full
}

// Start here
func main() {

	// Create a type people
	person1 := people{"Rich", "Davis", "51"}
	person2 := people{"Rich", "Davis", "55"}

	fmt.Println(person1.getFullName())
	fmt.Println(person2.getFullName())

}
