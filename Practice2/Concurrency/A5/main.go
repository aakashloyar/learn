package main
import "fmt"
func add(s,e int) int{
	sum:=0;
	for i:=s;i<e;i++ {
		sum+=i;
	}
	return sum;
}
func main() {
	ch:=make(chan int,100)
	for i:=0;i<1000;i+=100 {
		//ch<-add(i,min(i+100,1000))
		start, end := i, min(i+100, 1000)
		go func(s, e int) { // ✅ run in goroutine
			ch <- add(s, e)
		}(start, end)
	}

	// close(ch)//this will close it immediately
	// sum:=0;
	// for item := range ch {
	//     sum+=item
	// }
	sum := 0
	for i := 0; i < 1000; i += 100 { // read exactly 10 results
		sum += <-ch
	}
	fmt.Println(sum)
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}