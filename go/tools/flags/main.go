package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	var val int

	args := strings.Join(os.Args[1:], "-")

	flag.IntVar(&val, "loop", 1, "number of times to loop")
	flag.Parse()

	for i := 0; i < val; i++ {
		fmt.Println(args)
	}
	flag.VisitAll(VerfiyFlag)
}

func VerfiyFlag(f *flag.Flag) {
	fmt.Printf("%s=%s\n", f.Name, f.Value)
}
