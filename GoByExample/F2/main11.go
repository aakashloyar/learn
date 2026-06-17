package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func findUser(id int) error {
    return fmt.Errorf("db lookup failed: %w", ErrNotFound)
}

func main() {
	err := findUser(42)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("User not fcannoound")
	}
}


// func Is(err, target error) bool
