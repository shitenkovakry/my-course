package main

import "fmt"

func main() {
	number1 := 0
	number2 := 0
	number3 := 0

	limitForNumber1 := 10
	limitForNumber2 := 100
	limitForNumber3 := 1000

	for {
		if number1 > limitForNumber1 {
			break
		}
		fmt.Println("первое число:", number1)
		number1++
	}

	for {
		if number2 > limitForNumber2 {
			break
		}
		fmt.Println("второе число:", number2)
		number2++
	}

	for {
		if number3 > limitForNumber3 {
			break
		}
		fmt.Println("третье число:", number3)
		number3++
	}

}
