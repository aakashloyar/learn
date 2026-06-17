package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("./F43_Files/Dir/data/logs/2025", 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Directories created")
}
