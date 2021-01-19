package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println(os.TempDir())

	ss := time.Now()
	fmt.Println(ss)

}
