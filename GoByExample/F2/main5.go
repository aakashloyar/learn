package main

import "fmt"

type rect struct {
    width, height int
}

//can change the parameter
func (r *rect) area() int {
	r.width++;
    return r.width * r.height
}

//cannot change the parameters
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())//
    fmt.Println("perim:", r.perim())// 

    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())

	fmt.Println(r.width)

	//go automaticall handles method calls with pointer/value receivers
	//make it easier to switch between pointer and value receivers
}