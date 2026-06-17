package main

import (
    "fmt"
    "time"
)

func main() {

    // timer1 := time.NewTimer(2 * time.Second)

    // <-timer1.C
    // fmt.Println("Timer 1 fired")

    // timer2 := time.NewTimer(time.Second)
    // go func() {
    //     <-timer2.C
    //     fmt.Println("Timer 2 fired")
    // }()
    // stop2 := timer2.Stop()
    // if stop2 {
    //     fmt.Println("Timer 2 stopped")
    // }

    // time.Sleep(2 * time.Second)

	//can this stop be false
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()


	time.Sleep(2* time.Second)

	stop2 := timer2.Stop()
	fmt.Println(stop2)//false
	//here timer is already executed so it has not time to fire stop 
}
//why timers not sleep bcz timers are interruptable

