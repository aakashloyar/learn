//Is -> is this error equivalent to target error
//As -> is this error of target type


package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Msg string 
	code int
}

// As checks if the error is of the target type.
//func As(err error, target any) bool
//address of target should be passed

func (v *MyError) Error() string {
    return fmt.Sprintf("field %s: %s", v.Msg, v.code)
}

func main() {
	var err error = &MyError{Msg: "something went wrong", code: 500}
	var target *MyError
	if errors.As(err, &target) {
		fmt.Println("both error types matches")
	}
}