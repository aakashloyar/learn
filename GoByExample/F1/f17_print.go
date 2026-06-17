package main 
import "fmt"

// %s → string

// %d → int

// %f → float

// %v → any value

// %T → type

func main() {
	fmt.Println("Closures",40);//print with spaces
	fmt.Printf("My name is %s and I am %d years old.\n","Aakash",20);//formatted print 
	msg:=fmt.Sprintf("My name is %s and I am %d years old.\n","Aakash",20);//does not print, returns string
	print(msg)
}