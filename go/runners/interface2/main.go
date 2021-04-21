package main

import "fmt"

type Boss struct {
	Name   string
	Salary float64
}

func (b Boss) pay() float64 {
	fmt.Println("Paying the salary boss...")
	return b.Salary / 26
}

type Worker struct {
	Name  string
	Rate  float64
	Hours float64
}

func (w Worker) pay() float64 {
	fmt.Println("Paying the hourly employee...")
	return w.Rate * w.Hours
}

// Employee interface call defines anything that can 'pay() float64'
type Employee interface {
	pay() float64
}

func writeChecks(e Employee) float64 {
	return e.pay()
}

func main() {

	b := new(Boss)
	b.Name = "Richard"
	b.Salary = 52544.33

	fmt.Printf("Boss: %s $%.2f\n", b.Name, writeChecks(b))

	e := new(Worker)
	e.Name = "Justin"
	e.Hours = 22.0
	e.Rate = 14.55

	fmt.Printf("Worker: %s $%.2f\n", e.Name, writeChecks(e))

}
