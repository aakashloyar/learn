package main

import "fmt"

type Direction int

const (
	North Direction = iota // 0
	East                   // 1
	South                  // 2
	West                   // 3
)

func main() {
	fmt.Println(North, East, South, West) // 0 1 2 3
}

//go does not have direct enum 
//it has iota