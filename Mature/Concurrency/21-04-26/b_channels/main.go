//channels
package main 
import (
	"fmt"
	"time"
)
func main() {
	msg:= make(chan string)
	go func() {
		fmt.Println("Hi")
	}()
	go func(input string) {
		fmt.Println("Inside putting")
		msg<- input 
		fmt.Println("Outside putting")
	}("aakash")

	go func() {
		fmt.Println("Inside removing")
		fmt.Println(<-msg)
		fmt.Println("Outside removing")
	}()
	time.Sleep(time.Second)
	fmt.Println("Exit")
}