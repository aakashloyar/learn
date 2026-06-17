package main
import "fmt"
type Person struct {
	Name string
	Age int
	City string
}
func main() {
	// p1 := Person{Name: "Aakash", Age: 21, City: "Hyderabad"}

    // // Access fields
    // fmt.Println(p1.Name) // Aakash
    // fmt.Println(p1.Age)  // 21

    // // Modify fields
    // p1.City = "Delhi"
    // fmt.Println(p1.City) // Delhi
	// 1. Using field names (recommended)

	// 3. Zero values
	var p3 Person
	fmt.Println(p3) // {"" 0 ""}

	//p1 := Person{Name: "Raj", Age: 30}

	// 2. Without field names (order matters)
	//p2 := Person{"Anita", 25, "Mumbai"}

	

}