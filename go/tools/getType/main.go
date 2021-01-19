package main

import "fmt"

func main() {
	var i interface{} = "a"
	fmt.Println(getType(i))
}

func getType(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
func loopString() {
	a := "hello,richard"

	for k, v := range a {
		fmt.Printf("%5v %5v %5v\n", k, v, string(v))
	}
}
