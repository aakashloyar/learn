package main

import (
	"encoding/xml"
	"fmt"
)

type User struct {
	XMLName xml.Name `xml:"user"`
	Skills  []string `xml:"skills>skill"`
}

func main() {
	u := User{
		Skills: []string{"Go", "Docker", "Kubernetes"},
	}

	data, err := xml.MarshalIndent(u, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}


// type User struct {
// 	XMLName xml.Name `xml:"user"`
// 	Skills  []string `xml:"skills>skill"`
// }

// <user>
//   <skills>
//     <skill>Go</skill>
//     <skill>Next.js</skill>
//   </skills>
// </user>
