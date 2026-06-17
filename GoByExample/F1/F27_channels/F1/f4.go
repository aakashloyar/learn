// //now we will play around infinite for loop 

// // package main

// // import (
// // 	"fmt"
// // 	"time"
// // )

// // func main() {
// // 	for {
// // 		fmt.Println("working...")
// // 		time.Sleep(1 * time.Second) // cannot be interrupted
// // 	}
// // }


package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})

	y:=make(chan int)
	close(y)
	fmt.Println(<-y);

	// Stop the loop after 3 seconds
	go func() {
		time.Sleep(3 * time.Second)
		close(stop)
	}()

	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("working...")
		case x:=<-stop:
			fmt.Println("stopped")
			fmt.Println(x)
			return
		}
	}
}


//time.after can be stopped also