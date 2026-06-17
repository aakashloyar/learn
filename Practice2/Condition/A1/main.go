package main
import "fmt"
func main() {
	length := 5
	if length < 1 {
		fmt.Println("Email is invalid")
	}
	if length := 5; length < 1 {//in this case length is available in parent scope
		fmt.Println("Email is invalid")
	}
}


// if INITIAL_STATEMENT; CONDITION {
// }