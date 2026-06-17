package main
import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("Function f in f2.go")
	time.Sleep(2 * time.Second)
	fmt.Println("Exit from f")//this line will not be executed
}
func main() {
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("Exit from main")
}
//goroutine are different from threads
//they are lightweight and managed by go runtime
//they require less memory