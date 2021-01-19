package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	Name   string
	Age    int
	Status string
}

func (d *Data) ShowTime() {
	time := time.Now()
	fmt.Printf("%v", time)
}

var wg sync.WaitGroup

var c1 chan Data = make(chan Data)

func main() {

	wg.Add(1)
	go worker(1, time.Second*1)

	for _, x := range []Data{
		{"rich", 50, "run"},
		{"rich", 50, "run"},
		{"rich", 50, "run"},
		{"rich", 50, "run"},
		{"rich", 50, "run"},
		{"rich", 50, "run"},
		{"rich", 50, "done"},
	} {
		c1 <- x
	}

	wg.Wait()

}

func worker(i int, s time.Duration) {
	var data Data
	for {
		data = <-c1
		if data.Status == "done" {
			break
		}
		fmt.Printf("Routine %v: %v - %v\n", i, data.Name, data.Age)
		time.Sleep(s)
	}
	data.ShowTime()
	wg.Done()
}
