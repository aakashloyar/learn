package main

import "fmt"

func plus(a int, b int) int {

    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}


//multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

//named return values

func add(a, b int) (sum int) {
    sum = a + b
    return  // returns sum automatically
}


//function as values
f := func(a, b int) int {
    return a + b
}

fmt.Println(f(3, 4))


func update(x *int) {
    *x = 100
}


//defer run after the function completes
func test() {
    defer fmt.Println("world")
    fmt.Println("hello")
}

func main() {

    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
	a := 10
    update(&a)
    fmt.Println(a) // 100
}