package main

import "fmt"

func main() {
	fmt.Println("Введите любое число:")
	var number int
	fmt.Scan(&number)

	run := 0

	for {
		if run > number {
			break
		}

		fmt.Println(run)

		run++
	}
}
