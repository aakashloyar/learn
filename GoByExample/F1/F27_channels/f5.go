//buffered channels

package main
import (
	"fmt"
)
func f(ch chan int) {
	x := <-ch   // receive
	fmt.Println("Value received from channel:", x)	
	ch<-30
}
func main() {
	var ch chan int
	ch=make(chan int,2)//buffered channel of capacity 2
	ch<-10
	ch<-20
	go f(ch);
	//ch<-30//this will result in deadlock as buffer is full
	fmt.Println(<-ch);
	fmt.Println(<-ch);
	//10 20 30
}
