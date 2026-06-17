package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println("i =", i)
    }
	i := 0
    for i < 5 {
        fmt.Println("i =", i)
        i++
    }
	// for {
    //     fmt.Println("runs forever")
    // }
	for i := 0; i < 10; i++ {
        if i == 3 {
            continue // skip 3
        }
        if i == 7 {
            break // stop at 7
        }
        fmt.Println("i =", i)
    }
	numbers := []int{10, 20, 30, 40}
	for idx, val := range numbers {
		fmt.Println("index:", idx, "value:", val)
	}
	m := map[string]int{"a": 1, "b": 2}
	for key, val := range m {
		fmt.Println("key:", key, "value:", val)
	}
	s := "hello"
	for i, r := range s {
		fmt.Printf("index %d, rune %c\n", i, r)
	}

	primes := [6]int{2, 3, 5, 7, 11, 13}
    lice := primes[1:4]
	fmt.Println(lice)

	mySlice := []string{"I", "love", "go"}
    fmt.Println(len(mySlice)) // 3
}
