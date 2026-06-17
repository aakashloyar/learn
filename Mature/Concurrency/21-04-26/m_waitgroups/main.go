package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) 

        go func(i int) {
            defer wg.Done()
            time.Sleep(time.Second)
            fmt.Println("Job done:", i)
        }(i)
    }

    wg.Wait() 
    fmt.Println("All jobs finished")
}
//wg,Go do both steps itself