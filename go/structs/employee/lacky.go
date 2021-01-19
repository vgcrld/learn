package employee

import "fmt"

// Lacky - This creates a low level employee/
type Lacky struct {
	Name        string
	Position    string
	Age         int
	Description string
}

func (e Lacky) Display() {
	fmt.Printf("Your age, %s, is %X in hex", e.Name, e.Age)
}
