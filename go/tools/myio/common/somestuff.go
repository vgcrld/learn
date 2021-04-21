package common

import (
	"fmt"
)

func init() {
	fmt.Println("Hi - Welcome frim init()")
}
func DoSomethingRandom(s string) string {
	return "hello, " + s
}
