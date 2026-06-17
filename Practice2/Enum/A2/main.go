//go does not have default enum for string
package main
import "fmt"
type Priority int

const (
	Low Priority = iota + 1 // starts at 1
	Medium
	High
)
func (p Priority) String() string {
	switch p {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println(Low, Medium, High) // Low Medium High
}
