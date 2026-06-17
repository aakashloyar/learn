package main

import "fmt"

func main() {
    chInts := make(chan int, 1)
    chStrings := make(chan string, 1)

    //chInts <- 42
    //chStrings <- "hello"
	//select instead of switch
    select {
		case i, ok := <-chInts:
			if ok {
				fmt.Println("Received int:", i)
			}
		case s, ok := <-chStrings:
			if ok {
				fmt.Println("Received string:", s)
			}
		default:
			fmt.Println("no values ready")	
    } 
}
