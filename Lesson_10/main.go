package main

import (
	"errors"
	"fmt"
)

func Hello(s string) (string, error) {
	if len(s) < 20 {
		err := errors.New("Len less than 20")
		return "", err
	}
	return s, nil
}

func main() {
	z := "Just do it!"
	fmt.Println(Hello(z))
}
