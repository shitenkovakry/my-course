package main

import "fmt"

func main() {
	fmt.Println("введите число")
	var number int
	fmt.Scan(&number)

	fmt.Println("введите другое число")
	var number2 int
	fmt.Scan(&number2)

	if number%number2 == 0 {
		devision := number / number2
		fmt.Println("остатка нет. деление будет равно", devision)
		return
	}

	fmt.Println("остаток есть")
}
