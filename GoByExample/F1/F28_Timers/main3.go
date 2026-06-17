package main 
import (
	"fmt"
	"time"
)


func tryRequest() {
	time.Sleep(3 * time.Second)
	fmt.Println("Request succeeded")
	//return true
}
func main() {
	fmt.Println("Starting main...")
	//success := 
	go tryRequest()
	time.Sleep(2 * time.Second)
	fmt.Println("Main finished");
}