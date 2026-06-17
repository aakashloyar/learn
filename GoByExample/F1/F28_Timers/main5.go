package main

import (
	"fmt"
	"time"
)

func fakeAPI() string {
	time.Sleep(1 * time.Second) // Simulate slow API
	return "API Response"
}

func main() {
	result := make(chan string)

	go func() {
		result <- fakeAPI()
	}()

	timer := time.NewTimer(2 * time.Second)

	select {
	case res := <-result:
		fmt.Println("Got:", res)

	case <-timer.C:
		fmt.Println("Timeout hit!")
	}
}

//this is correct
//in the last code the executuion will still happen 
//so there can be money loss