package main 
import "fmt"
//go does not have real enums like other languages
//we can use const and iota to create enums

//basic enum using iota
const (
	a=iota
	b
	c
	d
)

//enum with custom type
type Size int

const (
    Small Size = iota
    Medium
    Large
)

//enu with string output
func (s Size) String() string {
    switch s {
    case Small:
        return "Small"
    case Medium:
        return "Medium"
    case Large:
        return "Large"
    default:
        return "Unknown"
    }
}


//bit signed enum
type Permission uint8

const (
    Read Permission = 1 << iota
    Write
    Execute
)
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(Small)
	fmt.Println(Medium)
	fmt.Println(Large)

	s := Medium
    fmt.Println(s)      // Medium
	p := Read | Write
	fmt.Println(p)      // 3
}