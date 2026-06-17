package main 

import (
	"fmt"
	"encoding/json"
	"log"
)

type Men struct {
	Id   string `json:"id"`
	Name string `json:"name"`//if you are not capitalising these fields you will not get them
}

func main() {
	fmt.Println("YES")
	track := Men {
		Id: "aalu",
		Name: "aakash",
	}
	data, err := json.Marshal(track)//go struct -> json bytes
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Println(data)//this will return character
	fmt.Println(string(data))

	jsonData := []byte(`{"id":"1","title":"Believer","artist":"Imagine Dragons"}`)
	var x Men 
	err = json.Unmarshal(jsonData, &x)//json bytes-> go struct  
	if err != nil {
		log.Fatal(err)
	}//now here if will take the values which are present leave everything else empty
	fmt.Println(x)
	fmt.Println(x.Id)
	fmt.Println(x.Name)
	var y Men 
	err = json.Unmarshal(data,&y)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(y)
}

