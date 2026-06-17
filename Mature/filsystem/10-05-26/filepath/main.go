package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := filepath.Join("data", "users", "file.txt")
	fmt.Println(path)
}