package main 
import "fmt"
type Rect struct {
	w, h int
}
func (r Rect) area() int {
	return r.w * r.h
}
func main() {
	r := Rect{w: 10, h: 5}
	fmt.Println("Area:", r.area())
}