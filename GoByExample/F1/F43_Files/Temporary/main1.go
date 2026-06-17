package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a temporary file in system temp directory
	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		panic(err)
	}

	// Always clean up
	defer os.Remove(tmpFile.Name())

	// Write data
	_, err = tmpFile.WriteString("Hello from temporary file!")
	if err != nil {
		panic(err)
	}

	// Close the file
	tmpFile.Close()

	// Read the file
	data, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		panic(err)
	}

	fmt.Println("Temporary file name:", tmpFile.Name())
	fmt.Println("File content:", string(data))
}
