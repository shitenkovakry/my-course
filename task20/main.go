package main

import "fmt"

func main() {
	fmt.Println("первое число:")
	var number1 int
	fmt.Scan(&number1)

	fmt.Println("второе число:")
	var number2 int
	fmt.Scan(&number2)

	fmt.Println("третье число:")
	var number3 int
	fmt.Scan(&number3)

	if number1 > number2 && number1 > number3 {
		fmt.Println(number1)
		return
	} else if number2 > number3 {
		fmt.Println(number2)
		return
	} else {
		fmt.Println(number3)
	}

}
