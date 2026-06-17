package main 
import (
	"fmt"
)
func main() {
	ch := make(chan string,2)
	ch<-"hello"
	close(ch)
	//ch<-"world" //panic: send on closed channel
	fmt.Println(<-ch);//receiver can still receive values from closed channel
	//close ch//panic: close of closed channel//now no one can send or receive values from closed channel
}

//afer a channel is closed after empty it returns default value
//for int -> 0 for struct -> {}