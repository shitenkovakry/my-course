package main

import "fmt"

func main() {
	fmt.Println("Введите первое число:")
	var number1 int
	fmt.Scan(&number1)

	fmt.Println("Введите второе число:")
	var number2 int
	fmt.Scan(&number2)

	if number1 < number2 {
		fmt.Println("Меньшее из двух", number1)
	} else if number1 > number2 {
		fmt.Println("Меньшее из двух", number2)
	} else {
		fmt.Println("Два числа равны")
	}
}
