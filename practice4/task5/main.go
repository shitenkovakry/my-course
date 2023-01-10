package main

import "fmt"

func main() {
	fmt.Println("Ресторан")

	fmt.Println("Введите день недели:")
	var day int
	fmt.Scan(&day)

	fmt.Println("Введите число гостей:")
	var guestsInCompany int
	fmt.Scan(&guestsInCompany)

	fmt.Println("Введите сумму чека")
	var amountCheck int
	fmt.Scan(&amountCheck)

	discount := 0
	additionalService := 0

	if day == 1 {
		discount = amountCheck * 10 / 100
		fmt.Println("Скидка по понедельникам:", discount)
	}

	if day == 5 && amountCheck > 10000 {
		discount = amountCheck * 5 / 100

		fmt.Println("Скидка по пятницам:", discount)
	}

	if guestsInCompany > 5 {
		additionalService = amountCheck * 10 / 100
		fmt.Println("Надбавка за обслуживание", additionalService)
	}

	finalAmountCheck := amountCheck + additionalService - discount
	fmt.Println("Сумма к оплате:", finalAmountCheck)
}
