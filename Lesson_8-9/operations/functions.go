package operations

import "fmt"

type User struct {
	Name string
	Age  int
}

func Create(userMap map[string]User, key string, value User) {
	userMap[key] = value
}

func Read(userMap map[string]User, key string) User {
	return userMap[key]
}

func Update(userMap map[string]User, key string, value User) {
	if _, ok := userMap[key]; ok {
		userMap[key] = value
	} else {
		fmt.Printf("Нет такого ключа '%s' в мапе.\n", key)
	}
}

func Delete(userMap map[string]User, key string) {
	if _, ok := userMap[key]; ok {
		delete(userMap, key)
	} else {
		fmt.Printf("Нет такого ключа '%s' в мапе.\n", key)
	}
}
