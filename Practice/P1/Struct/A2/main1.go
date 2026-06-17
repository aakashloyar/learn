package main
import "fmt"
type Address struct {
    City  string
    State string
}

type Employee struct {
    Name    string
    Address Address
}
func main() {
	emp := Employee{
		Name: "Aman",
		Address: Address{City: "Delhi", State: "Delhi"},
	}
	fmt.Println(emp.Address.City) // Delhi
}