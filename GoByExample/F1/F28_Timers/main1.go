package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	time.Sleep(2 * time.Second)// Pause execution for 2 seconds
	fmt.Println("After 2 seconds")
}


// Use this when:

// You just need a simple pause

// No need to cancel it

// ❌ But you cannot stop or reset Sleep.