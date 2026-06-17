//Channels
//channels helps in communication between goroutines
//these are very helpful in taking and passing values
package main
import "fmt"
func main() {
	//var ch chan int;
	ch:=make(chan int)
	//send a value in to a channel
	// send in a goroutine
	go func() {
		ch <- 42
	}()//cannot do it diectly go in deadlock
	x:= <-ch;
	fmt.Println(x)
}


// ch := make(chan int)
// ch <- 42        // (1) Send blocks here
// x := <-ch       // (2) Never reached
// fmt.Println(x)
// You created an unbuffered channel (make(chan int)).

// When you do ch <- 42, the main goroutine blocks until another goroutine is ready to receive the value.

// But in your program, there is no receiver goroutine yet. The next line (x := <-ch) can’t be reached because the program is stuck at the send.

// Result: all goroutines are asleep → deadlock.