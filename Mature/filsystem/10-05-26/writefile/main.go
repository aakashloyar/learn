package main

import (
	"os"
)

func main() {
	// owner: read+write (6)
    // group: read      (4)
    // others: read     (4)
	file, err := os.OpenFile(
		"../README.md",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	//this will write a new line in the file 
	file.WriteString("\nNew log entry\n")
}