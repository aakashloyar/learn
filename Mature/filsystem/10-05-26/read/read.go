package main

import (
	"fmt"
	"os"
)

func main() {
	// this path is relative path 
	// relative to terminal path
	data, err := os.ReadFile("../README.md")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data))
}