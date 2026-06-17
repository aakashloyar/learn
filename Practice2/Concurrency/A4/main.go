package main
import "fmt"
func main() {
	ch:=make(chan int,2)
	ch<-22
	ch<-23

	close(ch)  // ✅ signal no more values will be sent otherwise for loop  will continue forever 
	for item := range ch {
    // item is the next value received from the channel
	    fmt.Println(item)//deadlock
	}
}