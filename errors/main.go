package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func main() {
	
	err := randomError(0);

	// is neccessary to create it as a pointer to use errors.As
	//var myErr *MyError
	//errors.As(err, &myErr)

	// if we can cast the error to MyError
	if errors.As(err, new(*MyError)) {
		fmt.Println("my error")
	} else {
		fmt.Println("unknown error")
	}
}

func randomError(i int) error {
	if i == 0 {
		return errors.New("random error")
	}
	return &MyError{Message: "my error"}
}

func ptr[T any](v T) **T {
	a := &v
	return &a
}
