package main
import (
	"fmt"
	"errors"
)
func main() {
	a, _ := getPoint()//ignoring return value
	fmt.Println(a)
	b, y,err := getCoords(0)
	fmt.Println(b, y,err)

	
}
func getPoint() (x int, y int) {
    return 3, 4
}


//*** named return
func getCoords(b int) (x, y int,err error) {
	// x and y are initialized with zero values
	if b == 0 {
      return 0, 0, errors.New("can't divide by zero")
    }
	return // automatically returns x and y
}

func getCoord() (x, y int) {
	return x, y // this is explicit
}
func getCords() (x, y int) {
    return 5, 6 // this is explicit, x and y are NOT returned
}

func getCrds() (x, y int) {
    return // implicitly returns x and y
}