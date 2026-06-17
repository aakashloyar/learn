package main
import (
	"fmt"
	"maps"
)
func main() {
	//An empty map created with make(map[K]V) starts with 1 bucket internally.
    //Each bucket can hold 8 key/value pairs.
    //When more space is needed, Go doubles the number of buckets and redistributes keys.
	m:=make(map[string]int)
	m["k1"]=7;
	m["k2"]=13;
	fmt.Println("map:",m);
	fmt.Println("len:",len(m));
	delete(m,"k2")
	fmt.Println("map:",m);

	 _, prs := m["k2"]
	fmt.Println("prs:",prs)
	
	

	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }

	for k,v:=range n {
		fmt.Println(k,v);
	}
}
