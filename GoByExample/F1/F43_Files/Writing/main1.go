package main

import (
	"os"
)

func main() {
	data := []byte("Hello Aakash!\nWelcome to Go file writing.")

	err := os.WriteFile("./F43_Files/Writing/main.txt", data, 0644)
	if err != nil {
		panic(err)
	}
}



// []byte(...) → Go writes bytes, not strings

// 0644 → file permissions

// If file doesn’t exist → it is created

// If file exists → it is overwritten