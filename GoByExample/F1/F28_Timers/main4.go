package main 
import (
	"fmt"
	"time"
)


func tryRequest() bool {
	time.Sleep(4 * time.Second)
	fmt.Println("Request succeeded")
	return true
}
func main() {
	fmt.Println("Starting main...")
	go func() {
		fmt.Println("Trying request...")
		res := tryRequest()
		fmt.Println("Got:", res)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Main finished");
}