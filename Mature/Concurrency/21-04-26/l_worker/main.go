package main 
import (
	"fmt"
	"time"
)

func worker(id int,input <-chan int ,output chan<- int) {
	for j:= range input{
		fmt.Println("Starting execution of ", j, "by worker ", id)
		time.Sleep(time.Second)
		output<- j*2
		fmt.Println("Done execution of ", j, "by worker ", id)
	}
}
func main() {
	input:= make(chan int,15)
	output:= make(chan int, 15)
	for i:=0;i<3;i++ {
		go worker(i,input,output)
	}
	for i:=0;i<10;i++ {
		input<-i
	}
	
	close(input)
	for j:=0;j<10;j++ {
		fmt.Println(<-output)
	}
}