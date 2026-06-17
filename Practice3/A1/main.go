package main

import (
	"fmt"
	"time"
)

func main() {
	mychan:=make(chan int)
	//mychan<-1
	//res:=<-mychan
	//fmt.Println(res);
	go func() {
		fmt.Println("func started")
		mychan<-1
		fmt.Println("func also gone")
	}()//now here the internal channel will go in deadlock
	//but the true deadlock is when every channel in deadlock(mainly main program)
	time.Sleep(5*time.Second)
	fmt.Println("gone")
}