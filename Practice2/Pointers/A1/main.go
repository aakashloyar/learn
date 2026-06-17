package main
import "fmt"

func main() {
	x := 42
	var p *int=&x;
	var q **int=&p;
	fmt.Println(x,p,*p,q,*q,**q);

}

