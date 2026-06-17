package main
import "fmt"
func main() {
	str:="Hello"
	fmt.Println(len(str))
	fmt.Println(str[1])
	fmt.Println(str+" World")
	//fmt.Println(str+" World"[6])

	for i,ch:=range str {
		fmt.Println(i,ch)
	}

	// s := "नमस्ते"
    // fmt.Println(len(s)) 
	// fmt.Println(s)

	// for i, r := range "नमस्ते" {
	// 	fmt.Printf("%d: %c\n", i, r)
	// }


	s := "hello🙂"
	r := []rune(s)

	fmt.Println(r)
	fmt.Printf("%c\n", r[5]) // 🙂

}
//byte-> 8 bits value
//rune-> int32 value

// ASCII → 1 byte

// Hindi, Chinese → 2–3 bytes

// Emojis → 4 bytes

//in other langauges-> 1 single emoji uses 2 characters //called surrogate pairs
//Swift understands “what humans see” as one character.