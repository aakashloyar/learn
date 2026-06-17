package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("README.md")
	if err != nil {
		panic(err)
	}

	fmt.Println("Name:", info.Name())
	fmt.Println("Size:", info.Size())
	fmt.Println("IsDir:", info.IsDir())
	fmt.Println("Modified:", info.ModTime())
	fmt.Println("Permissions:", info.Mode())
}