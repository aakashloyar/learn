package main 
import "fmt"

func a() func() int {
	i:=0
	return func() int {
		i++;
		return i
	}
}
func main() {
	f := a();
	fmt.Println(f());
	fmt.Println(f());

	g:= a();
	fmt.Println(g())
}