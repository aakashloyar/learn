package main
import "fmt"

func main() {
	arr:=make([]int,4);
	fmt.Println(arr);
	rows := [][]int{};
	fmt.Println(rows);
	var res [3]int;
	fmt.Println(res);

	numRows, numCols := 3, 4

	// Create a slice with `numRows` elements
	matrix := make([][]int, numRows)

	// For each row, create a slice with `numCols` elements
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]int, numCols)
	}
}

