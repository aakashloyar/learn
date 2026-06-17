package main 
import (
	"fmt"
)
type Years int
type Person struct {
	Name  string
	Age   Years
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}
func (p *Person) Birthday() {
	p.Age += 1
}
func main() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p)
	fmt.Println(p.Greet())
	fmt.Println(p.Age)
	p.Birthday()
	fmt.Println(p.Age)
}