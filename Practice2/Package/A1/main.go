package main //this package convert this file to executable//entry point of program
import (
	"fmt"
	"math/rand"
)//pulls external packages

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}