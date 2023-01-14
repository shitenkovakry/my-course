package main

import "fmt"

func main() {
	fmt.Println("Введите число:")
	var number int
	fmt.Scan(&number)

	fmt.Println("Введите степень числа:")
	var powerOf int
	fmt.Scan(&powerOf)

	run := 1
	result := 1

	for {
		if run > powerOf {
			break
		}

		result *= number

		run++
	}

	fmt.Println("если возвести", number, "в степень", powerOf, "то будет ответ", result)
}
