//Directions in channels
//in this file we will see how sending and receiving works in channels
package main 
import (
	"fmt"
)
func receiveOnly(ch<-chan int) {
	x:=<-ch
	fmt.Println("Value received from receiveOnly channel:",x)
}

func sendOnly(ch chan<- int) {
	//x:=<-ch;//this will result in error as we cant receive from send only channel
	ch<-20;
	fmt.Println("Value sent to sendOnly channel");
}
func main() {
	ch1 := make(chan int,2)
	ch2 := make(chan int,3)
	go receiveOnly(ch1)
	go sendOnly(ch2)
}