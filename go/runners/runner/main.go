package main

import (
	"log"
	"time"

	runner "github.com/vgcrld/learngo/structs/caf"
)

type helloWorlder struct{}

func (h *helloWorlder) Run() {
	log.Println("Hello, world!")
}

func main() {
	r := runner.NewRunner()
	h := &helloWorlder{}
	r.Schedule(h, 5*time.Second)
	r.Run()
	time.Sleep(20 * time.Second)
}
