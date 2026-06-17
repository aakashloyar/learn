package main

import (
	"os"
	"fmt"
)

func main() {
	err := os.Remove("output.txt")
	if err != nil {
		fmt.Println("error while removing file",err)
	}
	os.RemoveAll("temp")//delete file recursively

	os.Mkdir("data", 0755) // make directory

	os.MkdirAll("a/b/c", 0755) // make all directory
}