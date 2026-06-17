package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./F43_Files/main2.go")//can read file only
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	fmt.Println("File opened successfully")
}
