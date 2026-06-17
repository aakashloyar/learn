package main
import (
	"fmt"
	"errors"
)


// func f(arg int) (int, error) {
//     if arg == 42 {
//         return -1, errors.New("can't work with 42")
//     }
//     return arg + 3, nil
// }


type argError struct {
    arg     int
    message string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}


func f(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}
func main() {
	fmt.Println(f(42))
	var ErrOutOfTea = errors.New("no more tea available")
	fmt.Println(ErrOutOfTea)

	for _, i := range []int{7, 42} {
        if r, e := f(i); e != nil {
            fmt.Println("f failed:", e)
        } else {
            fmt.Println("f worked:", r)
        }
    }
}