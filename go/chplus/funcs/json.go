package funcs

import (
	"encoding/json"
	"fmt"
)

type Size uint

// iota starts as 0 and is increased for each new use.
// - will absorb the 0 or any value that is next
// Basocally it is used for new call inside the const
const (
	_     Size = iota
	Small      = iota * 10
	Medium
	Large
	Extra
)

const (
	a = iota
	b
	c
)

func RunJSON() {
	fmt.Println("Hello from JSON")

	fmt.Println("iota Stuff")
	fmt.Println("Size ", Small, Medium, Large, Extra)
	fmt.Println("restarted ", a, b, c)

	simpleUnmarshal()
}

/*
	Unmarshal and checking some types on the way in. Convert string / float back and forth:
	https://www.derpturkey.com/golang-unmarshal-json-fields/
*/

// You marshal a struct into JSON.
func marsh() {

}

// You unmarshal JSON into a struct
func unmarsh() {

}

func simpleUnmarshal() {
	xx := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}

	data := []byte(`{
	  "name": "richard",
	  "age": 50
	}`)

	err := json.Unmarshal(data, &xx)
	if err != nil {
		panic(err)
	}

	fmt.Println("xx name: ", xx.Name)
	fmt.Println("xx age: ", xx.Age)

}
