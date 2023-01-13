package main

import "fmt"

func main() {
	fmt.Println("Введите первое число для проверки:")
	var number1 int
	fmt.Scan(&number1)

	fmt.Println("Введите второе число для проверки:")
	var number2 int
	fmt.Scan(&number2)

	fmt.Println("Введите третье число для проверки:")
	var number3 int
	fmt.Scan(&number3)

	positiveNumber := false

	if number1 > 0 || number2 > 0 || number3 > 0 {
		positiveNumber = true
	}

	if positiveNumber {
		fmt.Println("да, одно число является положительным")
	} else {
		fmt.Println("нет положительного числа")
	}
}
