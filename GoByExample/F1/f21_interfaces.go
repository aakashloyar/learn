package main
import (
    "fmt"
    //"math"
)

type rect struct {
    width, height float64
}

type circle struct {
	radius float64
}

type geometry interface {
    area() float64
    perim() float64
}
//required methods for interface
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * 3.14 * c.radius
}

//****************************
func detectCircle(g geometry) {
    if c, ok := g.(circle); ok {
        fmt.Println("circle with radius", c.radius)
    } else {
		fmt.Println("not a circle")
	}
}

//****************************
func measure(g geometry) {
    fmt.Println(g.area())
    fmt.Println(g.perim())
}
//****************************

func main() {
	r := rect{width: 3, height: 4}
	fmt.Println(r.area())

	//here it donot require interface 
	//but if we have multiple types implementing same methods
	//then we can use interface to generalize the code

	c := circle{radius: 5}
	measure(r)
	measure(c)
	
}

// Struct = "This is what I am."

// Method = "This is what I can do."

// Interface = "Anyone who can do X and Y is acceptable to me."