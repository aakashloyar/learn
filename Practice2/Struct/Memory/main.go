package main
import (
	"fmt"
	"unsafe"
)
//to make faster go left some padding due to which some memory lost
//to optimise use same byte type near
type stats struct {
	a uint16
	b uint8
	c uint8
}
type statsup struct {
	a uint8
	b uint8
	c uint16
}
func main() {
	s := statsup{a: 100, b: 10, c: 20}
	fmt.Println("Instance:", s)
	fmt.Println("Size of stats struct:", unsafe.Sizeof(s))
	// Field offsets
	fmt.Println("Offset a:", unsafe.Offsetof(s.a))
	fmt.Println("Offset b:", unsafe.Offsetof(s.b))
	fmt.Println("Offset c:", unsafe.Offsetof(s.c))
}