
// Simple setup

package main


import (
  "fmt"
)

func closeStuff(url string) (string,error) {
  fmt.Println("DEFER:: In closeStuff - and closing.")
  return "https://ibm.com", nil
}

func main() {

  defer closeStuff("https://ibm.com")

  msg := `
    This is some code.
    How does it look?
    When I print it.
  `

  fmt.Println("This is my work that is happening.")
  fmt.Println("we are done here")
  fmt.Println(msg)




}
