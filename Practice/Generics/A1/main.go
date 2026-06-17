package main
import "fmt"

type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64 |
    ~string
}

func Max[T Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
func main() {
	fmt.Println(Max(3, 5))       // int
    fmt.Println(Max(2.3, 4.1))   // float64
    fmt.Println(Max("go", "zo")) // string
}