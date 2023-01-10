package main

import "fmt"

func main() {
	fmt.Println("Введите первое число:")
	var number1 int
	fmt.Scan(&number1)

	fmt.Println("Введите второе число:")
	var number2 int
	fmt.Scan(&number2)

	fmt.Println("Введите сумму чисел:")
	var total int
	fmt.Scan(&total)

	summaNumbers := number1 + number2

	if total == summaNumbers {
		fmt.Println("Вы посчитали верно")
		return
	} //else {

	fmt.Println("Ответ неправильный. Попробуйте еще раз")
	// }
}
