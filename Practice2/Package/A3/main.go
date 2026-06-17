//packge is collection of .go file in a single directory with same package name
//module is collection of packages rooted at a diectory that has go.mod file
// module (project)
//  ├── go.mod        (defines the module path & dependencies)
//  ├── main.go       (package main)
//  └── utils/
//      └── math.go   (package utils)

// Here:

// main.go belongs to package main.

// utils/math.go belongs to package utils.

// Together they belong to one module (because they live under a directory with go.mod).

// 2. What’s inside go.mod

// When you run go mod init mymodule, Go creates:

// module mymodule

// go 1.23

// This means:

// The module’s import path is mymodule.

// Any package inside this folder can be imported using that path.

// If you later import other modules, Go adds them here, e.g.:

// require (
//     github.com/google/uuid v1.6.0
//     golang.org/x/crypto v0.25.0
// )


// 3. How imports relate to modules

// Say you have:

// go.mod → module github.com/you/awesome


// And inside:

// awesome/
//    go.mod
//    main.go       → package main
//    helper/
//       helper.go  → package helper


// In main.go, you can import your helper package like this:

// import "github.com/you/awesome/helper"


// because the module path is github.com/you/awesome and the subdirectory is helper.


// 4. Why modules matter

// They tell Go where your code lives (local or GitHub).

// They define your dependencies (and their versions).

// They make your project reproducible (anyone cloning your repo can run go mod tidy and get the exact same versions).


package main
import "math/rand"
import "fmt"
func main() {
	rand.Seed(42)
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("", rand.Intn(100))
}