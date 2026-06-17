package main
import "fmt"
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
    mySlice := primes[1:4]
    // mySlice = {3, 5, 7}
	fmt.Println(mySlice)
	nums := []int{10, 20, 30} // directly creates slice
	fmt.Println(nums)

	s := make([]int, 3, 5) // length=3, capacity=5//this help in reducing realocation//length //cap
    fmt.Println(s)         // [0 0 0]

	p := []int{1, 2, 3}
	p = append(p, 4, 5)
	fmt.Println(p) // [1 2 3 4 5]

}
//go donot have while loop

