package main

import (
	"fmt"

	"github.com/vgcrld/learngo/employee"
)

func main() {

	var emp1 employee.Lacky

	emp1.Name = "Rich"
	emp1.Position = "Nobody"
	emp1.Description = "A guy that is cool."
	emp1.Age = 49

	fmt.Println("hello")
	fmt.Println(emp1.Name)
	fmt.Println(employee.Salary(115343.00, .10))

}
