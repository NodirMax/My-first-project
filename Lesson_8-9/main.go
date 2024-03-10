package main

import (
	"fmt"
	"operations/operations"
)

func main() {
	userMap := make(map[string]operations.User)
	person := operations.User{Name: "John", Age: 32}
	person2 := operations.User{Name: "Bill", Age: 23}
	operations.Create(userMap, "1", person)    // CREATE
	fmt.Println(operations.Read(userMap, "1")) // READ
	operations.Update(userMap, "1", person2)   // UPDATE
	fmt.Println(operations.Read(userMap, "1"))
	operations.Delete(userMap, "1")
	fmt.Println(operations.Read(userMap, "1"))
}
