package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo started at:", time.Now().Format("15:04:05"))

	// -------------------------------
	// 1. time.After (fires once)
	// -------------------------------
	go func() {
		fmt.Println("Waiting 2s for time.After...")
		<-time.After(2 * time.Second)
		fmt.Println("time.After fired at", time.Now().Format("15:04:05"))
	}()

	// -------------------------------
	// 2. time.Tick (ticks forever, cannot stop)
	// -------------------------------
	go func() {
		for t := range time.Tick(1 * time.Second) {
			fmt.Println("time.Tick ticked at", t.Format("15:04:05"))
		}
	}()

	// -------------------------------
	// 3. time.NewTicker (stoppable ticker)
	// -------------------------------
	go func() {
		ticker := time.NewTicker(1500 * time.Millisecond)
		defer ticker.Stop() // important: prevent leaks

		for i := 0; i < 3; i++ {
			<-ticker.C
			fmt.Println("NewTicker tick at", time.Now().Format("15:04:05"))
		}
		fmt.Println("NewTicker stopped")
	}()

	// -------------------------------
	// 4. time.NewTimer (stoppable timer)
	// -------------------------------
	go func() {
		timer := time.NewTimer(5 * time.Second)
		go func() {
			<-timer.C
			fmt.Println("NewTimer fired at", time.Now().Format("15:04:05"))
		}()

		// Stop the timer before it fires
		time.Sleep(3 * time.Second)
		if timer.Stop() {
			fmt.Println("NewTimer stopped before firing")
		}
	}()

	// Keep program running long enough to see output
	time.Sleep(20 * time.Second)
	fmt.Println("Demo finished")
}
