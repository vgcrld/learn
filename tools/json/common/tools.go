package common

import (
	"fmt"
)

type Account struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func Hello(s string) {
	fmt.Println("Hello, ", s)
}
