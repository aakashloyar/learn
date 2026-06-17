package main
// var s []int       // s is nil
// var c chan string // c is nil

// var s = make([]int, 5) // s is initialized and not nil
// var c = make(chan int) // c is initialized and not nil

// A Send to a Nil Channel Blocks Forever
// var c chan string        // c is nil
// c <- "let's get started" // blocks


// A Receive from a Nil Channel Blocks Forever
// var c chan string // c is nil
// fmt.Println(<-c)  // blocks


// A Receive from a Closed Channel Returns the Zero Value Immediately
// var c = make(chan int, 100)
// close(c)
// fmt.Println(<-c) // 0


// A Send to a Closed Channel Panics
// var c = make(chan int, 100)
// close(c)
// c <- 1 // panic: send on closed channel


import "fmt"
func main() {
	// var s []int       // s is []
	// var c chan string // c is nil
	// fmt.Println(s,c);

	// var s = make([]int, 5) // s is initialized and not nil
	// var c = make(chan int) // c is initialized and not nil
    //fmt.Println(s,c);
	// var c chan string // c is nil
    // fmt.Println(<-c)  // blocks//deadlock

	// var c = make(chan int, 100)
	// close(c)
	// c <- 1 // panic: send on closed channel//error
	// fmt.Println("done")
	
	var c = make(chan int, 100)
	close(c)
	fmt.Println(<-c) // 0
}