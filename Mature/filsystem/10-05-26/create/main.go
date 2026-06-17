package main

import (
	"os"
)

func main() {
	file, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}