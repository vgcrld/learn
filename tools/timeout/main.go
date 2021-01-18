package main

import (
	"fmt"
	"time"
)

func main() {

	// Make chan, start routine then select waits for a channel.
	// Select is not swtich. Remember that.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// the go routine sleeps for 2 and time.After waits for 1 so
	// this will timeout.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// This go routine waits for 2 seconds and time.After for 3
	// so this will not timeout.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
