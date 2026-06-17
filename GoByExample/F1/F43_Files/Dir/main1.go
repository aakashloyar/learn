package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("./F43_Files/Dir/mydir", 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Directory created")
}
