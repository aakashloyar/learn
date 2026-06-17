package main

import (
	"fmt"
	"time"
)

func main() {
	// timer := time.NewTimer(1 * time.Second)

	// time.Sleep(3 * time.Second)

	// stopped := timer.Stop()
	// fmt.Println("Timer stopped:", stopped)

	// time.Sleep(3 * time.Second)
	// fmt.Println("Program finished")
	timer := time.NewTimer(1 * time.Second)

	time.Sleep(4 * time.Second)

	stopped := timer.Stop()
	fmt.Println(stopped)

}
