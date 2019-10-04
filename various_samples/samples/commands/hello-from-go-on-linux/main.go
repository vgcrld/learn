package main

/*
	Compile this stuff to run on windows like this

	GOOS=windows GOARCH=386 go build -o ~/compiled/hello-from-go-on-linux.exe hello-from-go-on-linux.go

*/

// You can run compile this to run on windows with

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("Hello Windows, Love Linux - %v\n", i)
	}

}
