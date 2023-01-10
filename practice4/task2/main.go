package main

import "fmt"

func main() {
	fmt.Println("Три числа.")

	fmt.Println("Введите первое число:")
	var number1 int
	fmt.Scan(&number1)

	fmt.Println("Введите второе число:")
	var number2 int
	fmt.Scan(&number2)

	fmt.Println("Введите третье число:")
	var number3 int
	fmt.Scan(&number3)

	max := 5

	if number1 > max {
		fmt.Println("Среди веденных чисел есть число больше 5.")
		return
	} else if number2 > max {
		fmt.Println("Среди введенных чисел есть число больше 5.")
		return
	} else if number3 > max {
		fmt.Println("Среди введенных чисел есть число больше 5.")
		return
	}

	fmt.Println("Среди введенных чисел нет числа больше 5.")
}
