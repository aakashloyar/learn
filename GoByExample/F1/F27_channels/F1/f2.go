//now in this we will try to see select with time.Sleep vs time.After
package main 
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Started")
	// select {
	// case <-time.Sleep(1*time.Second)://cannot do this 
	// }
	select {
	case x:=<-time.After(2*time.Second):
		fmt.Println("* ",x)
	case y:=<-time.After(1*time.Second):
		fmt.Println("** ",y)
	}

}

//so time.After works with select while time.Sleep donot