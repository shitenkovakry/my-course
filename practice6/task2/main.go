package main

import "fmt"

func main() {
	fmt.Println("Введите первое число:")
	var number1 int
	fmt.Scan(&number1)

	fmt.Println("Введите второе число:")
	var number2 int
	fmt.Scan(&number2)

	summaOfNumbers := number1 + number2

	for number1 < summaOfNumbers {
		number1++
	}

	fmt.Println(number1, "- сумма двух чисел", summaOfNumbers)
}
