package main
import "fmt"
func main() {
	x := 42
	var y *int=&x;
	fmt.Println(x);
	fmt.Println(y);
	fmt.Printf("value of p: %v\n", y)
	*y=10;
	fmt.Printf("value of p: %d\n", *y)
	//fmt.Printf(*y);//cannot
}