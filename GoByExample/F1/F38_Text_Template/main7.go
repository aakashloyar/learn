package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {

    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")

	//"p([a-z]+)ch"

	// This regex means:

	// p → literal "p"

	// ([a-z]+) → one or more lowercase letters (capturing group)

	// ch → literal "ch"
	//"p*****ch"

    fmt.Println(match)

    r, _ := regexp.Compile("p([a-z]+)ch")
	//this i a regularised pattern that we can use to match multiple strings
	fmt.Println(r);

    fmt.Println(r.MatchString("peach"))

    fmt.Println(r.FindString("punch peach punch"))//first match

    fmt.Println("idx:", r.FindStringIndex("peach punch"))//first match index suppose peach starts from 0 to 5

    fmt.Println(r.FindStringSubmatch("peach punch"))//find submatch -> first match and the capturing group

    fmt.Println(r.FindStringSubmatchIndex("peach punch")) //find submatch index

    fmt.Println(r.FindAllString("peach punch pinch", -1))//find all strings matching the pattern

    fmt.Println("all:", r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))//find all strings matching the pattern along with their indexes

    fmt.Println(r.FindAllString("peach punch pinch", 2))//find first 2 strings matching the pattern

	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", 2))//find first 2 strings matching the pattern along with their indexes

    fmt.Println(r.Match([]byte("peach")))//match against byte slice

    r = regexp.MustCompile("p([a-z]+)ch")//no error
    fmt.Println("regexp:", r)

    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))//replace all matching strings with the given string

    in := []byte("a peach")//input byte slice
	fmt.Println(in);
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}