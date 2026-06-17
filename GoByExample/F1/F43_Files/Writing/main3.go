//create and write to a file
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./F43_Files/Writing/create.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // VERY important

	n, err := file.Write([]byte("Writing using os.Create\n"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Bytes written:", n)
}
