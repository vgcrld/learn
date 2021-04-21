package main

import(
  "fmt"
  "os"
  "strings"
)

func main() {

  args := os.Args
  sep := args[1]
  str := args[2]

  split := strings.Split(str,sep)
  res   := strings.Join(split, " ")

  fmt.Printf("%s\n",res)



}
