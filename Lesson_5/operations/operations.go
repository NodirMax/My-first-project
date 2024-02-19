package operations

import (
	"fmt"
	"math/rand"
	"students/students"
)

func Cheaking(student students.Student) students.Student {
	if (student.Age < 15) || (student.Age > 19) {
		fmt.Println("Вы не можете присутствовать на экзамене")
	}
	student.Nomer_bileta = rand.Intn(30) + 1
	return student
}

func Exam(student students.Student, predmet students.Subject) (students.Student, students.Subject) {
	if float64(predmet.Kolichestvo_propuskov) >= 1/2*float64(predmet.Kolichestvo_urokov) {
		student.Proshel = false
	}
	if predmet.Osenka == 5 {
		student.Proshel = true
	}
	return student, predmet
}
