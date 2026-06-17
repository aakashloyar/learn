package main

import (
	"fmt"
)

type Op interface{
	Sum(a,b int) int 
}

type Num struct {
	a int 
	b int 
}

func (n Num) Sum(a,b int) int {
	return a+b
}
func main() {
	fmt.Println("Start...")
	var x Num 
	x= Num{2,3}
	fmt.Println(x)
	var y Op 
	y= Num{2,3}
	fmt.Println(y)
}