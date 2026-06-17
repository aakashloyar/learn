package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)

	fmt.Println("Waiting for timer...")
	<-timer.C   // blocks until timer fires

	fmt.Println("Timer finished!")
}


// Important Points

// timer.C is a channel

// <-timer.C waits until time is over

// This triggers only ONCE