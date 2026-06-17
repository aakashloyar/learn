package main

import (
	"encoding/xml"
	"fmt"
)

type User struct {
    Name string `xml:"name"`
    Age  int    `xml:"age"`
}

func main() {
	data := []byte(`
	<user>
		<name>Aakash</name>
		<age>21</age>
	</user>
	`)

	var u User
	xml.Unmarshal(data, &u)
	fmt.Println(u.Name, u.Age)

}

//verbosity-> extra text around the data
//json<xml //in verbosity
//xml is designed for documents not for api data exchange




// | Feature           | XML    | JSON  |
// | ----------------- | ------ | ----- |
// | Verbosity         | High ❌ | Low ✅ |
// | Human readability | Medium | High  |
// | Attributes        | ✅      | ❌     |
// | Schema maturity   | ✅      | ❌     |
// | Namespaces        | ✅      | ❌     |
// | Document support  | ✅      | ❌     |
// | APIs              | ❌      | ✅     |



// Final mental model 🧠

// JSON → app-to-app communication

// XML → structured documents + strict contracts