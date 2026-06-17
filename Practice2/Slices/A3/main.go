package main
import "fmt"
//variadic
func hello(nums ...int) {
	for i:=0;i<len(nums);i++ {
		fmt.Println(nums[i]);
	}
}

func print(strings ...string) {
	for i:=0;i<len(strings);i++ {
		fmt.Println(strings[i]);
	}
}
func main() {
	hello(0,1,0,2)	
	names := []string{"bob", "sue", "alice"}
    print(names...)//spread operator
}

