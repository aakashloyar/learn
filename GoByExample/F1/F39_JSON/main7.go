package main

import (
	"encoding/json"
	"os"
	"fmt"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u1 := User{Name: "Aakash", Age: 21}
    json.NewEncoder(os.Stdout).Encode(u1)


	data := strings.NewReader(`{"name":"Aakash","age":21}`)

	var u User
	json.NewDecoder(data).Decode(&u)

	fmt.Println(u)


}

// Marshal / Unmarshal → work with []byte
// Encoder / Decoder   → work with streams (io.Writer / io.Reader)


//Go value → JSON → []byte (in memory)


//Go value → JSON → written directly to stream  //encoder
// JSON from stream → Go value   //decoder
