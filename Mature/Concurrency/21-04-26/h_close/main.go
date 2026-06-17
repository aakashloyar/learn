package main

import (
	"time"
	"fmt"
)

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
		time.Sleep(time.Second)
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)//more you are getting from here if closed then more is false otherwise true
    fmt.Println("sent all jobs")

    <-done

    _, ok := <-jobs
    fmt.Println("received more jobs:", ok)
}