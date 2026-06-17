package main
import "fmt"

func main() {
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 24
	ages["Mary"] = 21 // overwrites 24

	fmt.Println(ages)
	ages = map[string]int{
	"John": 37,
	"Mary": 21,
	}
	fmt.Println(ages)

	elem, ok := ages["John"]
	fmt.Println(elem,ok)
	delete(ages, "John")
}

