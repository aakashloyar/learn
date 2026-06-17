package main

import (
    "fmt"
    "strconv"
)

func main() {

    f, _ := strconv.ParseFloat("1.234", 64)//64 means return a float64
    fmt.Println(f)

    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)
	//ParseInt(s string, base int, bitSize int)
	//base 0 ->auto detect base from string prefix
	//0x1c8
	//from here it will know its hexadecimal// base 16//from 0x
	//bitSize 64 -> return int64
	//1c8= 1*16^2 + 12*16^1 + 8*16^0 = 256 + 192 + 8 = 456

    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    k, _ := strconv.Atoi("135")
    fmt.Println(k)
	//atoi ->ascii to integer

    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}