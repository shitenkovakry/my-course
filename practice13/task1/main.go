package main

import (
	"fmt"
)

func CountevenNumbers(number1, number2 int) int {
	summa := 0

	for i := number1; i <= number2; i++ {
		if i%2 == 0 {
			summa += i
		}

	}

	for i := number1; i >= number2; i-- {
		if i%2 == 0 {
			summa += i
		}
	}

	return summa
}

func main() {
	number1 := 2
	number2 := 6

	result := CountevenNumbers(number1, number2)
	fmt.Println(result)
}
