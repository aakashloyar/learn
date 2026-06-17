//in this module we will learn time.Sleep vs time.After 
//time.Sleep -> blocks the goroutine forever
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	time.Sleep(2 * time.Second) // blocks main goroutine
	fmt.Println("End")
}


// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("Waiting...")
// 	x:=time.After(2*time.Second)//x not got receive only channel type//it can only receive value cannot add
// 	fmt.Println(<-x)
// 	//fmt.Println(<-x)//deadlock
// 	//fmt.Println(<-time.After(2 * time.Second)) // receive from channel
// 	fmt.Println("Done")
// }  //After(duration ms) <-time.Time
//functionality same as time,Sleep but internally uses timer.channel


//so moral of the story is time.Sleep -> it will not return anything
//while time.After return a recieve only channel

