package main

import (
	"errors"
	"fmt"
)

func main() {
	z, err := divide(6, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(z)
	}
}

func divide(a, b int) (result float64, err error) {
	if b == 0 {
		err = errors.New("Number below than 1")
		return result, err
	}
	result = float64(a) / float64(b)
	return result, err
}
