package main
import (
	"fmt"
)

func Min[T ~int | ~float64 |~string ](a, b T) T {
    if a < b {
        return a
    }
    return b
}

func main() {
	a:=20;
	b:=10;
	fmt.Println("Min is:",min(a,b));
	fmt.Println("Min is:",min("string1","string2"));
	//fmt.Println("Min is:",min([]int{1,2,3},[]int{1,2,3,4}));
}
