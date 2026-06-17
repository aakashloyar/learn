package main 
import (
	"fmt"
	"sync/atomic"
)

func main() {
	var c int64=0
	atomic.StoreInt64(&c,0)
	atomic.AddInt64(&c,1)
	atomic.LoadInt64(&c)
	atomic.CompareAndSwapInt64(&c,1,30)
	fmt.Println(atomic.LoadInt64(&c))
		
}

//cas
// It does:
//	atomic.CompareAndSwapInt64(&c,old,new)
// Check if x == old
// If yes → set x = new and return true
// If no → do nothing and return false