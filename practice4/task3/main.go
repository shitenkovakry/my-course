package main

import "fmt"

func main() {
	fmt.Println("Банкомат.")

	fmt.Println("Введите сумму снятия со счета:")
	var amountInvoice int
	fmt.Scan(&amountInvoice)

	if amountInvoice%100 == 0 && amountInvoice < 100000 {
		fmt.Println("Операция успешно выполнена.")
		fmt.Println("Вы сняли", amountInvoice, "рублей.")
	} else if amountInvoice%100 != 0 {
		fmt.Println("К сожалению, сумму снять не получилось по причине привышения купюры номиналом 100 рублей")
	} else if amountInvoice > 100000 {
		fmt.Println("К сожалению, не получилось снять сумму денег. Сумма привышает 100000 рублей.")
	}
}
