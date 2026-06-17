package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	start := flag.NewFlagSet("start", flag.ExitOnError)
	port := start.Int("port", 8080, "port")

	if len(os.Args) < 2 {
		fmt.Println("use: start")
		return
	}

	if os.Args[1] == "start" {
		start.Parse(os.Args[2:])
		fmt.Println("start on", *port)
	}
}

//subcommand
// git push origin main
// here push is subcommand
// origin and main are args to subcommand