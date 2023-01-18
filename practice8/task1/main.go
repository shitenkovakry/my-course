package main

import "fmt"

func main() {
	fmt.Println("Времена года")

	fmt.Println("Введите месяц:")
	var month string
	fmt.Scan(&month)

	switch month {
	case "январь", "февраль", "декабрь":
		fmt.Println("зима")
	case "март", "апрель", "май":
		fmt.Println("весна")
	case "июнь", "июль", "август":
		fmt.Println("лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("осень")
	default:
		fmt.Println("некорректный ввод")
	}
}
