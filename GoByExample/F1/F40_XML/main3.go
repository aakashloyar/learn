package main

import (
    "encoding/xml"
    "fmt"
)

type Person struct {
    Name string `xml:"name"`
    Age  int    `xml:"age"`
}

func main() {
    data := []byte(`<person><name>Aakash</name><age>21</age></person>`)

    var p Person
    xml.Unmarshal(data, &p)

    fmt.Println(p)
}
