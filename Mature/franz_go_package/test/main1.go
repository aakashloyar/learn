package main

import (
	"fmt"
)
type Man struct {
	Name *string
	Age  int
	check bool 
}

func NewMan(name string, age int, is bool) *Man{
	return &Man {
		Name: &name,
		Age: age,
		check: is,
	}
}
func main() {
	//var Man man 
	man := NewMan("aakash",18,true)
	fmt.Println(man)
	fmt.Println(*man)
	//fmt.Print("%v/n",*man)
	fmt.Printf("%v\n",*man)
	fmt.Printf("%+v\n",*man)
	fmt.Printf("%#v\n",*man)
	fmt.Println((*man).Name)
	fmt.Println(*man.Name)//will dereference all the pointers
}