package main
import (
	"fmt"
	"encoding/json"
)
type User struct {
	Name string  `json:"name"`
	Age  int     `json:"age"`
}


func main() {
	u := User{Name: "Aakash", Age: 21}
	b, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
	fmt.Println(b)
	fmt.Println(string(b))
}