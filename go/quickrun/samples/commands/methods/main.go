package main

import "fmt"

/*
	This is how to create a method. The method is not a function.
	a Method is a function attached to a structure. (that may be wrong)
	but IS basically what I see happening here.

	First we create the struct
	Then the code after "func" name is the method.

*/

//declared the structure named emp
// notice emp is the structure
type emp struct {
	name    string
	address string
	age     int
}

//Declaring a function with receiver of the type emp
//The receiver is emp "e" so when you call e.display
//it knows which struct data to use
func (e emp) display() {
	fmt.Printf("Thank you, %s, of age: %d\n", e.name, e.age)
}

func main() {

	//declaring a variable of type emp and assign values to members
	empdata1 := emp{"Raj", "Building-1, Paris", 25}
	empdata2 := emp{"Bob", "Building-1, Paris", 25}

	//Invoking the method using the receiver of the type emp
	// syntax is variable.methodname()
	empdata1.display()
	empdata2.display()
}
