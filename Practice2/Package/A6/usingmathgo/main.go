package main

import (
    "fmt"
    "github.com/aakash/mathgo"
)

func main() {
    fmt.Println("2 + 3 =", mathgo.Sum(2, 3))
}


//change go.mod file update and run it
// aakash-loyar@aakash-loyar-VivoBook-ASUSLaptop-X421EAYB-X413EA:~/Programming/VsCodeProjects/Go/Practice2/Package/A6/usingmathgo$ go mod edit -replace=github.com/aakash/mathgo=../mathgo
// go mod tidy
//mow you will be able to run it locally
//go run .

