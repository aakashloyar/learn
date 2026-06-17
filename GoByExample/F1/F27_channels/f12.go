//✅ 5. API Call with Timeout Simulation
package main
import (
	"fmt"
	"time"
)


func main() {
	done:=make(chan bool,2)
	//done<-true
	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Task finished")

	case <-time.After(2 * time.Second):
		fmt.Println("Task took too long, exiting")
	}//so if time taken is more than 2 seconds then timeout case will be executed
}