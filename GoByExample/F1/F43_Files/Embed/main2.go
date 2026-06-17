package main

import (
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var data []byte

func main() {
	fmt.Println(string(data))
}


// We use embed so that:

// our Go program does NOT depend on external files at runtime

// Everything it needs is packed inside the single compiled binary.