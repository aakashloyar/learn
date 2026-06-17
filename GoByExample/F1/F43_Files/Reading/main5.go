//reading a file line by line
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./F43_Files/main2.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
