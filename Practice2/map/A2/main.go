package main

import "fmt"

func main() {
    // Define a nested map: map[string]map[string]int
    students := make(map[string]map[string]int)

    // Initialize inner map before using it
    students["Alice"] = make(map[string]int)
    students["Alice"]["math"] = 90
    students["Alice"]["physics"] = 85

    students["Bob"] = make(map[string]int)
    students["Bob"]["math"] = 75
    students["Bob"]["physics"] = 95

    // Access values
    fmt.Println("Alice math:", students["Alice"]["math"])       // 90
    fmt.Println("Bob physics:", students["Bob"]["physics"])     // 95

    // Loop through nested map
    for name, subjects := range students {
        fmt.Println("Student:", name)
        for subject, score := range subjects {
            fmt.Printf("  %s: %d\n", subject, score)
        }
    }
}
