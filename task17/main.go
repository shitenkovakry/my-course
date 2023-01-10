package main

import "fmt"

func main() {
	fmt.Println("Введите число дня недели:")
	var number int
	fmt.Scan(&number)

	fmt.Println("Сумма чека:")
	var check int
	fmt.Scan(&check)

	discount := check * 10 / 100
	sum := check - discount

	if number == 1 {
		fmt.Println("Конечная сумма со скидкой", sum)
		return
	} else if number == 2 {
		fmt.Println("Конечная сумма со скидкой", sum)
		return
	} else if number == 3 {
		fmt.Println("Конечная сумма со скидкой", sum)
		return
	} else if number == 4 {
		fmt.Println("Конечная сумма со скидкой", sum)
		return
	} else if number == 5 {
		fmt.Println("Конечная сумма со скидкой", sum)
		return
	} else if number == 6 && check > 10000 {
		fmt.Println("Конечная сумма со скидкой", sum)
		return
	} else if number == 7 && check > 10000 {
		fmt.Println("Конечная сумма со скидкой", sum)
		return
	}

	fmt.Println("К сожалению, не можем выполнить скидку")
}
