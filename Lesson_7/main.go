package main

import (
	"Reverser/Reverser"
	"Reverser/SREZ"
	"fmt"
)

func main() {
	massiv := [...]interface{}{"Hello", "how", "are", "you", "?"}
	fmt.Println(Reverser.RVS(massiv[:]))
	user_input := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(SREZ.RNS(user_input[:]))
}
