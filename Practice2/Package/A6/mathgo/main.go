// ~/go-dev/mymath/mymath.go
package mathgo

// Sum returns the sum of two ints.
// Note: exported names must start with a capital letter.
func Sum(a, b int) int {
    return a + b
}


// go mod init github.com/aakash/mathgo->after this command go.mod file formed with content

//module github.com/aakash/mathgo
//go 1.25.0


//go mod tidy
// 🔹 When to run it

// After adding new imports → tidy ensures go.mod gets updated.

// After removing imports → tidy cleans up unused dependencies.

// Before committing code → keeps go.mod and go.sum minimal and reproducible.
