package main
import "fmt"
func main() {
	var arr [5]int;
	fmt.Println("emp:", arr);
	arr[4]=100

	arr1:=[5]int{1,2,3,4,5}
	fmt.Println("dcl:", arr1);

	//var arr2 [3][3]int={{1,2,3},{2,3,4},{4,5,6}}//not allowed
	arr2:= [2][3]int{{1,2},{4,5,6}};
	fmt.Println("2d:", arr2);
}