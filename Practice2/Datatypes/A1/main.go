package main
import "fmt"
func main() {
	var r rune = 'A'
	var x int32 = 'A'
	fmt.Println(r == x) // true
	fmt.Println(r) // 65
	fmt.Printf("%c\n", r) // prints the character → A
}
