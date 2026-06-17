package main 
import (
	"fmt"
	//"slices"
)
func main() {
	s:=[]string{"a","b","c"}
	fmt.Println(s[1:3])
	fmt.Println(s[:2])
	fmt.Println(s[1:])
	fmt.Println(s[:])
	fmt.Println(s);
	s=append(s,"e")
	s=append([]string{"f"},s...)
	//s=append("f","g")//wrong
	fmt.Println(s)

	//make will create slice with specified length and capacity
	//now it will have some reserved space to avoid multiple allocations
	//if length exceeds capacity, it will allocate new underlying array with double capacity
	//if capacity is not specified, it will be equal to length
	//while if we create array, its length is fixed
	//when we add elements to such array, we need to create new array and copy elements

	st := make([]string, 3, 10)//each string initialized to ""
	fmt.Println(st);
	fmt.Println(len(st),cap(st))
	fmt.Println(st[2])
	arr:=[3]int{1,2,3}
	slice:=arr[:]
	slice[0]=0;
	fmt.Println(arr);
	//now just see the magic
	slice=append(slice,4)
	slice[1]=9;
	fmt.Println(arr);//array changed as underlying array is changed
}