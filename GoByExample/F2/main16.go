package main

import "fmt"

func div(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}
	fmt.Printf("Result: %d\n", a/b)
}

func main() {
	div(10, 0)
	fmt.Println("Program continues...")
}
//recover stops a panic and return the value passed to panic.