package main
import (
 "fmt"
 "time"
)
func main() {
	fmt.Println("Yes")
	x := time.NewTimer(1*time.Second)

	fmt.Println("Waiting for timer...")
	fmt.Println(<-x.C)
	fmt.Println(x)
}

// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     ticker := time.NewTicker(2 * time.Second)

//     for t := range ticker.C {
//         fmt.Println("Tick at:", t)
//     }
// }
