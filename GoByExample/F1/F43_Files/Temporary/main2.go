package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "mydir-*")
	if err != nil {
		panic(err)
	}

	// Always clean up directory + contents
	defer os.RemoveAll(tmpDir)

	fmt.Println("Temporary directory:", tmpDir)

	// Create a file inside the temp directory
	filePath := filepath.Join(tmpDir, "data.txt")

	err = os.WriteFile(filePath, []byte("Hello from temp directory!"), 0644)
	if err != nil {
		panic(err)
	}

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println("File inside temp dir:", filePath)
	fmt.Println("File content:", string(data))
}


//default files and directories are created in system temp location
//in linux /tmp