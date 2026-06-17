//go routine delay

package main 

import (
	"fmt"
	"time"
)

func A() {
	//time.Sleep(1*time.Second)

	//<-time.After(1*time.Second)
	x:=time.After(1*time.Second)//this only block this way now directly
	fmt.Println(<-x);
}

func main() {
	fmt.Println("Start")
	A()
	fmt.Println("Done")
}