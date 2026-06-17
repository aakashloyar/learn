//range over channels
package main

import "fmt"

func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    for elem := range queue {
        fmt.Println(elem)
    }
    for elem := range queue {
        fmt.Println(elem)
    }//this will not execute

	fmt.Println(<-queue);
	fmt.Println(<-queue)
}

