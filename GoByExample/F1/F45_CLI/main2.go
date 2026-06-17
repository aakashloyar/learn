package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	fmt.Println("Number:", num)

	for i, arg := range os.Args[1:] {
		fmt.Printf("Arg %d: %s\n", i, arg)
	}
}
