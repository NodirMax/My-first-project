package main

import "fmt"

func main() {
	//--- 1 ---
	// var a int
	//fmt.Scan(&a)
	// fmt.Println(4*a)

	//--- 2 ---
	// var a, b int
	// fmt.Scan(&a, &b)
	// fmt.Println(2 * (a + b))

	//--- 3 ---
	// var a, b int
	// fmt.Scan(&a, &b)
	// fmt.Println(a*b)

	//--- 4 ---
	var age uint
	fmt.Scan(&age)
	fmt.Println((10 - age%10) % 10)
}
