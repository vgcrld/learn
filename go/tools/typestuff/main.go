package main

import (
	"fmt"
	"regexp"

	"github.com/davecgh/go-spew/spew"
)

// User struct for a user def
type User struct {
	name  string
	age   int
	email email
}

// Type of email that is a string
type email string

// Then we can add action to a simple string
// since it is a string. Notice it needs to be cast: string(e)
func (e email) verify() bool {
	eString := string(e)
	v, _ := regexp.MatchString("com$", eString)
	return v
}

// Users A way to shortcut creating something like a map
type Users map[int]User

func main() {

	x := Users{
		1:   {"Rich", 50, "vgcrld@gmail.com"},
		100: {"Rich2", 50, "vgcrld@gmail.com"},
	}
	x[2] = User{"bob", 54, "bob@bob.c"}
	x[3] = User{"sue", 52, "sue"}

	var em email = "vgcrld@gmail.comm"
	fmt.Println(em.verify())
	spew.Dump(x)

}
