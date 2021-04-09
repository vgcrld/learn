// Simple setup

package main

import (
	"fmt"
	// "strings"
)

/*
  Maps are basically hashes but more confusing than hashes
  you create it with map[key_type]val_type{}
  It pretty stupid syntax.
*/

func main() {

	my_map := map[string]int{}
	my_map["richard"] = 1
	my_map["davis"] = 15
	for i, val := range my_map {
		fmt.Println(i)
		fmt.Println(val)
	}
	fmt.Println(my_map)

}
