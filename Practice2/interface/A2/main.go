package main
import "fmt"
func PrintAnything(x interface{}) {//empty interface
    fmt.Println(x)
}
func main() {
	var x interface{} = nil      // nil interface
	fmt.Println(x)
    PrintAnything(42)
    PrintAnything("hello")
    PrintAnything([]int{1, 2, 3})
}

//types
//empty
//nil
//single method
//multi method
//embedded