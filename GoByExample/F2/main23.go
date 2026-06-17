package main

import "fmt"

// Smallest common behavior
type Worker interface {
	Work()
}

// Extra capability
type Eater interface {
	Eat()
}

type Human struct{}

func (Human) Work() {
	fmt.Println("Human is working")
}

func (Human) Eat() {
	fmt.Println("Human is eating")
}

type Robot struct{}

func (Robot) Work() {
	fmt.Println("Robot is working")
}

func main() {
	var w Worker

	w = Human{}
	w.Work()

	w.Eat() //❌ compile-time error (as desired)

	var e Eater

	e = Human{}
	e.Eat()

	w = Robot{}
	w.Work()

	// e = Robot{} ❌ compile-time error
}
