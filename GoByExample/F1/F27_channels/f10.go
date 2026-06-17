//✅ 1. Basic Delay (Like time.Sleep, but Channel-Based)
package main
import (
	"fmt"
	"time"
)


func main() {
	fmt.Println("Waiting...")

	<-time.After(2 * time.Second)

	fmt.Println("2 seconds passed")

}