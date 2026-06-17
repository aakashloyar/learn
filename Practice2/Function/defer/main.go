package main
import "fmt"
func main() {
	check(2);
}
func check(b int) int{
	if b==3 {
		return 2;
	}
	a:=2
	defer fmt.Println("Hi done")//after this statment whenever function return this statement will execute
	//not if function return initially//execute at the last
	fmt.Println(a)
	return 0;
}
