package main

import (
	"fmt"
	"students/operations"
	"students/students"
)

func main() {
	human := students.Student{Name: "Алеша Попов", Age: 17, Klass: 9, Proshel: false}
	predmet := students.Subject{Name: "Math", Osenka: 4, Kolichestvo_propuskov: 12, Kolichestvo_urokov: 60}
	human = operations.Cheaking(human)
	fmt.Println(human)
	fmt.Println(predmet)
}
