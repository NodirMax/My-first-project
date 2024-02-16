package main

import (
	"fmt"
	"strings"
)

func main() {

	// var stroka string
	// fmt.Scan(&stroka)
	// stroka_1 := []rune(stroka)
	// fmt.Println(len(stroka_1))
	z := "Hello man "
	fmt.Println(len(z))
	result := strings.TrimSpace(z)
	fmt.Println(len(result))

}
