package main

import (
	"encodepw/pwtools"
	"fmt"
)

func main() {

	cred := pwtools.NewCredentials("rdavis", "thisIsThe-plain-password123")
	// fmt.Println(cred.Encrypt)
	// fmt.Println(cred.User)
	msg, err := cred.GetPlain("rdavis")
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(msg)
		fmt.Println(cred.Encrypt)
	}
	fmt.Println(cred.IsValid())
}
