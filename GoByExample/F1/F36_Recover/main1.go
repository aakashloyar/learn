package main

import "fmt"

//recover is used to regain control of a panicking goroutine.

func safeFunc() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()//play by commenting this function

    fmt.Println("About to panic...")
    panic("Something went wrong!")
}

func main() {
    safeFunc()
    fmt.Println("Program continues normally.")
}
