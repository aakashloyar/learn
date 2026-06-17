package main
import "fmt"


type Number interface {
	~int | ~float64 | ~float32 |string
}

func Min[T Number](a, b T) T {//comparable means the types that we are passing must be comparable
	if a < b {
		return a
	}
	return b
}
func Print [T any] (a T) {
	fmt.Println(a)
}

func main() {
	fmt.Println(Min(3, 7))        // int
	fmt.Println(Min(2.5, 1.8))    // float64
	fmt.Println(Min("go", "java")) // string
	Print("hello");
}
