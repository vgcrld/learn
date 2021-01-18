package main

import (
	"fmt"
)

func main() {

	// maps must use make or be created in order to add to them
	var m = make(map[string]interface{})
	m["rich"] = "bob"
	m["bob"] = 1

	printit(m)

	z := make(map[string]interface{})
	println("\nremove data1")
	z["data1"] = 1
	z["data2"] = "richard"
	delete(z, "data1")

	printit(z)
}

func printit(m map[string]interface{}) {
	fmt.Println("\nhello")
	for k, v := range m {
		fmt.Printf("%T %v\n", k, v)
	}
}
