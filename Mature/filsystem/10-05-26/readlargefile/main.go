package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../README.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
//do not use readFile for large file 
//use bufio it reads line by line so low memory