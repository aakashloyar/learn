package main
import (
    "fmt"
    "strconv"
    "errors"
)
type error interface {
    Error() string
}
func main() {
    
    i, err := strconv.Atoi("42")
    fmt.Println(i,err)
    var er error = errors.New("something went wrong")
    fmt.Println(er)
}
