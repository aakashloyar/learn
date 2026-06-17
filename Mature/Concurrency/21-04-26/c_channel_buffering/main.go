package main 
import "fmt"

func main() {
	msg:= make(chan string,2)
	defer close(msg)
	msg<-"A"
	msg<-"B"
	fmt.Println(<-msg)
	x:= <-msg
	fmt.Println(x)
}