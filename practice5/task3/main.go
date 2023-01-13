package main

import "fmt"

func main() {
	fmt.Println("Введите первое число:")
	var number1 int
	fmt.Scan(&number1)

	fmt.Println("Введите второе число:")
	var number2 int
	fmt.Scan(&number2)

	fmt.Println("Введите третье число:")
	var number3 int
	fmt.Scan(&number3)

	if number1 == number2 && number1 == number3 && number2 == number3 {
		fmt.Println("все числа совпадают")
		return
	} else if number1 != number2 && number1 != number3 && number2 != number3 {
		fmt.Println("числа не совпадают")
		return
	}

	fmt.Println("хотя бы два совпадают")
}
