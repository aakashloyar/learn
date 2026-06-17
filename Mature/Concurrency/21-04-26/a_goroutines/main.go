// go routines
package main

import (
	"fmt"
	"time"
)

func B() {
	fmt.Println("Inside B")
}
func main() {
	go fmt.Println("Hello babes")
	go B()
	go func(name string) {
		fmt.Println("Inside a function with",name)
	}("aakash")
	time.Sleep(time.Second)
}