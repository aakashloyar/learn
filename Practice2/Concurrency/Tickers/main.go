//time.tick()-> standard library that return a channel that send value on a given interval
//time.After()->send value once after duration is passed//returns a channel that delivers a result after a specific interval
//time.sleep()-> blocks current go routine for  a specific duration of time
package main
import (
    "fmt"
    "time"
)
func main() {
    // for t := range time.Tick(1 * time.Second) {
    //    fmt.Println("Tick at", t)
    // }
	// fmt.Println("Waiting 2 seconds...")
	// <-time.After(2 * time.Second)
	// fmt.Println("Done!")

    ticker := time.NewTicker(1 * time.Second)
	//defer ticker.Stop()  // important: free resources

	for i := 0; i < 3; i++ {
		<-ticker.C
		fmt.Println("Tick", i)
	}
}