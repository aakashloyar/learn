package main
import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("Function f in f2.go")
	time.Sleep(2 * time.Second)
	fmt.Println("Exit from f")
}
func main() {
	time.Sleep(3 * time.Second)
	go f()
	fmt.Println("Exit from main")
}
