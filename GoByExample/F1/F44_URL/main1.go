package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://user:pass@example.com:8080/path/to/resource?name=Aakash&age=21#section1"

	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("User:", u.User)
	fmt.Println("Host:", u.Host)
	fmt.Println("Path:", u.Path)
	fmt.Println("RawQuery:", u.RawQuery)
	fmt.Println("Fragment:", u.Fragment)


	q := u.Query()

	fmt.Println(q["name"])   // [Aakash]
	fmt.Println(q["age"])    // [21]

    name := q.Get("name")
	age := q.Get("age")

	fmt.Println(name) // Aakash
	fmt.Println(age)  // 21


	q.Set("city", "Hyderabad")
	q.Add("skill", "Go")

	u.RawQuery = q.Encode()

	fmt.Println(u.String())


	host := u.Hostname()
	port := u.Port()

	fmt.Println("Host:", host)
	fmt.Println("Port:", port)



}
