package main
import (
	"fmt"
	"time"
)
//go doSomething()-> concurrency
func main() {
	go func() {
		fmt.Println("I have started mine homework");
		time.Sleep(3 * time.Second) // sleep for 3 seconds
		fmt.Println("finished work now")
	}()
	fmt.Println("I am playing mine work is done");

	 // Give goroutine time to finish before main exits
    time.Sleep(4 * time.Second)
}
