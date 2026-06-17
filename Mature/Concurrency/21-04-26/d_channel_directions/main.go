package main

import (
	"fmt"
	"time"
)

func input(done chan<- string, name string) {
	done<- name 
}
func output(done <-chan string) string{
	return <-done 
}
func main() {
	done:= make(chan string)
	defer fmt.Println(done)
	go input(done,"aakash")
	var name string 
	go func() {
		name = output(done)
	}()
	time.Sleep(time.Second)
	fmt.Println(name)

}
