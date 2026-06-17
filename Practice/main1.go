package main  // This is required for an executable program

import "fmt"   // Import the fmt package for printing

func main() {
    // Declare variables
    name := "Aakash"  // Go can infer the type automatically
    age := 20

    // Print using fmt
    fmt.Println("Hello, Go!")
    fmt.Println("My name is", name)
    fmt.Printf("I am %d years old\n", age)
}
