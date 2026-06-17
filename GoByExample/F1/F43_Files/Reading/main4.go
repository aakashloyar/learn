package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"./F43_Files/main2.go",
		os.O_RDWR|os.O_CREATE,
		0644,
	)//read, write, create
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	fmt.Println("File opened successfully")
}
