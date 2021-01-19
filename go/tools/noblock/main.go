package main

import "fmt"

func main() {
	Test1()
}
func Test2() {

	// This is ok - buffer will get the data
	data := make(chan string, 1)
	data <- "hello"
	fmt.Println(<-data)

	// This is ok because the go routine will get the data
	// and <-data is blocking - so execution will wait until
	// the routine puts data on it.
	data3 := make(chan string, 0)
	go func() { data3 <- "from routine" }()
	fmt.Println(<-data3)

	// This is bad - no buffer or receiver to take the data
	// The compiler knows to kill this so nothing above even runs.
	data2 := make(chan string, 0)
	data2 <- "hello"
}

func Test1() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here’s a non-blocking receive. If a value is available on messages then select
	// will take the <-messages case with that value.
	// If not it will immediately take the default case.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// A non-blocking send works similarly. Here msg cannot be sent to the messages channel,
	// because the channel has no buffer and there is no receiver.
	// Therefore the default case is selected.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use multiple cases above the default clause to implement a multi-way
	// non-blocking select. Here we attempt non-blocking receives on
	// both messages and signals.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
