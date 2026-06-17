package main
import "fmt"

func main() {
	fmt.Println("Hi boy");
	fmt.Print("Hello")
    fmt.Println(" World") // Output: Hello World

	fmt.Println("Hello", "World") // Output: Hello World\n
    name := "Aakash"
	age := 21
	fmt.Printf("My name is %s and I am %d years old\n", name, age)

}