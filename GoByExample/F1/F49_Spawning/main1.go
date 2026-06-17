package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Println("Worker", id)
}

func main() {
	for i := 1; i <= 5; i++ {
		go worker(i) // spawn 5 goroutines
	}

	time.Sleep(1 * time.Second)
}
