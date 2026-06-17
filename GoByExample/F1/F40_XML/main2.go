package main

import (
	"encoding/xml"
	//"os"
	"fmt"
	//"strings"
)

type User struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func main() {
	u := User{"Aakash", 21}
	b, _ := xml.MarshalIndent(u, "", "  ")
	fmt.Println(string(b))
}

