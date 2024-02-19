package myFunc

import "fmt"

func Day1_Task1(a float64) float64 {
	return 4 * a
}

func Day1_Task2(a, b float64) float64 {
	return 2 * (a + b)
}

func Day1_Task3(a, b float64) float64 {
	return a * b
}

func Day1_Task4(age int) int {
	return (10 - age%10) % 10
}

// Таблица умножения
func Bonus_Task1() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d * %d = %d", i, j, i*j)
			fmt.Println()
		}
	}
}

// Числа Фибоначчи
func Bonus_Task2(a int) int {
	if a == 1 || a == 2 {
		return 1
	}
	return Bonus_Task2(a-1) + Bonus_Task2(a-2)
}

func Summa(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b float64) float64 {
	if b == 0 {
		panic("Деление на 0")
	} else {
		return a / b
	}
}
