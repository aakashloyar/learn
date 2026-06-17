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
	data := []byte(`{"name":"Aakash","age":21}`)
	var u User

	err := json.Unmarshal(data, &u)//storing data into go value
	if err != nil {
		panic(err)
	}

	fmt.Println(u)
}