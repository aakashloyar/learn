package main

import "fmt"

//variadic
// func concat(strs ...string) string {
//     final := ""
//     // strs is just a slice of strings
//     for i := 0; i < len(strs); i++ {
//         final += strs[i]
//     }
//     return final
// }

// func main() {
//     final := concat("Hello ", "there ", "friend!")
//     fmt.Println(final)
//     // Output: Hello there friend!
// }


//spread operator
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}