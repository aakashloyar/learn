package main
import "fmt"
func main() {
	name := "Aakash"
    age := 21
    msg := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
	fmt.Println(msg)
	n, err := fmt.Printf("Hello, %s! You are %d years old.", name, age)
	fmt.Println(n, err)
	
}
//Print:- printing
//Printf:- formatted print 
//Sprintf:- returns string donot print anyting
//%v -> default print
//%t -> boolean
