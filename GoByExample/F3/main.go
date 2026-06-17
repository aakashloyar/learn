package main

// import (
// 	"fmt"
// )

// type Inter interface {
// 	Sum(a,b int) int 
// }

// type S struct { 
// 	a int 
// 	b int 
// }
// func (s S) Sum(a,b int) int {
// 	return a+b
// }

// type check struct {
// 	a int 
// 	b int 
// 	x Inter
// }

// func (c check) Execute() {
// 	y:=c.x.Sum(c.a,c.b)
// 	fmt.Println(y)
// }

// func main() {
// 	// x:=S{}
// 	// fmt.Println(x)

// 	// var y Inter 
// 	// y= S{0,0}
// 	// fmt.Println(y)
// 	z:=check{a: 2,b: 3,x: S{}}
// 	z.Execute()
	
// }