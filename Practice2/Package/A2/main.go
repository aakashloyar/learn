package main
//package random
//import "github.com/mailio/rand"//just take the last keyword when importing//mostly package name is that only
//but if it is different then you have to use that// but for convention you  have to give the same name
import "math/rand"
import "fmt"
func main() {
	rand.Seed(42)
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("", rand.Intn(100))
}