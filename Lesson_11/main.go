package main

import (
	"fmt"
	pluschange "function/PlusChange"
	shifrsezary "function/Shifr_Sezary"
	"function/function"
)

func main() {
	f := function.Conv_to_byte("Supra")
	fmt.Println("Конвертация строки в байт: ", f)
	fmt.Println("Конвертация байта в строку: ", function.Conv_to_string(f))
	z := shifrsezary.Shifrovka("hello")
	fmt.Println("Шифровка:", z)
	fmt.Println("Расшефровка:", shifrsezary.RasShifrovka(z))
	stroka := "Мы+получили+текст,+в+котором+по+какой-то+ошибке+все+пробелы+заменились+на+плюсы.+Надо+это+исправить"
	fmt.Println(pluschange.Change(stroka))
}
