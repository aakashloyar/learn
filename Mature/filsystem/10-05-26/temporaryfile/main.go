package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(file.Name())
}
//created in os temp directory
//can give a random naming regex which os will handle