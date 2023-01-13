package main

import "fmt"

func main() {
	fmt.Println("введите возраст клиента:")
	var ageOfClient int
	fmt.Scan(&ageOfClient)

	ageYoungerForDiscount := 18
	ageOlderForDiscount := 60

	if ageOfClient < ageYoungerForDiscount || ageOfClient > ageOlderForDiscount {
		fmt.Println("необходимо предоставить скидку")
		return
	}

	fmt.Println("скидки нет")
}
