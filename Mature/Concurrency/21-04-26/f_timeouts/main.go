package main

import (
	"fmt"
	"time"
)

func main() {
	c1:= make(chan string)
	c2:= make(chan string)
	go func() {
		fmt.Println("Inside func1")
		time.Sleep(3*time.Second)
		c1<- "a"
	}()
	go func() {
		fmt.Println("Inside function2")
		time.Sleep(4*time.Second)
		c2<- "b"
	}()
	for range 3 {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)	
		case <-time.After(1*time.Second):
			fmt.Println("Exit")
		}
	}
	fmt.Println(<-time.After(5*time.Second))
}