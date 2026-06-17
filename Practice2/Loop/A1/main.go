package main
import "fmt"
func main() {
	for i:=0;i<5;i++ {
		fmt.Println(i)
	}
	for i:=0;;i+=2 {
		if i==2 {
			break;
		}
	}
	i:=0;
	for i<5 {
		i+=2;
		fmt.Println(i)
	}
	// for {
	// 	fmt.Println("YES");
	// }
	var arr [5]int;
	fmt.Println(arr)
	for i,el:=range arr {
		fmt.Println(i,el)
	}
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix[1][2]) // 6
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // 6
}
//go donot have while loop

