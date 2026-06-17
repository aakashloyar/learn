package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := map[string]interface{}{
		"name": "Aakash",
		"age":  21,
		"skills": []string{"Go", "Next.js"},
	}

	b, err := json.MarshalIndent(data, "#", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

//b, err := json.MarshalIndent(data, prefix, indent)
//prefix is added at the beginning of each line
//indent-> put indentation at each level  l1->1*l  l2->2*l