package operations

import (
	"testing"
)

func TestCreate(t *testing.T) {
	userMap := make(map[string]User)
	key := "1"
	value := User{Name: "Alice", Age: 19}
	Create(userMap, key, value)
	if userMap[key] != value {
		t.Errorf("Unexpected value for key '%s'. Expected: %v, Got: %v", key, value, userMap[key])
	}
}

func TestUpdate(t *testing.T) {
	userMap := map[string]User{"1": {Name: "Alice", Age: 19}}
	key := "1"
	value := User{Name: "John", Age: 12}
	Update(userMap, key, value)
	if userMap[key] != value {
		t.Errorf("Unexpected value for key '%s'. Expected: %v, Got: %v", key, value, userMap[key])
	}
}

func TestDelete(t *testing.T) {
	value := User{Name: "Alice", Age: 19}
	userMap := map[string]User{"1": value}
	key := "1"
	Delete(userMap, key)
	if userMap[key] == value {
		t.Errorf("ERROR")
	}
}
