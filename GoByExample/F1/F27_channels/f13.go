//✅ 6. LLM / Streaming Token Timeout
package main
import (
	"fmt"
	"time"
)


func main() {
	stream:=make(chan string,2)
	go func() {
		tokens:=[]string{"tic ","tac ","toe "}
		for _,token:=range tokens{
			stream<-token
			fmt.Println(token+" Inside")
			time.Sleep(1 *time.Second)
		}
		//close(stream)
	}()
	for {
		select {
			case token := <-stream:
				fmt.Println(token)

			case <-time.After(3 * time.Second):
				fmt.Println("\nStream timed out")
				return
		}
	}

}