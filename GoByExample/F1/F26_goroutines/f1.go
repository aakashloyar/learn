//goroutines are lightweight threads managed by Go runtime
package main 
import (
	"fmt"
	"time"
)
func main() {
	go fmt.Println("Hello from goroutine")
	time.Sleep(time.Second)//wait for goroutine to finish
	//othewise main may exit before goroutine runs
	fmt.Println("Hello from main")
}