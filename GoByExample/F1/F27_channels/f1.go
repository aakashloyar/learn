package main
import (
	"fmt"
	//"time"
)

func main() {
	//declaration of channel
    var ch chan int
	//initialization of channel
	ch = make(chan int)

	ch <- 10    // send
    x := <-ch   // receive

	fmt.Println("Value received from channel:", x)
	
} 

//this code will result in a deadlock because as our channel is unbuffered and our sender and receiver are in the same goroutine(main goroutine).
