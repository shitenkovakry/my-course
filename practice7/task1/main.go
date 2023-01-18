package main

import "fmt"

const (
	mirrorTicket = "зеркальный билет"
)

func ConvertInto4Digits(ticket int) (int, int, int, int, int, int) {
	number := ticket

	number1 := number / 100000
	number = number - number1*100000

	number2 := number / 10000
	number = number - number2*10000

	number3 := number / 1000
	number = number - number3*1000

	number4 := number / 100
	number = number - number4*100

	number5 := number / 10
	number6 := number - number5*10

	return number1, number2, number3, number4, number5, number6
}

func CheckTicket(ticket int) string {
	number1, number2, number3, number4, number5, number6 := ConvertInto4Digits(ticket)

	if number1 == number6 && number2 == number5 && number3 == number4 {
		return mirrorTicket
	}

	return ""
}

func main() {
	minNumber := 100000
	maxNumber := 999999

	foundMirrorTickets := 0

	for numberCurrent := minNumber; numberCurrent <= maxNumber; numberCurrent++ {
		typeTicket := CheckTicket(numberCurrent)

		if typeTicket == mirrorTicket {
			foundMirrorTickets++
		}
	}

	fmt.Println("Количество зеркальных билетов среди всех шестизначных номеров от", minNumber, "до", maxNumber, ":")
	fmt.Println(foundMirrorTickets)
}
