package main

import (
    "fmt"
    "math/rand/v2"
)

func main() {

    fmt.Print(rand.IntN(100), ",")
    fmt.Print(rand.IntN(100))
    fmt.Println()

    fmt.Println(rand.Float64())

    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

    s2 := rand.NewPCG(42, 1024)//will give same sequence every time
    r2 := rand.New(s2)
    fmt.Print(r2.IntN(100), ",")
    fmt.Print(r2.IntN(100))
    fmt.Println()

	

    s3 := rand.NewPCG(42, 1024)
    r3 := rand.New(s3)
    fmt.Print(r3.IntN(100), ",")
    fmt.Print(r3.IntN(100))
    fmt.Println()
}



// s2 := rand.NewPCG(42, 1024)
// What this does
// Creates a random number generator source

// Uses the PCG algorithm (Permuted Congruential Generator)
// Think of it as:

// “Create a machine that can produce random numbers, initialized with some starting state.”

// 42 → seed (initial state)

// 1024 → stream / sequence selector

// 📌 Same seeds ⇒ same random sequence every time

// This is called deterministic randomness.


// r2 := rand.New(s2)
// What this does
// Wraps the source (s2) into a full Rand object

// Rand provides methods like:

// IntN

// Float64

// Shuffle

// Mental model:
// scss
// Copy code
// PCG source (engine)
//         ↓
// Rand (user-friendly API)