package main 
import "fmt"

func main() {
	var a int
	a=10
	const b int=20//have to initialize a constant
	// b=20
	var c=&a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(add(&a,b));
}

func add(a *int, b int) int {
	return *a+b;
}