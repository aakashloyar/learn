package main
import (
	"fmt"
	"errors"
)
// type error interface {
//     Error() string
// }


func isEqual(a, b int) (bool, error) {
	if a<0 ||b<0 {
		return false, fmt.Errorf("negative numbers are not allowed")
	}
	return a==b, nil
}
func main() {
	eq, err := isEqual(5, -3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Are equal:", eq)
	}
	var ErrNotFound = errors.New("not found")
	fmt.Println(ErrNotFound)
}