package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int = 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	n := 1000
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			mu.Lock()   // 🔒 lock before critical section
			counter++  // safe
			mu.Unlock() // 🔓 unlock after
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}
//mu.Lock()
//defer mu.Unlock()



// RWMutex (Read-Write Mutex)

// Use when:

// Many goroutines read

// Few goroutines write

// var mu sync.RWMutex

// mu.RLock()   // many readers allowed ✅
// mu.RUnlock()

// mu.Lock()    // only one writer allowed ❌❌
// mu.Unlock()


// ✅ Faster than normal mutex when reads dominate.



// Mutex vs Atomic (Which Should You Use?)
// Situation	Use This
// Just incrementing a number	✅ atomic
// Updating a struct	✅ mutex
// Multiple related variables	✅ mutex
// Just reading & writing int	✅ atomic
// Maps, slices	✅ mutex
// Business logic	✅ mutex
// 🚀 Rule of Thumb:

// If it's more than one line → use a mutex.


// WHICH ONE SHOULD YOU USE?

// Since you're learning backend + system design, you should:

// ✅ Use mutexes by default
// ✅ Use atomic only for hot counters & metrics