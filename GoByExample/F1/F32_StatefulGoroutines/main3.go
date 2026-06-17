

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var counter int
// 	var mu sync.Mutex
// 	var wg sync.WaitGroup

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for j := 0; j < 1000; j++ {
// 				mu.Lock()
// 				counter++
// 				mu.Unlock()
// 			}
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("Final counter:", counter)
// }
// //here each go routine is increamenting the variable
// //correct only becoz of mutexes
// //“Everyone is touching the same variable — carefully.”



package main

import (
	"fmt"
	"sync"
)

func main() {
	increment := make(chan int)
	done := make(chan struct{})

	// Stateful goroutine (owns the counter)
	go func() {
		counter := 0
		for inc := range increment {
			counter += inc
		}
		fmt.Println("Final counter:", counter)
		close(done)
	}()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				increment <- 1
			}
		}()
	}

	wg.Wait()
	close(increment) // signal: no more updates
	<-done
}
