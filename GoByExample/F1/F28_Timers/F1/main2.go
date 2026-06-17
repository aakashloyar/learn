package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(1 * time.Second)
    done := time.After(5 * time.Second)

    for {
        select {
        case t := <-ticker.C:
            fmt.Println("Tick at:", t)

        case <-done:
            fmt.Println("Done!")
            ticker.Stop()
            return
        }
    }
}

//stopping a ticker