package main
import "fmt"
func main() {
	fmt.Println("start")
	defer fmt.Println("this was deferred")//always executed at the end of the function
	fmt.Println("about to panic");
	panic("something bad happened");
	fmt.Println("end");
}