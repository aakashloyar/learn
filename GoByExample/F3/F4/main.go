package main
import (
	"fmt"
)

type Inter interface {
	Sum(a,b int) int 
	Change(other int)
}
type X struct {
	a int 
	b int 
}

func (s X) Sum(a,b int) int {
	return a+b
}  
func (s *X) Change(other int) {
	s.a=other
}
func main() {

	a := &X{2,3}
	fmt.Println(a);
	a.Change(200)
	fmt.Println(a);

	var b Inter 
	b = &X{2,3}//need to pass by pointer as now this time struct has a pointer method 
	//fmt.Println(b.Sum(10,20))

	b.Change(200)
	fmt.Println(b)

	var i Inter
	i = &X{2, 3}   // value works
	i.Change(200)
	fmt.Println(i)
	//fmt.Println(i.Sum(10, 20))
}