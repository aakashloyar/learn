//Closures

//A closure is a function that remembers variables outside its scope.
package main
import "fmt"
func counter() func() int {
    x := 0

    return func() int {
        x++
        return x
    }
}


func main() {
	c1:=counter();
	fmt.Println(c1());
	fmt.Println(c1());
}
