package main

import (
	"fmt"

	"github.com/vgcrld/learngo/structs/employee"
)

// Person is a person
type Person struct {
	FirstName string
	LastName  string
	age       int
	gender    string
}

//Greet is a way to say hello
func (p Person) Greet() string {
	return "Hello my name is " + p.FirstName
}

func main() {

	x := employee.Lacky{Name: "Rich"}
	fmt.Println(x)
	people := []Person{
		Person{"rich", "davis", 49, "M"},
		Person{"bob", "davis", 53, "M"},
		Person{"mark", "davis", 45, "M"},
		Person{"sue", "davis", 51, "M"},
	}

	for _, person := range people {
		fmt.Printf("%d\n", person.age)
	}
}
