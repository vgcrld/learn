// Calculate maximal value in a slice
package main

import (
	"fmt"
)

// this for loop with range - loop over the slice
func main() {
	nums := []int{16, 8, 42, 4, 23, 15}

	max := nums[0] // Initialize max with first value
	// [1:] skips the first value
  // Here we loop over nums with range
	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}

	fmt.Println(max)
}
