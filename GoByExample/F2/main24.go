package main 
import (
	"fmt"
	s "strings"
)


var temp = map[string]bool{
    "password": true,
    "secret":   true,
}

type Logger struct {}

type SecureLogger struct {
    //Logger//exported everywhere//public 
	logger Logger//is not exported outside pacakge//private
}

func (l Logger) Log(msg string) {
    fmt.Println(msg)
}

func (l SecureLogger) Log(msg string) {
	arr:=s.Split(msg," ");
	res:=""
	for _, x := range arr {
		if _, ok:= temp[x];ok {
			res+="*** "
		} else {
			res+=x+" "
		}
	}
	//fmt.Println(res)
	l.logger.Log(res)
}

func main() {

	l := SecureLogger{Logger{}}
	l.Log("my password is 123") 
	// output: "my ****** is 123"
	//l.Logger.Log("my password is 123") 


}