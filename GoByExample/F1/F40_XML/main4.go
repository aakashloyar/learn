package main

import (
	"encoding/xml"
	"fmt"
)

type User struct {
	XMLName xml.Name `xml:"user"`
	ID      int      `xml:"id, attr"`//serializing as attribute //not as element
	Name    string   `xml:"name"`
}

func main() {
	u := User{
		ID:   101,
		Name: "Aakash",
	}

	data, err := xml.MarshalIndent(u, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}


