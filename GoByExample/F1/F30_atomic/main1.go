package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 // shared counter

	var wg sync.WaitGroup
	n := 1000
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			// increment atomically
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final counter:", atomic.LoadInt64(&counter))
}
//atomic let you do operations on shared variables safely without using mutexes


// atomic.AddInt64(&x, delta) — add and return new value (safe concurrently).

// atomic.LoadInt64(&x) — read value atomically.

// atomic.StoreInt64(&x, val) — write value atomically.

// atomic.CompareAndSwapInt64(&x, old, new) — CAS loop when needed.