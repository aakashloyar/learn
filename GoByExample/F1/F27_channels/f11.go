//✅ 2. Timeout for Receiving from a Channel (MOST COMMON)package main
import (
	"fmt"
	"time"
)


func main() {
	ch:=make(chan string,2)
	ch<-"hello"
	for range 3 {
		select {
			case msg := <-ch:
				fmt.Println("Received:", msg)

			case <-time.After(1 * time.Second):
				fmt.Println("Timeout: no message received")
		}
	}
	// select {
	// 	case msg := <-ch:
	// 		fmt.Println("Received:", msg)

	// 	case <-time.After(1 * time.Second):
	// 		fmt.Println("Timeout: no message received")
	// }

}