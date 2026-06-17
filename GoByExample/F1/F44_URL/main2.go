package main

import (
    "crypto/sha256"
    "fmt"
)

func main() {
    s := "sha256 this string"

    h := sha256.New()//create a new hash object

    h.Write([]byte(s))//convert string to byte slice and write to hash object

    bs := h.Sum(nil)//return the hash

    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}