//cli flags
package main

import (
	"flag"
	"fmt"
)

func main() {
	// define a flag
	port := flag.Int("port", 8080, "server port")//port is just address

	// parse command-line flags
	flag.Parse()

	// use the flag
	fmt.Println("Server running on port:", *port)
}
//go run F45_CLI/main3.go --port=9000  //so here port will be 9000
//this is how to run with flags