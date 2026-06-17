package main 
import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs<-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("Worker", id, "processing job", job)
	}
	
}
func main() {
	njobs:=10
	nworkers:=3
	jobChan:=make(chan int,njobs)
	var wg sync.WaitGroup

	//create workers
	for i:=1;i<=nworkers;i++ {
		wg.Add(1)
		go worker(i, jobChan, &wg)
	}
	//send jobs 
	for j:=1;j<=njobs;j++ {
		time.Sleep(1 * time.Second)
		jobChan<-j
	}
	close(jobChan)
	wg.Wait()
}

//they are just randomly picking the jobs from the channel and executing them