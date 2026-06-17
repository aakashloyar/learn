package main
import "fmt"


func main() {
	//anonymous struct
	myCar := struct {
	brand string
	model string
	} {
	brand: "Toyota",
	model: "Camry",
	}
	fmt.Println(myCar)
}