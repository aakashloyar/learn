package main
import "fmt"
func send(ch chan int) {
    ch <- 99
}

func main() {
    ch := make(chan int, 1)
    send(ch)
    //fmt.Println(<-ch) // 99
	channel:=make(chan int,1)
	channel<-22
	//channel<-23
	x:=<-channel
	fmt.Println(x)
	v, ok := <-ch
	fmt.Println(v,ok)//cheking it present
	close(ch)//closing channel
	// v, ok := <-ch
	// fmt.Prinlnt(v,ok);
}


// Unbuffered channel → make(chan int)

// Has capacity 0.

// Send (ch <- x) blocks until a receiver is ready.

// Receive (<-ch) blocks until a sender is ready.

// Sender and receiver must synchronize at the same time.

// Buffered channel → make(chan int, N)

// Has capacity N.

// Can hold up to N values without requiring an immediate receiver.

// Send blocks only when the buffer is full.

// Receive blocks only when the buffer is empty.


//channel:=make(chan int,1)
//channel<-22
//x:=channel //works fine

// channel:=make(chan int)//require immediate receiver 
// 	channel<-22//deadlock
//	x:=<-channel

// channel:=make(chan int,1)//1 only it can handle
// 	channel<-22
// 	channel<-23//problem here deadlock
// 	x:=<-channel



//    fmt.Println(<-ch)