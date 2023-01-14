package main

import (
	"fmt"
)

const (
	luckyTicket   = "счастливый билет"
	mirrorTicket  = "зеркальный билет"
	regularTicket = "обычный билет"
)

func ConvertInto4Digits(ticket int) (int, int, int, int) {
	number := ticket

	number1 := number / 1000
	number = number - number1*1000

	number2 := number / 100
	number = number - number2*100

	number3 := number / 10
	number4 := number - number3*10

	return number1, number2, number3, number4
}

func CheckTicket(ticket int) string {
	number1, number2, number3, number4 := ConvertInto4Digits(ticket)

	if number1 == number4 && number2 == number3 {
		return mirrorTicket
	}

	if number1+number2 == number3+number4 {
		return luckyTicket
	}

	return regularTicket
}

func main() {
	fmt.Println("Введите номер своего билета")
	var ticket int
	fmt.Scan(&ticket)

	if ticket > 10000 {
		fmt.Println("номер билета очень большой")

		return
	}

	typeOfTicket := CheckTicket(ticket)
	fmt.Println(typeOfTicket)
}
