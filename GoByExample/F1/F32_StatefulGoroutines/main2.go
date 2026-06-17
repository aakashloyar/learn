package main
import (
	"fmt"

)
type CounterRequest struct {
    Op    string
    Reply chan int
}

func counterServer(reqs chan CounterRequest) {
    counter := 0 // ✅ OWNED only by this goroutine

    for req := range reqs {
        switch req.Op {
        case "inc":
            counter++
            req.Reply <- counter

        case "get":
            req.Reply <- counter
        }
    }
}

func main() {
    requests := make(chan CounterRequest)

    go counterServer(requests)

    for i := 0; i < 5; i++ {
        go func() {
            reply := make(chan int)
            requests <- CounterRequest{Op: "inc", Reply: reply}
            fmt.Println("Counter:", <-reply)
        }()
    }

    time.Sleep(time.Second)
}