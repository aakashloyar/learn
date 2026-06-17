// Key Components
// Component	Purpose
// Task Queue	Stores pending jobs
// Workers	Threads/goroutines that execute tasks
// Dispatcher	Assigns tasks to workers
// Concurrency Limit	Max number of workers

package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("Worker", id, "processing job", job)
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Create workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}



// for {
//     job, ok := <-jobs
//     if !ok {
//         break
//     }
//     fmt.Println(job)
// }//this is range over channel means //ok become false only if empty and closed
