package main

import "fmt"

func main() {
	ops := make(chan string)

	// Stateful goroutine
	go func() {
		counter := 0

		for {
			op := <-ops

			if op == "inc" {
				counter++
			}

			if op == "print" {
				fmt.Println("Counter:", counter)
			}
		}
	}()

	// Send requests
	ops <- "inc"
	ops <- "inc"
	ops <- "print"
}

//multiple goroutines attacking on 1 variable
//use mutex or atomic
//but you can use stateful go rountine
//idea is simple -> 1 goroutines own the resource 
//rest talk to it


//	case <-done:
//this will get executed when someone sends value or close done