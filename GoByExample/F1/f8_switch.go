package main
import (
	"time"
	"fmt"
)
func main() {
	i:=8;
	switch i {
		case 1: {
			fmt.Println("one");
		}
		case 2: {
			fmt.Println("two");	
		}
		case 3: {
			fmt.Println("three");
		}
		default: {
			fmt.Println("unknown");
		}
	}
	//t:=time.Now().Weekday();
	//fmt.Println(t);
	fmt.Println(time.Now());
	fmt.Println(time.Now().Weekday());
	fmt.Println(time.Now().Hour());


	whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
	
}