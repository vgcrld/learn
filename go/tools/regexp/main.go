package main

import (
	"os"
	"regexp"
)

func main() {

	// The os arguments passed
	args := os.Args

	// The regex to check
	pattern := args[1]

	// The value to match
	value := args[2]

	// Exit 0 if it matches, 1 if not
	if match(pattern, value) {
		os.Exit(0)
	} else {
		os.Exit(1)
	}

}

func match(p, s string) bool {
	m, e := regexp.MatchString(p, s)
	if e != nil {
		return false
	} else {
		return m
	}
}
