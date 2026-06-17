package main
import "fmt"
func main() {
	fmt.Println(check(add,2,3));


	// using an anonymous***** function
	var res int
	res = check(func(a, b int) int {
		return a * b
	}, 1, 2)
	fmt.Println(res);
}
func check(f func(int, int) int,a,b int) int {
	return f(a,b);
}
func add(a,b int) int {
	return a+b;
}