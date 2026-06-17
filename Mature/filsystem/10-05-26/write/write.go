package main

import (
	"os"
)

func main() {
	content := []byte("Hello from Golang!")
	//this is not adding line this is just removing the old content and add new
	err := os.WriteFile("../README.md", content, 0644)
	if err != nil {
		panic(err)
	}
}