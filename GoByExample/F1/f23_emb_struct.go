package main 
import "fmt"

type Address struct {
	Street string
	City string
}
type Person struct {
	Name    string
	Age     int
	Address Address	
}
func main() {
	p := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
		},
	}
	fmt.Println(p)
	fmt.Println(p.Address.Street)
}