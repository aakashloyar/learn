package main
import "fmt"

func update(x *int) {
	*x++;
}
func main() {
	var x int=12
	fmt.Println(x)
	var y *int=&x;
	*y++;
	update(&x)
	
	fmt.Println(x,y,*y)
	
}