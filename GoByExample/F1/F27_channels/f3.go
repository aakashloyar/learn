package main
import (
	"fmt"
	//"time"
)

func f(ch chan int) {
	x := <-ch   // receive

	fmt.Println("Value received from channel:", x)
}

func main() {
	//declaration of channel
    var ch chan int
	//initialization of channel
	ch = make(chan int)
	go f(ch)//no deadlock
	ch <- 10    // send
    // go f(ch)//deadlock
} 

//at this line ch<-10
//go will check if any goroutine is waiting to receive
//if there is no goroutine waiting to receive it and it is not buffered channel 
//then it will result in deadlock
//but here we have a goroutine f waiting to receive so no deadlock occurs