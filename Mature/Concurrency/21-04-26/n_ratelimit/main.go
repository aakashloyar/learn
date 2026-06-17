package main
import (
	"fmt"
	"time"
)
func main() {
	request := make(chan int, 5)
	for i:=0;i<5;i++ {
		request<- i 
	}
	close(request)
	ticker:= time.Tick(2*time.Second)
	for i:= range request {
		<-ticker
		fmt.Println("Executing ", i)
	}
}

// Tick vs NewTicker vs Ticker
// for t := range time.Tick(time.Second) {
//     	fmt.Println("tick at", t)
// 	}//runs forever
//ticker := time.NewTicker(1 * time.Second)
//fmt.Println(<-ticker.C)
//ticker.Stop()