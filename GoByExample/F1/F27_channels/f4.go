package main
import (
	"fmt"
)

func f(ch chan int) {
	x := <-ch   // receive
	fmt.Println("Value received from channel:", x)
}

func main() {
	var ch chan int
	ch=make(chan int)
	x:=<-ch;//noone is sending value to channel so deadlock occurs
	ch<-10
	fmt.Println(x)
}

