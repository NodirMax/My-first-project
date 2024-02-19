package main

import (
	"fmt"
	"log"
	"myFunc/myFunc"
)

func main() {
	// Приветствие
	fmt.Println("Здравствуйте! Пожалуйста, выберите день обучения:\n- 1 День. Знакомство\n- 2 День. Первая программа\n- 3 День. Функции")
	var user_input int
	fmt.Scan(&user_input)
	switch user_input {
	case 1:
		// ДЕНЬ 1

		fmt.Println("Первый день. Это день Знакомство")

	case 2:
		// ДЕНЬ 2

		fmt.Println("День 2. Выберите номер задачи.\n1 \t 2\t 3\t 4")
		var user_number_of_task int
		fmt.Scan(&user_number_of_task)
		switch user_number_of_task {

		case 1:
			// Условие задачи
			fmt.Println("Условие задачи!\n Написать программу, которая находит периметр квадрата. Пользователь должен сам вводить числа.")
			// Ввод пользователья
			var storona_kvadrata float64
			fmt.Scan(&storona_kvadrata)
			// Вызов функции из пакета features
			fmt.Println(myFunc.Day1_Task1(storona_kvadrata))

		case 2:
			// Условие задачи
			fmt.Println("Условие задачи!\n Написать программу, которая находит периметр прямоугольника. Пользователь должен сам вводить числа.")
			// Ввод пользователья
			var a, b float64
			fmt.Scan(&a, &b)
			// Вызов функции из пакета features
			fmt.Println(myFunc.Day1_Task2(a, b))

		case 3:
			// Условие задачи
			fmt.Println("Условие задачи!\n Программа, которая находит площадь прямоугольника. Пользователь должен сам вводить числа.")
			// Ввод пользователья
			var a, b float64
			fmt.Scan(&a, &b)
			// Вызов функции из пакета features
			fmt.Println(myFunc.Day1_Task3(a, b))

		case 4:
			// Условие задачи
			fmt.Println("Условие задачи!\n Написать программу, которая будет вычислять, сколько пользователю ещё нужно лет до его юбилея.\n Т.е, если ему 25 лет, то до юбилея ему нужно ещё 5 лет.\n Если ему 30 лет,то до юбилея ему нужно: 0\n Если 37, то до юбилея ему нужно: 3\n Пользователь, очевидно, должен вводить свой возраст сам.\n Использовать можно только пройденный нами материал.\n То есть использование условий 'if' не нужно.")
			// Ввод пользователья
			fmt.Println("Вводите возраст пользоваьеля")
			var age int
			fmt.Scan(&age)
			log.Println("Программа запущена")
			if myFunc.Day1_Task4(age) == 0 {
				fmt.Println("Юбилей уже наступиль")
			} else {
				fmt.Printf("До юбилея осталось: %d", myFunc.Day1_Task4(age))
				fmt.Println()
			}
			log.Println("Программа завершена")
		}

	case 3:
		fmt.Println("Вводите числа для вывода: Суммы, Произведения, Деления")
		// Ввод от пользователя
		var a, b int
		fmt.Scan(&a, &b)
		// Вызов функции из пакета features
		fmt.Println("Вывод результатов:")
		fmt.Println(myFunc.Summa(a, b))
		fmt.Println(myFunc.Multiply(a, b))
		fmt.Println(myFunc.Divide(float64(a), float64(b)))

	default:
		fmt.Println("Извините, но такого задания нет\nХотите выбрать бонусные задачи\n1: Таблица умножения  2:числа Фибоначчи")
		var user_number_of_task int
		fmt.Scan(&user_number_of_task)

		switch user_number_of_task {
		case 1:
			fmt.Println("Условия:\n При помощи циклов вывести в терминал таблицу умножения.")
			// Вызов функции из пакета features
			myFunc.Bonus_Task1()

		case 2:
			fmt.Println("Условия:\n Используя рекурсию написать функцию, вычисляющую числа Фибоначчи.")
			var user_input int
			fmt.Scan(&user_input)
			// Вызов функции из пакета features
			fmt.Println(myFunc.Bonus_Task2(user_input))

		}
	}
}
