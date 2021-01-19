package main

import "fmt"

func main() {

	go func() {
		fmt.Println("hello")
	}()

	v, err := fmt.Scanln()
	if err != nil {
		fmt.Println(v)
	}

}
