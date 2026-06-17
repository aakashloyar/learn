package main

import (
    "fmt"
    "os"
)

func main() {

    defer fmt.Println("!")

    os.Exit(3)//defer will not run because os.Exit terminates the program immediately
}