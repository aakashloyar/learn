package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	//defer ticker.Stop()

	for t := range ticker.C {
		fmt.Println("Tick at:", t)
	}//like a set tinterval
}
