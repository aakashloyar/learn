//reading a file word by word
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./F43_Files/main2.go")
	if(err!=nil){
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	dat:= []byte("hello world\n")
	fmt.Print(string(dat))
}
