package main

import "fmt"

func main() {
	hits := make(map[string]map[string]int)
	user := "alice"
    page := "home"

    // Check if inner map exists
    if hits[user] == nil {
        hits[user] = make(map[string]int) // initialize it
    }

    // Now safe to increment
    hits[user][page]++
    hits[user]["about"]++

    fmt.Println(hits)
	attended := map[string]struct{}{
		"Ann": {},
		"Joe": {},
	}
	fmt.Println(attended)
	messages := []string{"Hello world", "hello there", "General Kenobi"}
	fmt.Println(messages)
}