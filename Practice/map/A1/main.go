package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 24
	ages["Mary"] = 21 // overwrites 24
	fmt.Println(len(ages)) // 2
	ages["John"] = 26
	my_age:= ages["John"]
	fmt.Println(my_age)
	delete(ages, "Mary")
	fmt.Println(ages);
	elem, ok := ages["John"]
	fmt.Println(elem,ok);
}