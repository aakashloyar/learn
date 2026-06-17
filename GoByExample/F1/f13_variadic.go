package main

import "fmt"


//...int -> creating into array
//x...-> passing array as variadic argument
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0

    for _, num := range nums {
        total += num
    }
	// for num of nums {
	// 	total += num;
	// }
    fmt.Println(total)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}