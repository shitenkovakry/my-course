package main

import "fmt"

func main() {
	number := -1
	sumOfNumbers := 0

	for {
		number++
		if number > 40 {
			break
		}

		if number%4 != 0 {
			continue
		}

		sumOfNumbers += number
	}

	fmt.Println(sumOfNumbers)
}
