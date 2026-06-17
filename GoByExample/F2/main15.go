package main 
import "fmt"

func main() {
	fmt.Println("Starting division...")
	func (a, b int)  {
		if b == 0 {
			panic("division by zero")
		}
		fmt.Printf("Result: %d\n", a/b)
	}(10, 1)
	fmt.Println("Program continues...")
}