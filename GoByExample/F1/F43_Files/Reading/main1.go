package main

import (
	"fmt"
	"os"
)

func main() {
	// wd, _ := os.Getwd()
    // fmt.Println("Working directory:", wd)

	//absolute path -> / direclty from root
	data, err := os.ReadFile("/home/aakash-loyar/Programming/VsCodeProjects/Go/GoByExample/F1/F43_Files/main2.go")

	//relative path -> from current working directory // ./ or directly folder name

	data, err := os.ReadFile("./F43_Files/main2.go")
	data, err := os.ReadFile("F43_Files/main2.go")
	

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
