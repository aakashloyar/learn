package main 
import "fmt"

func main() {
	s := "ह"
    fmt.Println(len(s)) // 3 bytes
	s = "hello世界"
	r := []rune(s)
	fmt.Println(len(r)) // 7 characters
	for _, ch := range s {
		fmt.Println(ch, string(ch))
	}


	var r rune = 'A'
	var r2 rune = 'ह'
	var r3 rune = '😀'

	fmt.Println(r, r2, r3)


}