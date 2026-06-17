package main
import (
	//"fmt"
	//"time"
)

func main() {
	//declaration of channel
    var ch chan int
	//initialization of channel
	ch = make(chan int)

	ch <- 10    // send
    //x := <-ch   // receive

	//fmt.Println("Value received from channel:", x)
	
} 


// Step	What Go Does
// 1	main() starts
// 2	Channel is created (unbuffered)
// 3	ch <- 10 executes
// 4	Go says: ❝Is any goroutine waiting to receive?❞
// 5	Answer: ❌ NO
// 6	So Go parks (sleeps) the main goroutine
// 7	Now zero runnable goroutines exist
// 8	Go runtime detects deadlock
// 9	❌ Program crashes with deadlock