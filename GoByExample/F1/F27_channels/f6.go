//buffered channels

package main
import (
	"fmt"
	"time"
)
func func1() {
	time.Sleep(3*time.Second)
}
func f(ch chan int) {
	time.Sleep(4 * time.Second)
	x := <-ch   // receive
	time.Sleep(2*time.Second)
	fmt.Println("Value received from channel:", x)	
	//ch<-30
}
func main() {
	var ch chan int
	ch=make(chan int,2)//buffered channel of capacity 2
	ch<-10
	ch<-20
	go f(ch);
	//func1();

	//****************************
	ch<-30//this will act like a pivot as it will wait for buffer to have space 
	//but when each thread is done executing main thread will exit
	//****************************
	time.Sleep(1 * time.Second)
	//10 20 30
	fmt.Println("Exited from main")
}
