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

	moreOrEqualMax := 0

	if number1 >= max {
		moreOrEqualMax++
	}

	if number2 >= max {
		moreOrEqualMax++
	}

	if number3 >= max {
		moreOrEqualMax++
	}

	fmt.Println("Среди введенных чисел", moreOrEqualMax, "больше или равны", max, ".")
}
